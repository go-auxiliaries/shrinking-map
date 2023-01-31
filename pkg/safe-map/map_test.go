package safe_map_test

import (
	"math/rand"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/go-auxiliaries/shrinking-map/pkg/safe-map"
	"github.com/thanhpk/randstr"
)

func Benchmark_Set_Tight_Nonexisting(b *testing.B) {
	testMap := safe_map.New[string, string](10000000)
	id := int64(0)
	runSingleThreadTest(b.N, 100, func() {
		val := atomic.AddInt64(&id, 1)
		strVal := strconv.FormatInt(val, 10)
		testMap.Set(strVal, strVal)
	})
}

func Benchmark_Set_Sync_Nonexisting(b *testing.B) {
	testMap := sync.Map{}

	id := int64(0)
	runSingleThreadTest(b.N, 100, func() {
		val := atomic.AddInt64(&id, 1)
		strVal := strconv.FormatInt(val, 10)
		testMap.Store(strVal, strVal)
	})
}

func Benchmark_Set_Tight_Nonexisting_Parallel(b *testing.B) {
	testMap := safe_map.New[string, string](10000000)

	id := int64(0)
	runParallelTest(b.N, 100, func() {
		val := atomic.AddInt64(&id, 1)
		strVal := strconv.FormatInt(val, 10)
		testMap.Set(strVal, strVal)
	})
}

func Benchmark_Set_Sync_Nonexisting_Parallel(b *testing.B) {
	testMap := sync.Map{}

	id := int64(0)
	runParallelTest(b.N, 100, func() {
		val := atomic.AddInt64(&id, 1)
		strVal := strconv.FormatInt(val, 10)
		testMap.Store(strVal, strVal)
	})
}

func Benchmark_Set_Tight_Existing(b *testing.B) {
	testMap := safe_map.New[string, string](10000000)

	fillTightMap(b.N, 100, testMap)
	id := int64(0)
	runSingleThreadTest(b.N, 100, func() {
		val := atomic.AddInt64(&id, 1)
		strVal := strconv.FormatInt(val, 10)
		testMap.Set(strVal, strVal)
	})
}

func Benchmark_Set_Sync_Existing(b *testing.B) {
	testMap := sync.Map{}

	fillSyncMap(b.N, 100, &testMap)
	id := int64(0)
	runSingleThreadTest(b.N, 100, func() {
		val := atomic.AddInt64(&id, 1)
		strVal := strconv.FormatInt(val, 10)
		testMap.Store(strVal, strVal)
	})
}

func Benchmark_Set_Tight_Existing_Parallel(b *testing.B) {
	testMap := safe_map.New[string, string](10000000)

	fillTightMap(b.N, 100, testMap)
	id := int64(0)
	runParallelTest(b.N, 100, func() {
		val := atomic.AddInt64(&id, 1)
		strVal := strconv.FormatInt(val, 10)
		testMap.Set(strVal, strVal)
	})
}

func Benchmark_Set_Sync_Existing_Parallel(b *testing.B) {
	testMap := sync.Map{}

	fillSyncMap(b.N, 100, &testMap)
	id := int64(0)
	runParallelTest(b.N, 100, func() {
		val := atomic.AddInt64(&id, 1)
		strVal := strconv.FormatInt(val, 10)
		testMap.Store(strVal, strVal)
	})
}

func Benchmark_Get_Tight_Nonexisting(b *testing.B) {
	testMap := safe_map.New[string, string](10000000)

	id := int64(0)
	runSingleThreadTest(b.N, 100, func() {
		val := atomic.AddInt64(&id, 1)
		strVal := strconv.FormatInt(val, 10)
		testMap.Get(strVal)
	})
}

func Benchmark_Get_Sync_Nonexisting(b *testing.B) {
	testMap := sync.Map{}

	id := int64(0)
	runSingleThreadTest(b.N, 100, func() {
		val := atomic.AddInt64(&id, 1)
		strVal := strconv.FormatInt(val, 10)
		testMap.Load(strVal)
	})
}

func Benchmark_Get_Tight_Nonexisting_Parallel(b *testing.B) {
	testMap := safe_map.New[string, string](10000000)

	id := int64(0)
	runParallelTest(b.N, 100, func() {
		val := atomic.AddInt64(&id, 1)
		strVal := strconv.FormatInt(val, 10)
		testMap.Set(strVal, strVal)
	})
}

