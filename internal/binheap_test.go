package binheap

import "testing"

func TestSortOne(t *testing.T) {
	binheap.Sort([]string{"b", "c", "a", "d"})

}
