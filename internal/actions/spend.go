package actions

import (
	"bondsim/internal/domain"
	"bondsim/internal/sim"
)

func Spend(amount domain.Cents) Action {
	return func(s sim.Sim) func() {
		return func() {
			// just spend some money
			s.Balance().Cash(-amount, "LoanPayments")
		}
	}
}