func Benchmark_Get_Sync_Nonexisting_Parallel(b *testing.B) {
	testMap := sync.Map{}

	id := int64(0)
	runParallelTest(b.N, 100, func() {
		val := atomic.AddInt64(&id, 1)
		strVal := strconv.FormatInt(val, 10)
		testMap.Load(strVal)
	})
}

func Benchmark_Get_Tight_Existing(b *testing.B) {
	testMap := safe_map.New[string, string](10000000)

	fillTightMap(b.N, 100, testMap)
	id := int64(0)
	runSingleThreadTest(b.N, 100, func() {
		val := atomic.AddInt64(&id, 1)
		strVal := strconv.FormatInt(val, 10)
		testMap.Get(strVal)
	})
}

func Benchmark_Get_Sync_Existing(b *testing.B) {
	testMap := sync.Map{}

	fillSyncMap(b.N, 100, &testMap)
	id := int64(0)
	runSingleThreadTest(b.N, 100, func() {
		val := atomic.AddInt64(&id, 1)
		strVal := strconv.FormatInt(val, 10)
		testMap.Load(strVal)
	})
}

func Benchmark_Get_Tight_Existing_Parallel(b *testing.B) {
	testMap := safe_map.New[string, string](10000000)

	fillTightMap(b.N, 100, testMap)
	id := int64(0)
	runParallelTest(b.N, 100, func() {
		val := atomic.AddInt64(&id, 1)
		strVal := strconv.FormatInt(val, 10)
		testMap.Set(strVal, strVal)
	})
}

func Benchmark_Get_Sync_Existing_Parallel(b *testing.B) {
	testMap := sync.Map{}

	fillSyncMap(b.N, 100, &testMap)
	id := int64(0)
	runParallelTest(b.N, 100, func() {
		val := atomic.AddInt64(&id, 1)
		strVal := strconv.FormatInt(val, 10)
		testMap.Load(strVal)
	})
}

func Benchmark_GetOrSet_Tight_Nonexisting(b *testing.B) {
	testMap := safe_map.New[string, string](10000000)

	id := int64(0)
	runSingleThreadTest(b.N, 100, func() {
		val := atomic.AddInt64(&id, 1)
		strVal := strconv.FormatInt(val, 10)
		testMap.GetOrSet(strVal, strVal)

	})
}

func Benchmark_GetOrSet_Sync_Nonexisting(b *testing.B) {
	testMap := sync.Map{}

	id := int64(0)
	runSingleThreadTest(b.N, 100, func() {
		val := atomic.AddInt64(&id, 1)
		strVal := strconv.FormatInt(val, 10)
		testMap.LoadOrStore(strVal, strVal)
	})
}

func Benchmark_GetOrSet_Tight_Nonexisting_Parallel(b *testing.B) {
	testMap := safe_map.New[string, string](10000000)

	id := int64(0)
	runParallelTest(b.N, 100, func() {
		val := atomic.AddInt64(&id, 1)
		strVal := strconv.FormatInt(val, 10)
		testMap.GetOrSet(strVal, strVal)

	})
}

func Benchmark_GetOrSet_Sync_Nonexisting_Parallel(b *testing.B) {
	testMap := sync.Map{}

	id := int64(0)
	runParallelTest(b.N, 100, func() {
		val := atomic.AddInt64(&id, 1)
		strVal := strconv.FormatInt(val, 10)
		testMap.LoadOrStore(strVal, strVal)
	})
}

func Benchmark_GetOrSet_Tight_Existing(b *testing.B) {
	testMap := safe_map.New[string, string](10000000)

	fillTightMap(b.N, 100, testMap)
	id := int64(0)
	runSingleThreadTest(b.N, 100, func() {
		val := atomic.AddInt64(&id, 1)
		strVal := strconv.FormatInt(val, 10)
		testMap.GetOrSet(strVal, strVal)
	})
}

func Benchmark_GetOrSet_Sync_Existing(b *testing.B) {
	testMap := sync.Map{}

	fillSyncMap(b.N, 100, &testMap)
	id := int64(0)
	runSingleThreadTest(b.N, 100, func() {
		val := atomic.AddInt64(&id, 1)
		strVal := strconv.FormatInt(val, 10)
		testMap.LoadOrStore(strVal, strVal)
	})
}

