package arraysslices

import (
	"fmt"
	"slices"
	"testing"
)

func TestReduce(t *testing.T) {
	tcs := []struct {
		name    string
		items   []int
		reducer func(int, int) int
		initial int
		total   int
	}{
		{
			"reduce ints",
			[]int{1, 2, 3, 4, 5},
			func(a, b int) int {
				return a + b
			},
			0,
			15,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			act := Reduce(tc.items, tc.reducer, tc.initial)
			if act != tc.total {
				t.Errorf("got %d expected %d given %v", act, tc.total, tc.items)
			}
		})
	}
}

func TestReduceStrings(t *testing.T) {
	tcs := []struct {
		name    string
		items   []string
		reducer func(string, string) string
		initial string
		total   string
	}{
		{
			"reduce ints",
			[]string{"h", "e", "l", "l", "o"},
			func(a, b string) string {
				return a + b
			},
			"",
			"hello",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			act := Reduce(tc.items, tc.reducer, tc.initial)
			if act != tc.total {
				t.Errorf("got %s expected %s given %v", act, tc.total, tc.items)
			}
		})
	}
}

func TestSum(t *testing.T) {
	tcs := []struct {
		name  string
		nums  []int
		total int
	}{
		{"Test 1", []int{1, 2, 3, 4, 5}, 15},
		{"Test 2", []int{5, 8, 9000, 3, 0}, 9016},
		{"With Negative", []int{-5, -8, 13, 3, -3}, 0},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			act := Sum(tc.nums)
			if act != tc.total {
				t.Errorf("got %d expected %d given %v", act, tc.total, tc.nums)
			}
		})
	}
}

func TestSumAll(t *testing.T) {
	tcs := []struct {
		name   string
		nums   [][]int
		totals []int
	}{
		{"Test 1", [][]int{{1, 2, 3}, {4, 5, 6}}, []int{6, 15}},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			act := SumAll(tc.nums)
			if !slices.Equal(tc.totals, act) {
				t.Errorf("got %d expected %d given %v", act, tc.totals, tc.nums)
			}
		})
	}
}

func TestSumVariadic(t *testing.T) {
	tcs := []struct {
		name   string
		nums   [][]int
		totals []int
	}{
		{"Test 1", [][]int{{1, 2, 3}, {4, 5, 6}}, []int{6, 15}},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			act := SumVariadic(tc.nums...)
			if !slices.Equal(tc.totals, act) {
				t.Errorf("got %d expected %d given %v", act, tc.totals, tc.nums)
			}
		})
	}
}

func TestSumTails(t *testing.T) {
	tcs := []struct {
		name   string
		nums   [][]int
		totals []int
	}{
		{"Test 1", [][]int{{1, 2, 3}, {4, 5, 6}}, []int{5, 11}},
		{"Empty slice", [][]int{}, []int{}},
		{"Empty slice in nums", [][]int{{}, {2, 3, 4, 5}}, []int{0, 12}},
		{"Single num slice in nums", [][]int{{5}, {2, 3, 4, 5}}, []int{0, 12}},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			act := SumTails(tc.nums...)
			if !slices.Equal(tc.totals, act) {
				t.Errorf("got %d expected %d given %v", act, tc.totals, tc.nums)
			}
		})
	}
}

func ExampleSum() {
	sum := Sum([]int{1, 2, 3, 4, 5})
	fmt.Println(sum)
	// Output: 15
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([]int{1, 2, 3, 4, 5})
	}
}

func ExampleSumAll() {
	sum := SumAll([][]int{{1, 2, 3, 4, 5}, {1, 2, 3}})
	fmt.Println(sum)
	// Output: [15 6]
}

func ExampleSumTails() {
	sum := SumTails([]int{1, 2, 3, 4, 5}, []int{1, 2, 3})
	fmt.Println(sum)
	// Output: [14 5]
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([][]int{{1, 2, 3, 4, 5}, {1, 2, 3}})
	}
}

func BenchmarkSumVariadic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumVariadic([]int{1, 2, 3, 4, 5}, []int{1, 2, 3})
	}
}

func BenchmarkSumTails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumTails([]int{1, 2, 3, 4, 5}, []int{1, 2, 3})
	}
}
