package main

import (
	"testing"
	"reflect"
	"runtime"
	"math/bits"
)


// A test sample with an average of 7 set bits, and an average length of 11 bits
var tests = []struct {
	input uint
	want  uint
}{
	{5, 2},
	{0, 0},
	{8, 1},
	{15, 4},
	{19, 3},
	{ 1<<20-1, 20},
	{ 1<<20, 1},
	{ 1<<19-1, 19},
}


func population_count_mathbits(n uint) uint {
	return uint(bits.OnesCount(n))
}

func Test_popcount_variants(t *testing.T) {

	funcs := []func(uint) uint{
		population_count,
		population_count_kernighan,
		population_count_table,
		population_count_divideconquer,
		population_count_mathbits,
	}

	for _, f:= range funcs {
		function_name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
		t.Run(function_name, func (t *testing.T) {
			for _, test := range tests {
				got := f(test.input)
				if got != test.want {
					t.Errorf("for input: %d, want output: %d, got: %d", test.input, test.want, got)
				}
			}			
		})
	}

}

var result uint 


func Benchmark_population_count(b *testing.B) {
	for i:=0; i<b.N; i++ {
		for _, test := range tests {
			r := population_count(test.input)
			result+=r
		}			
	}
}

func Benchmark_population_count_kernighan(b *testing.B) {
	for i:=0; i<b.N; i++ {
		for _, test := range tests {
			r := population_count_kernighan(test.input)
			result+=r
		}			
	}
}

func Benchmark_population_count_table(b *testing.B) {
	for i:=0; i<b.N; i++ {
		for _, test := range tests {
			r := population_count_table(test.input)
			result+=r
		}			
	}
}

func Benchmark_population_count_divideconquer(b *testing.B) {
	for i:=0; i<b.N; i++ {
		for _, test := range tests {
			r := population_count_divideconquer(test.input)
			result+=r
		}			
	}
}

func Benchmark_population_count_mathbits(b *testing.B) {
	for i:=0; i<b.N; i++ {
		for _, test := range tests {
			r := bits.OnesCount(test.input)
			result+= uint(r)
		}			
	}
}



//// Benchmark all of the functions 
//// Commented out and split into individual benchmarks to improve performance
// func Benchmark_popcount_variants(b *testing.B) {
//     funcs := []func(uint) uint{
//             population_count,
//             population_count_kernighan,
//             population_count_table,
//             population_count_divideconquer,
//             population_count_mathbits,
//     }

//     for _, f:= range funcs {
//             function_name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
//             b.Run(function_name, func (b *testing.B) {
//                     for i:=0; i<b.N; i++ {
//                             for _, test := range tests {
//                                     _ = f(test.input)
//                             }                       
//                     }
//             })
//     }
// }
