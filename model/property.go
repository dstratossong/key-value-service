package model

import (
	"sync/atomic"
)

type Properties map[string]string

var PropertyStore = make(map[uint64]*Properties)

var runCursor uint64 = 0

func NewServiceRun(service string, props *Properties) {
	id := atomic.AddUint64(&runCursor, 1)
	PropertyStore[id] = props
}

func GetPropertyForRun(id uint64) Properties {
	props, found := PropertyStore[id]
	if !found {
		// ERROR
		panic("id not found")
	}

	return *props
}
