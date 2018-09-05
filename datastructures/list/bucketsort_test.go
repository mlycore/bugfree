package list 
import (
	"fmt"
	"testing"
)

func Test_BucketSort(t *testing.T) {
	list := []int{12, 4, 36, 19, 8, 11}
	fmt.Printf("%v\n", list)
	r := BucketSort(list)
	fmt.Printf("%v\n", r)
}

func Test_BucketSort_1(t *testing.T) {
	list := []int{212, 4, 36, 19, 8, 11}
	fmt.Printf("%v\n", list)
	r := BucketSort(list)
	fmt.Printf("%v\n", r)
}
