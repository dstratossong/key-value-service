package model

type Run struct {
	Id         uint64
	Service    *Service
	Status     string
	Properties *Properties
	Result     *Properties
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

func FinishRun(id uint64, status string, data *Properties) {
	run, found := RunStore[id]
	if !found {
		// Muted Error
		return
	}

	run.Status = status
	run.Result = data
}
