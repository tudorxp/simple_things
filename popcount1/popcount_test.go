package main

import (
	"testing"
	"reflect"
	"runtime"
)

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


func Test_population_count(t *testing.T) {

	funcs := []func(uint) uint{
		population_count_1,
		population_count_2,
		population_count_3,
		population_count_4,

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


func Benchmark_population_count(b *testing.B) {

	funcs := []func(uint) uint{
		population_count_1,
		population_count_2,
		population_count_3,
		population_count_4,

	}

	for _, f:= range funcs {
		function_name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
		b.Run(function_name, func (b *testing.B) {
			for i:=0; i<b.N; i++ {
				for _, test := range tests {
					_ = f(test.input)
				}			
			}
		})
	}

}
