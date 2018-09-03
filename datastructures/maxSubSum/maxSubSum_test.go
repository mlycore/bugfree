package maxSubSum
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_MaxSubSum_1(t *testing.T) {
	list := []int{2, -1, 3, 4, -5, 6, 7}
	r := maxSubSum_1(list)
	assert.Equal(t, 16, r, "should be equal")
}

func Test_MaxSubSum_2(t *testing.T) {
	// list := []int{2, -1, 3, 4, -5, 6, 7}
	list := []int{-2, 3, 3, 4, -5, -6, -7}
	r := maxSubSum_2(list)
	// assert.Equal(t, 16, r, "should be equal")
	assert.Equal(t, 10, r, "should be equal")
}

func Test_MaxSubSum_3(t *testing.T) {
	list := []int{-2, 3, 3, 4, -5, -6, -7}
	r := maxSubSum_3(list)
	assert.Equal(t, 10, r, "should be equal")
}
