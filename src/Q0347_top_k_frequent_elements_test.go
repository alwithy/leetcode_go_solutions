package src

import (
	"fmt"
	"testing"
)

func TestQ0347(t *testing.T) {
	nums := []int{1,1,1,2,2,3}
	k := 2
	res := topKFrequent(nums, k)
	fmt.Println(res)
}
