package actions

import "bondsim/internal/sim"

func Daily(act Action) Action {
	return func(s sim.Sim) func() {
		inner := act(s)
		return func() {
			inner()
			next := s.Now().AddDate(0, 0, 1)
			s.Schedule(next, Monthly(act)(s))
		}
	}
}
