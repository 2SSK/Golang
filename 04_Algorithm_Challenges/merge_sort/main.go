package main

import (
	"fmt"
)

func merge(arr []int, l, m, r int) {
	n1 := m - l + 1
	n2 := r - m

	L := make([]int, n1)
	R := make([]int, n2)

	for i := 0; i < n1; i++ {
		L[i] = arr[l+i]
	}

	for i := 0; i < n2; i++ {
		R[i] = arr[m+1+i]
	}

	i, j, k := 0, 0, l
	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		k++
	}

	for i < n1 {
		arr[k] = L[i]
		i++
		k++
	}

	for j < n2 {
		arr[k] = R[j]
		j++
		k++
	}
}

func mergeSort(arr []int, l, r int) {
	if l < r {

		mid := l + (r-l)/2

		mergeSort(arr, l, mid)
		mergeSort(arr, mid+1, r)

		merge(arr, l, mid, r)

	}
}

func main() {
	fmt.Print("Merge Sort : \n\n")
}
