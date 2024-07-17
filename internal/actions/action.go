package actions

import "bondsim/internal/sim"

type Action = func(sim.Sim) func()
