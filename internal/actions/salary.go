package actions

import (
	"bondsim/internal/domain"
	"bondsim/internal/sim"
)

func Salary(amount domain.Cents) Action {
	return func(sim sim.Sim) func() {
		return func() {
			// add amount to free cash
			sim.Balance().Cash(amount, "Salary")
		}
	}
}
