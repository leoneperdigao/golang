package countinversions_test

import (
	"testing"

	"../countinversions"
)

const EXPECTED = 87

func TestBruteForce(t *testing.T) {
	array1 := []int{6, 10, 13, 5, 8, 2, 11, 15, 4, 25, 21, 5, 73, 1, 46, 3, 6, 31, 5, 7}
	array2 := make([]int, len(array1))
	copy(array2, array1)
	count := countinversions.BruteForce(array2)

	if count != EXPECTED {
		t.Errorf("Total of inverions was incorrect, got: %d, want: %d.", count, EXPECTED)
	}
}

func TestCount(t *testing.T) {
	array1 := []int{6, 10, 13, 5, 8, 2, 11, 15, 4, 25, 21, 5, 73, 1, 46, 3, 6, 31, 5, 7}
	array2 := make([]int, len(array1))
	copy(array2, array1)
	count := countinversions.Count(array1)
	if count != EXPECTED {
		t.Errorf("Total of inverions was incorrect, got: %d, want: %d.", count, EXPECTED)
	}
}
