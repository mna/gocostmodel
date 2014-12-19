# gocostmodel

This package was inspired by Brian W. Kernighan and Rob Pike's book "The Practice of Programming" (Addison-Wesley, 1999). In Chapter 7 on performance, section 7.6 (Estimation):

> "It's hard to estimate ahead of time how fast a program will be [...]. It's easy, though, to create a *cost model* for a language or a system, which will give you at least a rough idea of how long important operations take."

This is a *cost model* that benchmarks common basic operations for the Go language. For it to be meaningful, it should be run on your target hardware, and the relative performance of operations is probably more important than the actual numbers. Examples of such results on a variety of hardware is available in the `bench/` folder.

There may be places where compiler optimizations render some benchmarks moot. One known such instance is in the benchmarks that call the `sum` functions with 1, 2 and 3 arguments - the function gets inlined and the result is the timing of the body of the function, not the overhead of the function call, as desired. Any help as to how to avoid inlining while keeping the body of the callee simple is welcome.

## Example

Here is a short example of benchmark results, run on a Toshiba Chromebook 64-bit running Ubuntu (LXDE) with crouton:

```
BenchmarkConvIntToFloat64	1000000000	         2.16 ns/op	       0 B/op	       0 allocs/op
BenchmarkConvFloat64ToInt	2000000000	         1.26 ns/op	       0 B/op	       0 allocs/op
BenchmarkConvIntToFloat32	2000000000	         1.44 ns/op	       0 B/op	       0 allocs/op
BenchmarkConvFloat32ToInt	2000000000	         1.42 ns/op	       0 B/op	       0 allocs/op
BenchmarkConvStringToInt	30000000	        57.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkConvIntToString	20000000	       108 ns/op	       1 B/op	       1 allocs/op
BenchmarkStringToByte	20000000	       110 ns/op	       8 B/op	       1 allocs/op
BenchmarkByteToString	20000000	        84.8 ns/op	       4 B/op	       1 allocs/op
BenchmarkFindSliceInt5	100000000	        13.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkFindSliceStr5	20000000	        64.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkFindSliceInt100	10000000	       236 ns/op	       0 B/op	       0 allocs/op
BenchmarkFindSliceStr100	 1000000	      1248 ns/op	       0 B/op	       0 allocs/op
BenchmarkFindMapInt5	100000000	        22.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkFindMapStr5	20000000	        77.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkFindMapInt100	30000000	        58.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkFindMapStr100	20000000	        76.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkFloat32Assign	2000000000	         0.99 ns/op	       0 B/op	       0 allocs/op
```

## Installation

    $ go get github.com/PuerkitoBio/gocostmodel

## Running the benchmarks

    $ go test -run zz -bench . [-benchmem]

## Contributing

Pull requests are welcome, but the benchmarks should be for common basic operations, in the spirit of the existing benchmarks.

## License

This source code is released under the [BSD 3-clause license][bsd].

[bsd]: http://opensource.org/licenses/BSD-3-Clause

