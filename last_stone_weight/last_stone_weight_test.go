package last_stone_weight

import (
	"testing"
)

func TestSolution(t *testing.T) {
	var ans, expectAns int32
	
	ans = lastStoneWeight([]int32{1, 2, 3, 6, 7, 7})
	expectAns = 0
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}
	
	ans = lastStoneWeight([]int32{2, 4, 5})
	expectAns = 1
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}
}
