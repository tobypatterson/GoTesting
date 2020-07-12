package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortIncreasingOrder(t *testing.T) {

	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}

	BubbleSort(elements)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 0, elements[0])
	assert.EqualValues(t, 9, elements[len(elements)-1])

}

func TestSortIncreasingOrder(t *testing.T) {

	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}

	Sort(elements)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 0, elements[0])
	assert.EqualValues(t, 9, elements[len(elements)-1])
}

func BenchmarkBubbleSort(b *testing.B) {

	elements := GenElements(100)

	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkSort(b *testing.B) {

	elements := GenElements(100)

	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}

func TestWithTimeout(t *testing.T) {
	elements := GenElements(10)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 9, elements[0])
	assert.EqualValues(t, 0, elements[len(elements)-1])

	timeoutChan := make(chan bool, 1)

	go func() {
		BubbleSort(elements)
		timeoutChan <- false
	}()

	go func() {
		time.Sleep(500 * time.Millisecond)
		timeoutChan <- true
	}()

	if <-timeoutChan {
		assert.Fail(t, "Function Timed Out")
		return
	}

	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 0, elements[0])
	assert.EqualValues(t, 9, elements[len(elements)-1])
}
