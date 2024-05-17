package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {

	tcs := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"Add 1 and 2", 1, 2, 3},
		{"Add 2 and 2", 2, 2, 4},
		{"Add 0 and 0", 0, 0, 0},
		{"Add -1 and 1", -1, 1, 0},
		{"Add -1 and -1", -1, -1, -2},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := Add(tc.a, tc.b)
			if actual != tc.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tc.a, tc.b, actual, tc.expected)
			}
		})
	}

}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
