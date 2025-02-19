package sfu

import (
	"sync"

	"github.com/go-logr/logr"
)

// Logger is an implementation of logr.Logger. If is not provided - will be turned off.
var Logger logr.Logger = logr.Discard()

var (
	packetFactory *sync.Pool
)

func init() {
	// Init packet factory
	packetFactory = &sync.Pool{
		New: func() interface{} {
			b := make([]byte, 1460)
			return &b
		},
	}
}
