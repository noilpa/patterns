package singleton

import "sync"

var (
	uniqInstance *alone
	once sync.Once
)

type alone struct {
	cntr int
}

func (a *alone) Inc() {
	a.cntr++
}

func (a *alone) Cntr() int {
	return a.cntr
}

func New() *alone {
	once.Do(func() {
		if uniqInstance == nil {
			uniqInstance = new(alone)
		}
	})
	return uniqInstance
}
