
When I first set eyes on the problem, the naive solution immediately presented itself - 
checking the least significant bit of the integer, then shifting right, and looping over these two steps until we reach 0.

Having since done further research into the matter, I've implemented a number of alternative methods; 
However I will still submit the original solution as the default, since 
- I find it the most readable 
- In the absence of performance requirements, I judge it "good enough" for the purpose - since our integers have less than 20 bits, it loops at most 20 times.

Here is a sample run, and a test/benchmark run of all methods as tested on my machine (have included some tests with larger integers & larger bit counts to emphasise the difference between approaches)

[top@tmac popcount]$ go run popcount.go 
population_count(5) = 2
population_count(0) = 0
population_count(8) = 1
population_count(15) = 4
population_count(19) = 3

[top@tmac popcount]$ go test -bench .
goos: darwin
goarch: amd64
pkg: tudorxp/popcount
Benchmark_population_count-8                 	30000000	        54.1 ns/op
Benchmark_population_count_kernighan-8       	30000000	        47.3 ns/op
Benchmark_population_count_table-8           	50000000	        31.4 ns/op
Benchmark_population_count_divideconquer-8   	100000000	        14.8 ns/op
Benchmark_population_count_mathbits-8        	100000000	        14.1 ns/op
PASS
ok  	tudorxp/popcount	7.678s

If performance is really needed, it turns out that calling the Go standard library function in math/bits (bits.OnesCount) is the fastest.
However, that's not because their implementation (in https://golang.org/src/math/bits/bits.go) is faster, as most of the time that code is not even used.
It turns out the Go compiler will detect an invocation of these math/bits functions ( as seen in: https://github.com/golang/go/blob/master/src/cmd/compile/internal/gc/ssa.go#L3549-L3580 ).
Depending on architecture and CPU capabilities, they will more often than not be optimised to an (x86_64 in my case: POPCNTQ) assembler instruction.
It is likely this can be optimised further, but perhaps Go is not the best language for that then.