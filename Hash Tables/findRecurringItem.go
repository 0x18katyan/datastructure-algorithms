package main

import (
	"fmt"
)

func findDuplicate(arr []int) (int, error) {

	dict := make(map[int]int)

	for _, val := range arr {
		if dict[val] != 0 {
			return val, nil
		} else {
			dict[val] = 1
		}
	}
	return 0, fmt.Errorf("No Duplicates Found.")
}

func main() {
	// arr := []int{1, 2, 3, 8, 2, 3, 1, 0}
	arr := []int{1, 2}
	val, err := findDuplicate(arr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Found Duplicate:", val)
}
