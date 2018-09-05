package list

import (
	"fmt"
)

func pow(a, b int) int {
	p := 1
	if b == 0 {
		return 1
	}
	for i := 1; i < b+1; i++ {
		p = p * a
	}
	return p
}

func BucketSort(list []int) []int {
	// 0 - 9
	buckets := make([][]int, 10)
	for i := 0; i < 10; i++ {
		buckets[i] = make([]int, 0)
	}
	loc := make(map[int]int)

	for i := 0; i < 10; i++ {
		var flag bool
		for j := 0; j < len(list); j++ {
			if list[j] == 0 {
				continue
			}
			t := list[j] / pow(10, i)
			rem := t % 10
			if i > 0 && loc[list[j]] != rem {
				buckets[loc[list[j]]] = buckets[loc[list[j]]][0 : len(buckets[loc[list[j]]])-1]

			}
			if loc[list[j]] == rem {
				continue
			}
			loc[list[j]] = rem
			buckets[rem] = append(buckets[rem], list[j])

			if t == 0 {
				list[j] = 0
				flag = true
			}
			fmt.Printf("i: %d, t: %d, rem: %d, buckets: %v\n", i, t, rem, buckets)
		}
		if flag {
			break	
		}
	}

	re := make([]int, 0)
	for i:=0; i<len(buckets); i++ {
		for j:=0; j<len(buckets[i]); j++ {
			re = append(re, buckets[i][j])
		}
	}
	// return buckets[0]
	return re
}