func Benchmark_GetOrSet_Tight_Existing_Parallel(b *testing.B) {
	testMap := safe_map.New[string, string](10000000)

	fillTightMap(b.N, 100, testMap)
	id := int64(0)
	runParallelTest(b.N, 100, func() {
		val := atomic.AddInt64(&id, 1)
		strVal := strconv.FormatInt(val, 10)
		testMap.GetOrSet(strVal, strVal)

	})
}

func Benchmark_GetOrSet_Sync_Existing_Parallel(b *testing.B) {
	testMap := sync.Map{}

	fillSyncMap(b.N, 100, &testMap)
	id := int64(0)
	runParallelTest(b.N, 100, func() {
		val := atomic.AddInt64(&id, 1)
		strVal := strconv.FormatInt(val, 10)
		testMap.LoadOrStore(strVal, strVal)
	})
}

func Benchmark_Mixed_Tight_Parallel(b *testing.B) {
	testMap := safe_map.New[string, string](10000000)

	fillTightMap(b.N, 10, testMap)
	id := int64(0)
	runParallelTest(b.N, 100, func() {
		existing := rand.Intn(10000)%2 == 1
		val := atomic.AddInt64(&id, 1)
		strVal := strconv.FormatInt(val, 10)
		if !existing {
			strVal = "nonExisting_" + strVal
		}
		switch val % 5 {
		case 0:
			testMap.GetOrSet(strVal, strVal)
		case 1:
			testMap.Get(strVal)
		case 2:
			testMap.Set(strVal, strVal)
		case 3:
			testMap.Delete(strVal)
		case 4:
			testMap.GetAndDelete(strVal)
		}
	})
}

func Benchmark_Mixed_Sync_Parallel(b *testing.B) {
	testMap := sync.Map{}

	fillSyncMap(b.N, 10, &testMap)
	id := int64(0)
	runParallelTest(b.N, 100, func() {
		existing := rand.Intn(10000)%2 == 1
		val := atomic.AddInt64(&id, 1)
		strVal := strconv.FormatInt(val, 10)
		if !existing {
			strVal = "nonExisting_" + strVal
		}
		switch val % 5 {
		case 0:
			testMap.LoadOrStore(strVal, strVal)
		case 1:
			testMap.Load(strVal)
		case 2:
			testMap.Store(strVal, strVal)
		case 3:
			testMap.Delete(strVal)
		case 4:
			testMap.LoadAndDelete(strVal)

		}
	})
}

func Test_SafeMap_Mixed(t *testing.T) {
	testMap := safe_map.New[string, string](1000)
	knownIndexes := make([]string, 10000)
	for id := range knownIndexes {
		knownIndexes[id] = randstr.Base64(10)
	}
	lastKnownIndex := len(knownIndexes) - 1
	wg := sync.WaitGroup{}
	for n := 0; n < 100; n++ {
		wg.Add(1)
		go func() {
			for k := 0; k < 10000; k++ {
				var key string
				switch k % 2 {
				case 1:
					key = knownIndexes[rand.Intn(lastKnownIndex)]
				default:
					key = randstr.Hex(20)
				}
				switch k % 5 {
				case 0:
					testMap.GetOrSet(key, key)
				case 1:
					testMap.Get(key)
				case 2:
					testMap.Set(key, key)
				case 3:
					testMap.Delete(key)
				case 4:
					testMap.GetAndDelete(key)
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()

}

func runParallelTest(nRoutines, nIterates int, body func()) {
	wg := sync.WaitGroup{}
	for n := 0; n < nRoutines; n++ {
		wg.Add(1)
		go func() {
			for k := 0; k < nIterates; k++ {
				body()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func runSingleThreadTest(nTimes, nIterates int, body func()) {
	for n := 0; n < nTimes*nIterates; n++ {
		body()
	}
}

func fillTightMap(nTimes, nIterates int, testMap *safe_map.Map[string, string]) {
	for n := 1; n <= nTimes*nIterates; n++ {
		strVal := strconv.FormatInt(int64(n), 10)
		testMap.Set(strVal, strVal)
	}
}

func fillSyncMap(nTimes, nIterates int, testMap *sync.Map) {
	for n := 1; n <= nTimes*nIterates; n++ {
		strVal := strconv.FormatInt(int64(n), 10)
		testMap.Store(strVal, strVal)
	}
}
