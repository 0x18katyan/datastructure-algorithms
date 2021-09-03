package main

import "fmt"

func merge_sort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	} else {
		mid := len(arr) / 2
		L := arr[:mid]
		R := arr[mid:]
		return merge(merge_sort(L), merge_sort(R))
	}
}

func merge(L []int, R []int) []int {
	final_arr := []int{}
	i, j := 0, 0
	for i < len(L) && j < len(R) {
		if L[i] < R[j] {
			final_arr = append(final_arr, L[i])
			i++
		} else {
			final_arr = append(final_arr, R[j])
			j++
		}
	}
	for ; i < len(L); i++ {
		final_arr = append(final_arr, L[i])
	}
	for ; j < len(R); j++ {
		final_arr = append(final_arr, R[j])
	}
	return final_arr
}

func main() {
	arr := []int{123, 45, 1, 2, 67, 8, 10, 1, 456, 234, 456, 9}
	fmt.Println("unsorted array ", arr)
	fmt.Println("sorted array", merge_sort(arr))
}
