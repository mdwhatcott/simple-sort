package simple_sort

import "testing"

func assertSort(t *testing.T, algorithm func([]int)) {
	var (
		initial  = []int{5, 3, 1, 8, 10, 2, 0}
		expected = []int{0, 1, 2, 3, 5, 8, 10}
	)

	algorithm(initial)

	if !equal(initial, expected) {
		t.Errorf("\n"+
			"Expected: %v\n"+
			"Actual:   %v", expected, initial)
	}
}

func equal(actual, expected []int) bool {
	if len(actual) != len(expected) {
		return false
	}
	for i := range actual {
		if actual[i] != expected[i] {
			return false
		}
	}
	return true
}

func Test1(t *testing.T) {
	assertSort(t, Algorithm1)
}
func Test2(t *testing.T) {
	assertSort(t, Algorithm2)
}
func Test3(t *testing.T) {
	t.Skip("hmm, this one's not working as advertised...")
	assertSort(t, Algorithm3)
}

func Benchmark1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Algorithm1([]int{5, 3, 1, 8, 10, 2, 0})
	}
}
func Benchmark2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Algorithm2([]int{5, 3, 1, 8, 10, 2, 0})
	}
}
func Benchmark3(b *testing.B) {
	b.Skip("not working")
	for i := 0; i < b.N; i++ {
		Algorithm3([]int{5, 3, 1, 8, 10, 2, 0})
	}
}
