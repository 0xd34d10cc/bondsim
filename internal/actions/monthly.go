package actions

import "bondsim/internal/sim"

func Monthly(act Action) Action {
	return func(s sim.Sim) func() {
		inner := act(s)
		return func() {
			inner()
			next := s.Now().AddDate(0, 1, 0)
			s.Schedule(next, Monthly(act)(s))
		}
	}
}
