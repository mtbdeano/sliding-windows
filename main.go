package main

import (
	"math/rand"
	"fmt"

	"github.com/mtbdeano/sliding-windows/binheap"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
	h := binheap.NewMinHeap(5)
	h.Insert(4)
	fmt.Println(h.String())

	h.Insert(1)
	fmt.Println(h.String())
	h.Insert(10)
	fmt.Println(h.String())
	h.Insert(5)
	fmt.Println(h.String())
	h.Insert(2)
	fmt.Println(h.String())
	for i := 0; i < 10; i++ {
		h.Insert(rand.Intn(100))
	}
	fmt.Println(h.String())
}
