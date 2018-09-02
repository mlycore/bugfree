package kth 
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_Kth(t *testing.T) {
	dataset := [][]int{
		{2, 1, 5, 3, 6, 8, 4, 7},
	} 

	k := len(dataset[0]) / 2	
	kth := Kth(k, dataset[0])
	assert.Equal(t, 5, kth, "should be equal")
}
