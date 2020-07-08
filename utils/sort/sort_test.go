package utils

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {

	elements := []int{9,7,5,3,1,2,4,6,8,0}

	BubbleSort(elements)

	if elements[0] != 9 {
		t.Error("first element should be 9")
	}

	if elements[len(elements)-1] != 0 {
		t.Error("last element should be 0")
	}
}