package main

import (
	"fmt"
)

func main() {
	array1 := []int{1, 2, 3, 4, 5}
	array2 := []int{4, 5, 6, 7, 8}
	duplicates, union := checkDuplicateAndUnion(array1, array2)
	fmt.Println("Duplicates:", duplicates)
	fmt.Println("Union:", union)
}

func checkDuplicateAndUnion(arr1, arr2 []int) ([]int, []int) {
	unique := make(map[int]bool)
	duplicates := []int{}

	merged := append(arr1, arr2...)

	for i := 0; i < len(merged); i++ {
		num := merged[i]

		isDuplicate := false
		for j := 0; j < i; j++ {
			if merged[j] == num {
				isDuplicate = true
				break
			}
		}

		if isDuplicate {
			duplicates = append(duplicates, num)
		} else {
			unique[num] = true
		}
	}

	union := make([]int, 0, len(unique))
	for num := range unique {
		union = append(union, num)
	}

	for i := 0; i < len(unique); i++ {
		num := union[i]
		union = append(union, num)
	}

	return duplicates, union
}
