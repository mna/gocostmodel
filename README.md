# gocostmodel

This package was inspired by Brian W. Kernighan and Rob Pike's book "The Practice of Programming" (Addison-Wesley, 1998). In Chapter 7 on performance, section 7.6 (Estimation):

> "It's hard to estimate ahead of time how fast a program will be [...]. It's easy, though, to create a *cost model* for a language or a system, which will give you at least a rough idea of how long important operations take."

This is a *cost model* that benchmarks common basic operations for the Go language. For it to be meaningful, it should be run on your target hardware, and the relative performance of operations is probably more important than the actual numbers.

There may be places where compiler optimizations render some benchmarks moot. One known such instance is in the benchmarks that call the `sum` functions with 1, 2 and 3 arguments - the function gets inlined and the result is the timing of the body of the function, not the overhead of the function call, as desired. Any help as to how to avoid inlining while keeping the body of the callee simple is welcome.

## Installation

    $ go get github.com/PuerkitoBio/gocostmodel

## Running the benchmarks

    $ go test -run zz -bench . [-benchmem]

## Contributing

Pull requests are welcome, but the benchmarks should be for common basic operations, in the spirit of the existing benchmarks.

## License

This source code is released under the [BSD 3-clause license][bsd].

[bsd]: http://opensource.org/licenses/BSD-3-Clause
