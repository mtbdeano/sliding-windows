package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/mtbdeano/sliding-windows"
	"github.com/mtbdeano/sliding-windows/binheap"
)

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}
