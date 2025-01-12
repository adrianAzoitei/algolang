package mergesort

func MergeSort(A []int) (sorted []int, err error) {
	// assume len(A) is even for now
	n := len(A)

	// base cases
	if n == 1 || n == 0 {
		return A, nil
	}

	C, _ := MergeSort(A[0 : n/2])
	D, _ := MergeSort(A[n/2 : n])
	return Merge(C, D)
}

func Merge(C, D []int) (B []int, err error) {
	n1 := len(C)
	n2 := len(D)
	n := n1 + n2
	B = make([]int, n)

	i := 0
	j := 0
	for k := 0; k < n; k++ {
		if i < n1 && j < n2 {
			if C[i] < D[j] {
				B[k] = C[i]
				i++
			} else {
				B[k] = D[j]
				j++
			}
			continue
		}

		for i < n1 {
			B[k] = C[i]
			i++
			break
		}

		for j < n2 {
			B[k] = D[j]
			j++
			break
		}
	}

	return B, nil
}
