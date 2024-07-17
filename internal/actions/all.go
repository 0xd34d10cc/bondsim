package actions

import "bondsim/internal/sim"

func All(acts ...Action) Action {
	return func(s sim.Sim) func() {
		fns := make([]func(), 0, len(acts))
		for _, act := range acts {
			fns = append(fns, act(s))
		}

		return func() {
			for _, fn := range fns {
				fn()
			}
		}
	}
}
