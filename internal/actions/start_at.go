package actions

import (
	"bondsim/internal/sim"
	"time"
)

func StartAt(at time.Time, act Action) Action {
	return func(s sim.Sim) func() {
		return func() {
			s.Schedule(at, act(s))
		}
	}
}
