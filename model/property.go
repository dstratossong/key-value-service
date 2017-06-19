package model

import (
	"sync/atomic"
)

type Properties map[string]string

var PropertyStore = make(map[uint64]*Properties)

var runCursor uint64 = 0

func NewServiceRun(serviceStr string, props *Properties) uint64 {
	services, found := Services[serviceStr]
	if !found {
		panic("service not found")
	}
	service := services[0]

	id := atomic.AddUint64(&runCursor, 1)

	PropertyStore[id] = props

	run := new(Run)
	run.Id = id
	run.Service = service
	run.Status = "Dispatched"
	run.Properties = props
	RunStore[id] = run

	Enqueue(service, id)

	return id
}

func GetPropertyForRun(id uint64) Properties {
	props, found := PropertyStore[id]
	if !found {
		// ERROR
		panic("id not found")
	}

	return *props
}
