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
pkg: github.com/go-auxiliaries/shrinking-map/pkg/safe-map
cpu: 12th Gen Intel(R) Core(TM) i9-12900HK
Benchmark_TestSuite/Sync_Get_Existing_Parallel_100-20         	574680510	         2.040 ns/op
Benchmark_TestSuite/SafeMap_Get_Existing_Parallel_100-20      	16876723	        64.82 ns/op
Benchmark_TestSuite/Sync_Set_Existing_Parallel_100-20         	 4579580	       543.0 ns/op
Benchmark_TestSuite/SafeMap_Set_Existing_Parallel_100-20      	 3505146	       410.9 ns/op
Benchmark_TestSuite/Sync_Delete_Existing_Parallel_100-20      	507411679	         2.002 ns/op
Benchmark_TestSuite/SafeMap_Delete_Existing_Parallel_100-20   	 3937110	       361.9 ns/op
Benchmark_TestSuite/Sync_GetAndDelete_Existing_Parallel_100-20         	598342317	         2.109 ns/op
Benchmark_TestSuite/SafeMap_GetAndDelete_Existing_Parallel_100-20      	 5360181	       352.7 ns/op
Benchmark_TestSuite/Sync_GetOrSet_Existing_Parallel_100-20             	76702706	        18.15 ns/op
Benchmark_TestSuite/SafeMap_GetOrSet_Existing_Parallel_100-20          	 6074865	       323.1 ns/op
Benchmark_TestSuite/Sync_Mixed_Existing_Parallel_100-20                	152064001	         6.982 ns/op
Benchmark_TestSuite/SafeMap_Mixed_Existing_Parallel_100-20             	12711880	       157.9 ns/op
Benchmark_TestSuite/Sync_Get_Existing_Parallel_10000-20                	359368326	         3.193 ns/op
Benchmark_TestSuite/SafeMap_Get_Existing_Parallel_10000-20             	19517655	        60.87 ns/op
Benchmark_TestSuite/Sync_Set_Existing_Parallel_10000-20                	 2357757	       513.7 ns/op
Benchmark_TestSuite/SafeMap_Set_Existing_Parallel_10000-20             	 2529451	       518.3 ns/op
Benchmark_TestSuite/Sync_Delete_Existing_Parallel_10000-20             	346038777	         3.271 ns/op
Benchmark_TestSuite/SafeMap_Delete_Existing_Parallel_10000-20          	 3070018	       445.2 ns/op
Benchmark_TestSuite/Sync_GetAndDelete_Existing_Parallel_10000-20       	344920372	         3.223 ns/op
Benchmark_TestSuite/SafeMap_GetAndDelete_Existing_Parallel_10000-20    	 4241340	       365.9 ns/op
Benchmark_TestSuite/Sync_GetOrSet_Existing_Parallel_10000-20           	51026160	        22.38 ns/op
Benchmark_TestSuite/SafeMap_GetOrSet_Existing_Parallel_10000-20        	 3777860	       402.1 ns/op
Benchmark_TestSuite/Sync_Mixed_Existing_Parallel_10000-20              	82055820	        12.32 ns/op
Benchmark_TestSuite/SafeMap_Mixed_Existing_Parallel_10000-20           	 6319804	       231.1 ns/op
Benchmark_TestSuite/Sync_Get_Existing_Parallel_1000000-20              	100291227	        10.17 ns/op
Benchmark_TestSuite/SafeMap_Get_Existing_Parallel_1000000-20           	18184263	        62.27 ns/op
Benchmark_TestSuite/Sync_Set_Existing_Parallel_1000000-20              	 2596406	       477.8 ns/op
Benchmark_TestSuite/SafeMap_Set_Existing_Parallel_1000000-20           	 3475326	       379.1 ns/op
Benchmark_TestSuite/Sync_Delete_Existing_Parallel_1000000-20           	84168696	        12.50 ns/op
Benchmark_TestSuite/SafeMap_Delete_Existing_Parallel_1000000-20        	 3314630	       455.7 ns/op
Benchmark_TestSuite/Sync_GetAndDelete_Existing_Parallel_1000000-20     	95083435	        11.22 ns/op
Benchmark_TestSuite/SafeMap_GetAndDelete_Existing_Parallel_1000000-20  	 3912381	       372.5 ns/op
Benchmark_TestSuite/Sync_GetOrSet_Existing_Parallel_1000000-20         	30091417	        34.01 ns/op
Benchmark_TestSuite/SafeMap_GetOrSet_Existing_Parallel_1000000-20      	 3856132	       365.6 ns/op
Benchmark_TestSuite/Sync_Mixed_Existing_Parallel_1000000-20            	18821798	        53.71 ns/op
Benchmark_TestSuite/SafeMap_Mixed_Existing_Parallel_1000000-20         	 6336772	       224.2 ns/op
Benchmark_TestSuite/Sync_Get_Existing_Parallel_10000000-20             	 2987179	       423.0 ns/op
Benchmark_TestSuite/SafeMap_Get_Existing_Parallel_10000000-20          	17702481	        64.85 ns/op
Benchmark_TestSuite/Sync_Set_Existing_Parallel_10000000-20             	 2595688	       565.7 ns/op
Benchmark_TestSuite/SafeMap_Set_Existing_Parallel_10000000-20          	 3251793	       397.6 ns/op
Benchmark_TestSuite/Sync_Delete_Existing_Parallel_10000000-20          	 2948784	       516.5 ns/op
Benchmark_TestSuite/SafeMap_Delete_Existing_Parallel_10000000-20       	 3789914	       418.0 ns/op
Benchmark_TestSuite/Sync_GetAndDelete_Existing_Parallel_10000000-20    	 2654278	       539.9 ns/op
Benchmark_TestSuite/SafeMap_GetAndDelete_Existing_Parallel_10000000-20 	 3575700	       396.3 ns/op
Benchmark_TestSuite/Sync_GetOrSet_Existing_Parallel_10000000-20        	 4139848	       545.3 ns/op
Benchmark_TestSuite/SafeMap_GetOrSet_Existing_Parallel_10000000-20     	 4274433	       331.1 ns/op
Benchmark_TestSuite/Sync_Mixed_Existing_Parallel_10000000-20           	 1876588	       631.8 ns/op
Benchmark_TestSuite/SafeMap_Mixed_Existing_Parallel_10000000-20        	 8878618	       206.9 ns/op
Benchmark_TestSuite/Sync_Get_NonExisting_Parallel_100-20               	1000000000	         0.9913 ns/op
Benchmark_TestSuite/SafeMap_Get_NonExisting_Parallel_100-20            	19557538	        54.61 ns/op
Benchmark_TestSuite/Sync_Set_NonExisting_Parallel_100-20               	 2437075	       493.9 ns/op
Benchmark_TestSuite/SafeMap_Set_NonExisting_Parallel_100-20            	 4256010	       343.3 ns/op
Benchmark_TestSuite/Sync_Delete_NonExisting_Parallel_100-20            	1000000000	         1.025 ns/op
Benchmark_TestSuite/SafeMap_Delete_NonExisting_Parallel_100-20         	 9591711	       179.6 ns/op
Benchmark_TestSuite/Sync_GetAndDelete_NonExisting_Parallel_100-20      	1000000000	         1.088 ns/op
Benchmark_TestSuite/SafeMap_GetAndDelete_NonExisting_Parallel_100-20   	 8058096	       243.4 ns/op
Benchmark_TestSuite/Sync_GetOrSet_NonExisting_Parallel_100-20          	100000000	        14.24 ns/op
Benchmark_TestSuite/SafeMap_GetOrSet_NonExisting_Parallel_100-20       	 4491636	       308.9 ns/op
Benchmark_TestSuite/Sync_Mixed_NonExisting_Parallel_100-20             	191599881	         6.353 ns/op
Benchmark_TestSuite/SafeMap_Mixed_NonExisting_Parallel_100-20          	11156361	       143.5 ns/op
Benchmark_TestSuite/Sync_Get_NonExisting_Parallel_10000-20             	1000000000	         1.004 ns/op
Benchmark_TestSuite/SafeMap_Get_NonExisting_Parallel_10000-20          	18986616	        54.76 ns/op
Benchmark_TestSuite/Sync_Set_NonExisting_Parallel_10000-20             	 2453768	       489.2 ns/op
Benchmark_TestSuite/SafeMap_Set_NonExisting_Parallel_10000-20          	 2988786	       463.1 ns/op
Benchmark_TestSuite/Sync_Delete_NonExisting_Parallel_10000-20          	1000000000	         1.002 ns/op
Benchmark_TestSuite/SafeMap_Delete_NonExisting_Parallel_10000-20       	 4908854	       272.5 ns/op
Benchmark_TestSuite/Sync_GetAndDelete_NonExisting_Parallel_10000-20    	1000000000	         1.115 ns/op
Benchmark_TestSuite/SafeMap_GetAndDelete_NonExisting_Parallel_10000-20 	 4882572	       282.3 ns/op
Benchmark_TestSuite/Sync_GetOrSet_NonExisting_Parallel_10000-20        	49729287	        21.69 ns/op
Benchmark_TestSuite/SafeMap_GetOrSet_NonExisting_Parallel_10000-20     	 4430449	       346.5 ns/op
Benchmark_TestSuite/Sync_Mixed_NonExisting_Parallel_10000-20           	86142045	        13.39 ns/op
Benchmark_TestSuite/SafeMap_Mixed_NonExisting_Parallel_10000-20        	10921442	       201.8 ns/op
Benchmark_TestSuite/Sync_Get_NonExisting_Parallel_1000000-20           	1000000000	         1.125 ns/op
Benchmark_TestSuite/SafeMap_Get_NonExisting_Parallel_1000000-20        	19023562	        54.74 ns/op
Benchmark_TestSuite/Sync_Set_NonExisting_Parallel_1000000-20           	 1939941	       643.5 ns/op
Benchmark_TestSuite/SafeMap_Set_NonExisting_Parallel_1000000-20        	 3007370	       451.5 ns/op
Benchmark_TestSuite/Sync_Delete_NonExisting_Parallel_1000000-20        	1000000000	         1.057 ns/op
Benchmark_TestSuite/SafeMap_Delete_NonExisting_Parallel_1000000-20     	 4815855	       292.3 ns/op
Benchmark_TestSuite/Sync_GetAndDelete_NonExisting_Parallel_1000000-20  	1000000000	         1.015 ns/op
Benchmark_TestSuite/SafeMap_GetAndDelete_NonExisting_Parallel_1000000-20         	 5013561	       269.8 ns/op
Benchmark_TestSuite/Sync_GetOrSet_NonExisting_Parallel_1000000-20                	 2955369	       443.8 ns/op
Benchmark_TestSuite/SafeMap_GetOrSet_NonExisting_Parallel_1000000-20             	 3351228	       404.3 ns/op
Benchmark_TestSuite/Sync_Mixed_NonExisting_Parallel_1000000-20                   	 4925106	       542.0 ns/op
Benchmark_TestSuite/SafeMap_Mixed_NonExisting_Parallel_1000000-20                	 7892790	       229.5 ns/op
Benchmark_TestSuite/Sync_Get_NonExisting_Parallel_10000000-20                    	1000000000	         1.013 ns/op
Benchmark_TestSuite/SafeMap_Get_NonExisting_Parallel_10000000-20                 	19985877	        54.51 ns/op
Benchmark_TestSuite/Sync_Set_NonExisting_Parallel_10000000-20                    	 2355073	       535.4 ns/op
Benchmark_TestSuite/SafeMap_Set_NonExisting_Parallel_10000000-20                 	 3000261	       433.3 ns/op
Benchmark_TestSuite/Sync_Delete_NonExisting_Parallel_10000000-20                 	1000000000	         1.005 ns/op
Benchmark_TestSuite/SafeMap_Delete_NonExisting_Parallel_10000000-20              	 6978391	       284.4 ns/op
Benchmark_TestSuite/Sync_GetAndDelete_NonExisting_Parallel_10000000-20           	1000000000	         1.035 ns/op
Benchmark_TestSuite/SafeMap_GetAndDelete_NonExisting_Parallel_10000000-20        	 5544055	       246.2 ns/op
Benchmark_TestSuite/Sync_GetOrSet_NonExisting_Parallel_10000000-20               	 2765965	       418.7 ns/op
Benchmark_TestSuite/SafeMap_GetOrSet_NonExisting_Parallel_10000000-20            	 3392496	       458.0 ns/op
Benchmark_TestSuite/Sync_Mixed_NonExisting_Parallel_10000000-20                  	 3135346	       613.0 ns/op
Benchmark_TestSuite/SafeMap_Mixed_NonExisting_Parallel_10000000-20               	 6366879	       233.6 ns/op
Benchmark_TestSuite/Sync_Get_Existing_Linear_100-20                              	63669410	        18.51 ns/op
Benchmark_TestSuite/SafeMap_Get_Existing_Linear_100-20                           	64913563	        18.40 ns/op
Benchmark_TestSuite/Sync_Set_Existing_Linear_100-20                              	10951699	       105.7 ns/op
Benchmark_TestSuite/SafeMap_Set_Existing_Linear_100-20                           	38890647	        30.86 ns/op
Benchmark_TestSuite/Sync_Delete_Existing_Linear_100-20                           	70088430	        17.93 ns/op
Benchmark_TestSuite/SafeMap_Delete_Existing_Linear_100-20                        	36791875	        33.38 ns/op
Benchmark_TestSuite/Sync_GetAndDelete_Existing_Linear_100-20                     	66952568	        17.64 ns/op
Benchmark_TestSuite/SafeMap_GetAndDelete_Existing_Linear_100-20                  	36971284	        32.07 ns/op
Benchmark_TestSuite/Sync_GetOrSet_Existing_Linear_100-20                         	16298546	        71.37 ns/op
Benchmark_TestSuite/SafeMap_GetOrSet_Existing_Linear_100-20                      	39489004	        30.43 ns/op
Benchmark_TestSuite/Sync_Mixed_Existing_Linear_100-20                            	25721468	        44.84 ns/op
Benchmark_TestSuite/SafeMap_Mixed_Existing_Linear_100-20                         	41185646	        29.27 ns/op
Benchmark_TestSuite/Sync_Get_Existing_Linear_10000-20                            	31513636	        37.38 ns/op
Benchmark_TestSuite/SafeMap_Get_Existing_Linear_10000-20                         	32943967	        36.55 ns/op
Benchmark_TestSuite/Sync_Set_Existing_Linear_10000-20                            	10481876	       114.0 ns/op
Benchmark_TestSuite/SafeMap_Set_Existing_Linear_10000-20                         	29151861	        38.63 ns/op
Benchmark_TestSuite/Sync_Delete_Existing_Linear_10000-20                         	31749602	        36.20 ns/op
Benchmark_TestSuite/SafeMap_Delete_Existing_Linear_10000-20                      	21368912	        47.43 ns/op
Benchmark_TestSuite/Sync_GetAndDelete_Existing_Linear_10000-20                   	32354337	        36.58 ns/op
Benchmark_TestSuite/SafeMap_GetAndDelete_Existing_Linear_10000-20                	26819620	        44.48 ns/op
Benchmark_TestSuite/Sync_GetOrSet_Existing_Linear_10000-20                       	16062680	        76.43 ns/op
Benchmark_TestSuite/SafeMap_GetOrSet_Existing_Linear_10000-20                    	28636717	        41.72 ns/op
Benchmark_TestSuite/Sync_Mixed_Existing_Linear_10000-20                          	17729290	        61.54 ns/op
Benchmark_TestSuite/SafeMap_Mixed_Existing_Linear_10000-20                       	22947308	        44.87 ns/op
Benchmark_TestSuite/Sync_Get_Existing_Linear_1000000-20                          	14247619	        83.03 ns/op
Benchmark_TestSuite/SafeMap_Get_Existing_Linear_1000000-20                       	15688119	        77.00 ns/op
Benchmark_TestSuite/Sync_Set_Existing_Linear_1000000-20                          	 5849140	       197.3 ns/op
Benchmark_TestSuite/SafeMap_Set_Existing_Linear_1000000-20                       	11937541	        84.86 ns/op
Benchmark_TestSuite/Sync_Delete_Existing_Linear_1000000-20                       	14981034	        77.54 ns/op
Benchmark_TestSuite/SafeMap_Delete_Existing_Linear_1000000-20                    	10694175	        97.42 ns/op
Benchmark_TestSuite/Sync_GetAndDelete_Existing_Linear_1000000-20                 	14996866	        79.53 ns/op
Benchmark_TestSuite/SafeMap_GetAndDelete_Existing_Linear_1000000-20              	12510092	        95.59 ns/op
Benchmark_TestSuite/Sync_GetOrSet_Existing_Linear_1000000-20                     	 7006353	       154.5 ns/op
Benchmark_TestSuite/SafeMap_GetOrSet_Existing_Linear_1000000-20                  	13955450	        86.87 ns/op
Benchmark_TestSuite/Sync_Mixed_Existing_Linear_1000000-20                        	10743942	       110.1 ns/op
Benchmark_TestSuite/SafeMap_Mixed_Existing_Linear_1000000-20                     	12147165	        91.64 ns/op
Benchmark_TestSuite/Sync_Get_Existing_Linear_10000000-20                         	 9198398	       154.3 ns/op
Benchmark_TestSuite/SafeMap_Get_Existing_Linear_10000000-20                      	14745873	        85.11 ns/op
Benchmark_TestSuite/Sync_Set_Existing_Linear_10000000-20                         	 5493913	       211.5 ns/op
Benchmark_TestSuite/SafeMap_Set_Existing_Linear_10000000-20                      	13110157	        95.38 ns/op
Benchmark_TestSuite/Sync_Delete_Existing_Linear_10000000-20                      	 6869094	       166.6 ns/op
Benchmark_TestSuite/SafeMap_Delete_Existing_Linear_10000000-20                   	12890013	       110.6 ns/op
Benchmark_TestSuite/Sync_GetAndDelete_Existing_Linear_10000000-20                	 7507372	       151.1 ns/op
Benchmark_TestSuite/SafeMap_GetAndDelete_Existing_Linear_10000000-20             	 8877446	       156.3 ns/op
Benchmark_TestSuite/Sync_GetOrSet_Existing_Linear_10000000-20                    	 6520016	       191.4 ns/op
Benchmark_TestSuite/SafeMap_GetOrSet_Existing_Linear_10000000-20                 	14933388	        95.09 ns/op
Benchmark_TestSuite/Sync_Mixed_Existing_Linear_10000000-20                       	 6675783	       190.3 ns/op
Benchmark_TestSuite/SafeMap_Mixed_Existing_Linear_10000000-20                    	13785230	       109.4 ns/op
Benchmark_TestSuite/Sync_Get_NonExisting_Linear_100-20                           	100000000	        10.07 ns/op
Benchmark_TestSuite/SafeMap_Get_NonExisting_Linear_100-20                        	77151236	        15.44 ns/op
Benchmark_TestSuite/Sync_Set_NonExisting_Linear_100-20                           	10429992	       110.2 ns/op
Benchmark_TestSuite/SafeMap_Set_NonExisting_Linear_100-20                        	38925982	        30.80 ns/op
Benchmark_TestSuite/Sync_Delete_NonExisting_Linear_100-20                        	131633998	         9.116 ns/op
Benchmark_TestSuite/SafeMap_Delete_NonExisting_Linear_100-20                     	42306814	        28.54 ns/op
Benchmark_TestSuite/Sync_GetAndDelete_NonExisting_Linear_100-20                  	129460143	         9.652 ns/op
Benchmark_TestSuite/SafeMap_GetAndDelete_NonExisting_Linear_100-20               	39100035	        28.05 ns/op
Benchmark_TestSuite/Sync_GetOrSet_NonExisting_Linear_100-20                      	17955370	        67.31 ns/op
Benchmark_TestSuite/SafeMap_GetOrSet_NonExisting_Linear_100-20                   	39377995	        30.51 ns/op
Benchmark_TestSuite/Sync_Mixed_NonExisting_Linear_100-20                         	30521538	        40.46 ns/op
Benchmark_TestSuite/SafeMap_Mixed_NonExisting_Linear_100-20                      	36812632	        35.37 ns/op
Benchmark_TestSuite/Sync_Get_NonExisting_Linear_10000-20                         	100000000	        11.24 ns/op
Benchmark_TestSuite/SafeMap_Get_NonExisting_Linear_10000-20                      	77718800	        15.43 ns/op
Benchmark_TestSuite/Sync_Set_NonExisting_Linear_10000-20                         	10738016	       108.1 ns/op
Benchmark_TestSuite/SafeMap_Set_NonExisting_Linear_10000-20                      	30988747	        38.20 ns/op
Benchmark_TestSuite/Sync_Delete_NonExisting_Linear_10000-20                      	129764692	         9.223 ns/op
Benchmark_TestSuite/SafeMap_Delete_NonExisting_Linear_10000-20                   	41361781	        28.62 ns/op
Benchmark_TestSuite/Sync_GetAndDelete_NonExisting_Linear_10000-20                	129433466	         9.120 ns/op
Benchmark_TestSuite/SafeMap_GetAndDelete_NonExisting_Linear_10000-20             	43091794	        27.81 ns/op
Benchmark_TestSuite/Sync_GetOrSet_NonExisting_Linear_10000-20                    	16638021	        74.34 ns/op
Benchmark_TestSuite/SafeMap_GetOrSet_NonExisting_Linear_10000-20                 	29615571	        40.20 ns/op
Benchmark_TestSuite/Sync_Mixed_NonExisting_Linear_10000-20                       	22718679	        54.90 ns/op
Benchmark_TestSuite/SafeMap_Mixed_NonExisting_Linear_10000-20                    	28340762	        41.39 ns/op
Benchmark_TestSuite/Sync_Get_NonExisting_Linear_1000000-20                       	131575062	         9.219 ns/op
Benchmark_TestSuite/SafeMap_Get_NonExisting_Linear_1000000-20                    	77042383	        15.46 ns/op
Benchmark_TestSuite/Sync_Set_NonExisting_Linear_1000000-20                       	 4477153	       237.4 ns/op
Benchmark_TestSuite/SafeMap_Set_NonExisting_Linear_1000000-20                    	11701244	        95.81 ns/op
Benchmark_TestSuite/Sync_Delete_NonExisting_Linear_1000000-20                    	137279714	         8.774 ns/op
Benchmark_TestSuite/SafeMap_Delete_NonExisting_Linear_1000000-20                 	42228463	        28.65 ns/op
Benchmark_TestSuite/Sync_GetAndDelete_NonExisting_Linear_1000000-20              	129309384	         9.037 ns/op
Benchmark_TestSuite/SafeMap_GetAndDelete_NonExisting_Linear_1000000-20           	42561860	        27.90 ns/op
Benchmark_TestSuite/Sync_GetOrSet_NonExisting_Linear_1000000-20                  	 5101731	       201.4 ns/op
Benchmark_TestSuite/SafeMap_GetOrSet_NonExisting_Linear_1000000-20               	10424455	       104.7 ns/op
Benchmark_TestSuite/Sync_Mixed_NonExisting_Linear_1000000-20                     	10509524	       107.5 ns/op
Benchmark_TestSuite/SafeMap_Mixed_NonExisting_Linear_1000000-20                  	16981828	        71.11 ns/op
Benchmark_TestSuite/Sync_Get_NonExisting_Linear_10000000-20                      	133417846	         8.831 ns/op
Benchmark_TestSuite/SafeMap_Get_NonExisting_Linear_10000000-20                   	76370257	        15.51 ns/op
Benchmark_TestSuite/Sync_Set_NonExisting_Linear_10000000-20                      	 3098736	       397.0 ns/op
Benchmark_TestSuite/SafeMap_Set_NonExisting_Linear_10000000-20                   	 6665934	       199.8 ns/op
Benchmark_TestSuite/Sync_Delete_NonExisting_Linear_10000000-20                   	138650385	         8.696 ns/op
Benchmark_TestSuite/SafeMap_Delete_NonExisting_Linear_10000000-20                	40610791	        28.82 ns/op
Benchmark_TestSuite/Sync_GetAndDelete_NonExisting_Linear_10000000-20             	132316108	         9.026 ns/op
Benchmark_TestSuite/SafeMap_GetAndDelete_NonExisting_Linear_10000000-20          	41460210	        28.15 ns/op
Benchmark_TestSuite/Sync_GetOrSet_NonExisting_Linear_10000000-20                 	 2938242	       378.4 ns/op
Benchmark_TestSuite/SafeMap_GetOrSet_NonExisting_Linear_10000000-20              	 5678400	       252.5 ns/op
Benchmark_TestSuite/Sync_Mixed_NonExisting_Linear_10000000-20                    	 4031577	       370.9 ns/op
Benchmark_TestSuite/SafeMap_Mixed_NonExisting_Linear_10000000-20                 	12334532	       164.2 ns/op
PASS
ok  	github.com/go-auxiliaries/shrinking-map/pkg/safe-map	950.004s
```

#### Benchmark highlights ####
As you can see it is slower than `sync.Map` in almost all cases when parallel access takes place.
But on huge dataset it is getting `300%` faster:
```shell
Benchmark_TestSuite/Sync_Mixed_Existing_Parallel_10000000-20           	 1876588	       631.8 ns/op
Benchmark_TestSuite/SafeMap_Mixed_Existing_Parallel_10000000-20        	 8878618	       206.9 ns/op
Benchmark_TestSuite/Sync_Mixed_NonExisting_Parallel_10000000-20                  	 3135346	       613.0 ns/op
Benchmark_TestSuite/SafeMap_Mixed_NonExisting_Parallel_10000000-20               	 6366879	       233.6 ns/op
```

So, if you are expecting `>10000000` unique keys, `rwMutex` over map is your choice.
Otherwise, stick to `sync.Map`.
