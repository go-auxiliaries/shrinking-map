## Description ##
As you may know [maps are not shrinking in go](https://github.com/golang/go/issues/20135#issuecomment-616643867)
This library gives you easy-to-use implementation of the map that is shrinking automatically once number of deleted elements reach the limit

## How to use ##

```go
package main

import (
	"fmt"
	"github.com/go-auxiliaries/shrinking-map/pkg/safe-map"
)

func main() {
	testMap := safe_map.New[string, string](10000000)
	testMap.Set("someKey1", "someVal1")
	testMap.Set("someKey2", "someVal2")
	// map[someKey1:someVal1 someKey2:someVal2]
	fmt.Printf("%v\n", testMap.Values())
	testMap.Delete("someKey1")
	// map[someKey2:someVal2]
	fmt.Printf("%v\n", testMap.Values())
	// someVal3
	fmt.Printf("%v\n", testMap.GetOrSet("someKey3", "someVal3"))
	// someVal3
	fmt.Printf("%v\n", testMap.GetOrSet("someKey3", "someVal4"))
}
```

## Benchmarks ##
These are the benchmarks against `sync.Map`

```shell
goos: linux
goarch: amd64
pkg: github.com/go-auxiliaries/go-shrinking-map/pkg/safe-tight-map
cpu: 12th Gen Intel(R) Core(TM) i9-12900HK
Benchmark_Set_Tight_Nonexisting-20                         44162             33461 ns/op
Benchmark_Set_Sync_Nonexisting-20                          24254             49771 ns/op
Benchmark_Set_Tight_Nonexisting_Parallel-20                29949             38683 ns/op
Benchmark_Set_Sync_Nonexisting_Parallel-20                 10000            102689 ns/op
Benchmark_Set_Tight_Existing-20                            27194             44582 ns/op
Benchmark_Set_Sync_Existing-20                             16953             66588 ns/op
Benchmark_Set_Tight_Existing_Parallel-20                   22004             53512 ns/op
Benchmark_Set_Sync_Existing_Parallel-20                    13106             83825 ns/op
Benchmark_Get_Tight_Nonexisting-20                        285482              3994 ns/op
Benchmark_Get_Sync_Nonexisting-20                         356938              3587 ns/op
Benchmark_Get_Tight_Nonexisting_Parallel-20                29298             37498 ns/op
Benchmark_Get_Sync_Nonexisting_Parallel-20                436996              3020 ns/op
Benchmark_Get_Tight_Existing-20                            27914             42464 ns/op
Benchmark_Get_Sync_Existing-20                             17552             72341 ns/op
Benchmark_Get_Tight_Existing_Parallel-20                   23348             53597 ns/op
Benchmark_Get_Sync_Existing_Parallel-20                    16075             70877 ns/op
Benchmark_GetOrSet_Tight_Nonexisting-20                    39522             36092 ns/op
Benchmark_GetOrSet_Sync_Nonexisting-20                     25434             47856 ns/op
Benchmark_GetOrSet_Tight_Nonexisting_Parallel-20           27150             40160 ns/op
Benchmark_GetOrSet_Sync_Nonexisting_Parallel-20            15829             71015 ns/op
Benchmark_GetOrSet_Tight_Existing-20                       26704             41621 ns/op
Benchmark_GetOrSet_Sync_Existing-20                        17332             72910 ns/op
Benchmark_GetOrSet_Tight_Existing_Parallel-20              24048             52383 ns/op
Benchmark_GetOrSet_Sync_Existing_Parallel-20               14742             70663 ns/op
Benchmark_Mixed_Tight_Parallel-20                          29206             38757 ns/op
Benchmark_Mixed_Sync_Parallel-20                           18379             89391 ns/op
PASS
ok      github.com/go-auxiliaries/go-shrinking-map/pkg/safe-tight-map   88.683s

```

#### Benchmark highlights ####
As you can see it is `30-70%` faster than `sync.Map` in almost all cases.
But there is one particular case that stands out:
```shell
Benchmark_Get_Tight_Nonexisting_Parallel-20                29298             37498 ns/op
Benchmark_Get_Sync_Nonexisting_Parallel-20                436996              3020 ns/op
```
Which means that it is `1200%` slower on reading values that does not exist
So, if it is your case, please use `sync.Map` instead

At the same time it worth to mention that on reading existing values it is faster than `sync.Map` by `32%`
```shell
Benchmark_Get_Tight_Existing_Parallel-20                   23348             53597 ns/op
Benchmark_Get_Sync_Existing_Parallel-20                    16075             70877 ns/op
```

Please keep that in mind making decision.
If you are not sure regarding you workload, 
take a look at mixed test case, which shows that it is `230%` faster than `sync.Map`:
```shell
Benchmark_Mixed_Tight_Parallel-20                          29206             38757 ns/op
Benchmark_Mixed_Sync_Parallel-20                           18379             89391 ns/op
```