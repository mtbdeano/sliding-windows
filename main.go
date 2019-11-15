package main

import (
	"github.com/mtbdeano/sliding-windows/binheap"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
	binheap.Sort([]string{"b", "c"})
}
