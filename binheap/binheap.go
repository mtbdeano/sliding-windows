package binheap

import (
	"math"
	"strings"
	"fmt"

	log "github.com/sirupsen/logrus"
)

// Heap is a siple min or max heap binary tree array
type Heap struct {
	heap []int // use the slice length and capacity to store next position and total size
}

// MaxHeap is a heap where the max is the root
type MaxHeap Heap

// Left provides the left child
func (h Heap) Left(i int) int {
	return i*2 + 1
}

// Right provides the right child
func (h Heap) Right(i int) int {
	return i*2 + 2
}

// Parent provides the parent link
func (h Heap) Parent(i int) int {
	return int(math.Floor(float64(i) / 2.0))
}

// Root gets the root of the heap (min or max)
func (h Heap) Root() int {
	return h.heap[0]
}

func (h Heap) printTree(pos int, depth int) string {
	if pos >= len(h.heap) {
		return ""
	}
	s := h.printTree(h.Left(pos), depth+1)
	s += strings.Repeat("    ", depth) + fmt.Sprintf("%4d", h.heap[pos]) + "\n"
	s += h.printTree(h.Right(pos), depth+1)
	return s
}

// String stringifies the tree
func (h Heap) String() string {
	return h.printTree(0, 0)
}

// CompareFunc is a comparitor
type CompareFunc func(x, y int) bool

func (h *Heap) heapify(pos int, f CompareFunc) {
	log.Debugf("heapify(%d)", pos)
	if pos >= 0 && pos < len(h.heap) {
		p := h.Parent(pos)
		log.Debugf("heapify: pos %d, parent %d", pos, p)
		if p != pos {
			l := h.Left(p)
			r := h.Right(p)
			log.Debugf("heapify: pos %d, parent %d, left %d right %d len %d", pos, p, l, r, len(h.heap))

			if l < len(h.heap) && f(h.heap[p], h.heap[l]) {
				log.Debugf("at pos %d[%d:%d] with parent %d, left %d", p, l, r, h.heap[p], h.heap[l])

				h.heap[p], h.heap[l] = h.heap[l], h.heap[p]
				log.Debugf("Left Swapping [%d]:%d with [%d]:%d recursing up to %d", p, h.heap[p], l, h.heap[l], l)
				h.heapify(l, f)
			}
			if r < len(h.heap) && f(h.heap[p], h.heap[r]) {
				log.Debugf("at pos %d[%d:%d] with parent %d, left %d right %d", p, l, r, h.heap[p], h.heap[l], h.heap[r])

				h.heap[p], h.heap[r] = h.heap[r], h.heap[p]
				log.Debugf("Right Swapping [%d]:%d with [%d]:%d recursing up to %d", p, h.heap[p], r, h.heap[r], r)
				h.heapify(r, f)
			}
			h.heapify(p, f)
		} else {
			log.Debugf("Stopping parent and pos are the same!")
		}
	} else {
		log.Debugf("Stopping at the root!")
	}
}

// MinHeap is the minimum heap
type MinHeap Heap

// String pretty prints a tree
func (h MinHeap) String() string {
	return (Heap)(h).String()
}

// Min returns the minimum from the min heap
func (h MinHeap) Min() int {
	return Heap(h).Root()
}

// Insert adds this integer to the minheap
func (h *MinHeap) Insert(x int) {
	h.heap = append(h.heap, x)
	log.Debugf("Insert(%d)", x)
	(*Heap)(h).heapify(len(h.heap)-1, func(x, y int) bool {
		return x >= y
	})
}

// NewMinHeap creates a new min heap
func NewMinHeap(size int) (h *MinHeap) {
	var nh MinHeap
	nh.heap = make([]int, 0, size)
	log.Debugf("Created new heap with max size %d", size)
	return &nh
}
