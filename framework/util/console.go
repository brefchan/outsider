package util

import (
	"fmt"
)

// PrettyPrint 美观输出数组
func PrettyPrint(arr [][]string) {
	if len(arr) == 0 {
		return
	}
	rows := len(arr)
	cols := len(arr[0])

	lensMap := make(map[int]int)
	for j := 0; j < cols; j++ {
		for i := 0; i < rows; i++ {
			lens := len(arr[i][j])
			if lens > lensMap[j] {
				lensMap[j] = lens
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Print(arr[i][j])
			padding := lensMap[j] - len(arr[i][j]) + 2
			for p := 0; p < padding; p++ {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
