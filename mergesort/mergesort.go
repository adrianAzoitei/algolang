package mergesort

import "errors"

func MergeSort(A []int) (sorted []int, err error) {
	// assume len(A) is even for now
	n := len(A)

	// base cases
	if n == 1 || n == 0 {
		return A, nil
	}

	if n%2 != 0 {
		return nil, errors.New("Expected even length input array")
	}

	C, _ := MergeSort(A[0 : n/2])
	D, _ := MergeSort(A[n/2 : n])
	return Merge(C, D)
}

func Merge(C, D []int) (B []int, err error) {
	// assume len(C) == len(D) for now
	if len(C) != len(D) {
		return nil, errors.New("Expected equally sized input arrays")
	}
	n := len(C)
	B = make([]int, n*2)

	i := 0
	j := 0
	for k := 0; k < n*2; k++ {
		if i < n && j < n {
			if C[i] < D[j] {
				B[k] = C[i]
				i++
			} else {
				B[k] = D[j]
				j++
			}
			continue
		}

		for i < n {
			B[k] = C[i]
			i++
			break
		}

		for j < n {
			B[k] = D[j]
			j++
			break
		}
	}

	return B, nil
}
