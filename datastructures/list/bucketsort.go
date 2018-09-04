package list

import (
	"fmt"
)

func pow(a, b int) int {
	p := 1
	for i := 0; i < b; i++ {
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

	for i:=1; i<10; i++ {
		for j:=0; j<len(list); j++ {
			if list[j] == 0 {
				continue
			}
			t := list[j] / pow(10, i)
			rem := list[j] % 10
			buckets[rem] = append(buckets[rem], list[j])
			if t == 0 {
				list[j] = 0		
			} else {
				buckets[rem] = buckets[rem][0:len(buckets[rem]) - 1]
			}
			fmt.Printf("t: %d, rem: %d, buckets: %v\n", t, rem, buckets[rem])
			fmt.Printf("buckets: %v\n", buckets)
		}

	}

	result := make([]int, len(list))
	var k int
	for i := 0; i < len(buckets); i++ {
		for j := 0; j < len(buckets[i]); j++ {
			result[k] = buckets[i][j]
			k++
		}
	}
	return result
}
