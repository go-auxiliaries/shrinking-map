package safe_map_test

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"testing"

	safe_map "github.com/go-auxiliaries/shrinking-map/pkg/safe-map"
)

const nRoutines = 20
const neverShrink = 1000000000000

type testFunc func(i int) (fillExisting func(), testBody func(tag int, name string))
type funcName string

const (
	syncGet          = funcName("Sync_Get")          // -> Load
	syncSet          = funcName("Sync_Set")          // -> Store
	syncDelete       = funcName("Sync_Delete")       // -> Delete
	syncGetAndDelete = funcName("Sync_GetAndDelete") // -> LoadAndDelete
	syncGetOrSet     = funcName("Sync_GetOrSet")     // -> LoadOrStore

	syncMixed = funcName("Sync_Mixed") // -> All together

	safeMapGet          = funcName("SafeMap_Get")
	safeMapSet          = funcName("SafeMap_Set")
	safeMapDelete       = funcName("SafeMap_Delete")
	safeMapGetAndDelete = funcName("SafeMap_GetAndDelete")
	safeMapGetOrSet     = funcName("SafeMap_GetOrSet")

	safeMapMixed = funcName("SafeMap_Mixed") // -> All together
)

var funcNameToTestFunc = map[funcName]testFunc{
	syncGet: func(i int) (fillExisting func(), testBody func(tag int, name string)) {
		testMap := sync.Map{}
		return func() {
				fillSyncMap(i, &testMap)
			},
			func(tag int, name string) {
				testMap.Load(name)
			}
	},
	syncSet: func(i int) (fillExisting func(), testBody func(tag int, name string)) {
		testMap := sync.Map{}
		return func() {
				fillSyncMap(i, &testMap)
			},
			func(tag int, name string) {
				testMap.Store(name, name)
			}
	},
	syncDelete: func(i int) (fillExisting func(), testBody func(tag int, name string)) {
		testMap := sync.Map{}
		return func() {
				fillSyncMap(i, &testMap)
			},
			func(tag int, name string) {
				testMap.Delete(name)
			}
	},
	syncGetAndDelete: func(i int) (fillExisting func(), testBody func(tag int, name string)) {
		testMap := sync.Map{}
		return func() {
				fillSyncMap(i, &testMap)
			},
			func(tag int, name string) {
				testMap.LoadAndDelete(name)
			}
	},
	syncGetOrSet: func(i int) (fillExisting func(), testBody func(tag int, name string)) {
		testMap := sync.Map{}
		return func() {
				fillSyncMap(i, &testMap)
			},
			func(tag int, name string) {
				testMap.LoadOrStore(name, name)
			}
	},
	syncMixed: func(i int) (fillExisting func(), testBody func(tag int, name string)) {
		testMap := sync.Map{}
		return func() {
				fillSyncMap(i, &testMap)
			},
			func(tag int, name string) {
				switch tag % 5 {
				case 0:
					testMap.LoadOrStore(name, name)
				case 1:
					testMap.Load(name)
				case 2:
					testMap.Store(name, name)
				case 3:
					testMap.Delete(name)
				case 4:
					testMap.LoadAndDelete(name)
				}
			}
	},
	safeMapGet: func(i int) (fillExisting func(), testBody func(tag int, name string)) {
		testMap := safe_map.New[string, string](neverShrink)
		return func() {
				fillSTagsMap(i, testMap)
			},
			func(tag int, name string) {
				testMap.Get(name)
			}
	},
	safeMapSet: func(i int) (fillExisting func(), testBody func(tag int, name string)) {
		testMap := safe_map.New[string, string](neverShrink)
		return func() {
				fillSTagsMap(i, testMap)
			},
			func(tag int, name string) {
				testMap.Set(name, name)
			}
	},
	safeMapDelete: func(i int) (fillExisting func(), testBody func(tag int, name string)) {
		testMap := safe_map.New[string, string](neverShrink)
		return func() {
				fillSTagsMap(i, testMap)
			},
			func(tag int, name string) {
				testMap.Delete(name)
			}
	},
	safeMapGetAndDelete: func(i int) (fillExisting func(), testBody func(tag int, name string)) {
		testMap := safe_map.New[string, string](neverShrink)
		return func() {
				fillSTagsMap(i, testMap)
			},
			func(tag int, name string) {
				testMap.GetAndDelete(name)
			}
	},
	safeMapGetOrSet: func(i int) (fillExisting func(), testBody func(tag int, name string)) {
		testMap := safe_map.New[string, string](neverShrink)
		return func() {
				fillSTagsMap(i, testMap)
			},
			func(tag int, name string) {
				testMap.GetOrSet(name, name)
			}
	},
	safeMapMixed: func(i int) (fillExisting func(), testBody func(tag int, name string)) {
		testMap := safe_map.New[string, string](neverShrink)
		return func() {
				fillSTagsMap(i, testMap)
			},
			func(tag int, name string) {
				switch tag % 5 {
				case 0:
					testMap.GetOrSet(name, name)
				case 1:
					testMap.Get(name)
				case 2:
					testMap.Set(name, name)
				case 3:
					testMap.Delete(name)
				case 4:
					testMap.GetAndDelete(name)
				}
			}
	},
}

