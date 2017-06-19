package model

type Run struct {
	Id         uint64
	Service    *Service
	Status     string
	Properties *Properties
}

var RunStore = make(map[uint64]*Run)

func GetRun(id uint64) *Run {
	run, found := RunStore[id]

	if !found {
		// Wrong ID
		return nil
	}

	return run
}
