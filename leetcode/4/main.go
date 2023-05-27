package main

import "fmt"

func main() {
	res := findMedianSortedArrays([]int{1, 2, 3}, []int{4, 3, 3})

	fmt.Printf("%v\n", res)

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	s_arr, l := sortList(append(nums1, nums2...))

	if l%2 == 0 {
		return (float64(s_arr[l/2]) + float64(s_arr[(l/2)-1])) / 2
	} else {
		return float64(s_arr[(l-1)/2])
	}
}

func sortList(arr []int) ([]int, int) {
	var min_ind int
	l := len(arr)
	for i := 0; i < l; i++ {
		min_ind = i
		for j := i + 1; j < l; j++ {
			if arr[min_ind] > arr[j] {
				min_ind = j
			}
		}
		arr[i], arr[min_ind] = arr[min_ind], arr[i]
	}

	return arr, len(arr)
}
