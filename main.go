// Package simplesort implements the algorithms described here:
// https://arxiv.org/pdf/2110.01111.pdf
//
// Also see discussion here:
// https://news.ycombinator.com/item?id=28758106
//
package simplesort

func Algorithm1(A []int) {
	n := len(A)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if A[i] < A[j] {
				A[i], A[j] = A[j], A[i]
			}
		}
	}
}

func Algorithm2(A []int) {
	n := len(A)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if A[i] > A[j] {
				A[i], A[j] = A[j], A[i]
			}
		}
	}
}

func Algorithm3(A []int) {
	n := len(A)
	for i := 0; i < n; i++ {
		for j := 1; j < i-1; j++ {
			if A[i] < A[j] {
				A[i], A[j] = A[j], A[i]
			}
		}
	}
}
