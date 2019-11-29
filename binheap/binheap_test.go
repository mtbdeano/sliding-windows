package binheap

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMovingAround(t *testing.T) {
	h := NewMinHeap(8)
	h.Insert(5)
	h.Insert(1)
	h.Insert(3)
	h.Insert(4)
	h.Insert(2)
	h.Insert(6)
	for i := 0; i < 10; i++ {
		h.Insert(rand.Intn(100))
	}
	fmt.Println(h)

}
