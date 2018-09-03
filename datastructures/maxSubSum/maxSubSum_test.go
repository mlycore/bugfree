package maxSubSum
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_MaxSubSum(t *testing.T) {
	list := []int{2, -1, 3, 4, -5, 6, 7}
	r := maxSubSum(list)
	assert.Equal(t, 16, r, "should be equal")
}
