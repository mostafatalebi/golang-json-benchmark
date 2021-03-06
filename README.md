# golang-json-benchmark
This is not for normal projects with ordinary traffic volume or resource generosity.

## To Whom It May Concern?
This benchmark is for those people whose serve is serving millions of rps, or whose service must run under limited resources and have considerable resource constraints. 

### Avoid Using GoLang's default JSON
Avoid using it, yes.
#### Why?
Because it has a generic approach toward unmarshalling. It accepts any sort of data, and maps incoming json to the passed data type (which usually is a struct).
Now, there are good libraries out there. The best of them I say is fastjson:
https://github.com/valyala/fastjson

Your must write more lines of code, indeed. You have to assign each field declaratively; but who doesn't know that GoLang is not intended to be a language with generic data types and structures? Then avoid doing it. Write declarative code.

Run the benchmark, and it simply shows you the result. It already has been tested on production. 

### Benchmark Result
```shell 
BenchmarkNormalJsonDecoding-8            5799097              6302 ns/op             976 B/op         32 allocs/op
BenchmarkFastJJsonDecoding-8             9503856              4528 ns/op            6368 B/op         32 allocs/op

```

#### Sample command for running bench
```shell
 go test -bench=. -benchmem -benchtime=30s -cpuprofile ../json.bench.out
```