type testCase struct {
	name string
	body func(*testing.B)
}

func Benchmark_TestSuite(b *testing.B) {
	testCases := make([]testCase, 0)
	for _, parallel := range []bool{true, false} {
		for _, existing := range []bool{true, false} {
			for _, nUniqueKeys := range []int{100, 10000, 1000000, 10000000} {
				for _, fName := range []funcName{
					syncGet, safeMapGet,
					syncSet, safeMapSet,
					syncDelete, safeMapDelete,
					syncGetAndDelete, safeMapGetAndDelete,
					syncGetOrSet, safeMapGetOrSet,
					syncMixed, safeMapMixed,
				} {
					parallelLocal := parallel
					existingLocal := existing
					nUniqueKeysLocal := nUniqueKeys
					fNameLocal := fName
					p := "Parallel"
					if parallelLocal == false {
						p = "Linear"
					}
					e := "Existing"
					if existingLocal == false {
						e = "NonExisting"
					}
					testCases = append(testCases, testCase{
						name: fmt.Sprintf("%s_%s_%s_%d", fNameLocal, e, p, nUniqueKeysLocal),
						body: func(b *testing.B) {
							i := b.N
							createExistingFn, testBody := funcNameToTestFunc[fNameLocal](nUniqueKeysLocal)
							if existingLocal {
								createExistingFn()
							}
							if parallelLocal {
								runParallelTest(b, i, nUniqueKeysLocal, testBody)
							} else {
								runSingleThreadTest(b, i, nUniqueKeysLocal, testBody)
							}
						},
					})
				}
			}
		}
	}
	for _, t := range testCases {
		b.Run(t.name, t.body)
	}
}

var cachedTagNames = make([]string, 0)

func getNamesList(nIterates int) []string {
	if len(cachedTagNames) < nIterates {
		for n := len(cachedTagNames); n <= nIterates; n++ {
			cachedTagNames = append(cachedTagNames, strconv.Itoa(n))
		}
	}
	return cachedTagNames
}

func runParallelTest(b *testing.B, nIterates, nUniqueKeys int, body func(tag int, name string)) {
	wg := sync.WaitGroup{}
	iteratesPerRoutine := nIterates / nRoutines
	tagNames := getNamesList(nUniqueKeys)
	runtime.GC()
	b.ResetTimer()
	for n := 0; n < nRoutines; n++ {
		wg.Add(1)
		go func() {
			for k := 0; k < iteratesPerRoutine; k++ {
				tag := k % nUniqueKeys
				body(tag, tagNames[tag])
			}
			wg.Done()
		}()
	}
	wg.Wait()
	b.StopTimer()
}

func runSingleThreadTest(b *testing.B, nIterates, nUniqueKeys int, body func(tag int, name string)) {
	tagNames := getNamesList(nUniqueKeys)
	runtime.GC()
	b.ResetTimer()
	for n := 0; n < nIterates; n++ {
		tag := n % nUniqueKeys
		body(tag, tagNames[tag])
	}
	b.StopTimer()
}

func fillSTagsMap(nIterates int, testMap *safe_map.Map[string, string]) {
	for n := 0; n <= nIterates; n++ {
		strVal := strconv.Itoa(n)
		testMap.Set(strVal, strVal)
	}
}

func fillSyncMap(nIterates int, testMap *sync.Map) {
	for n := 0; n <= nIterates; n++ {
		strVal := strconv.Itoa(n)
		testMap.Store(strVal, strVal)
	}
}
