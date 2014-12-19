package main

var (
	i1 int
	i2 int = 10
)

// skeleton program to inspect things like generated assembly, compiler
// options impact on benchmarks, etc.
func main() {
	for i := 0; i < 1000; i++ {
		i1 = i2
	}
}
