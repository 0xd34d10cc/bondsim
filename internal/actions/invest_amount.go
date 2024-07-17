package actions

import (
	"bondsim/internal/domain"
	"bondsim/internal/sim"
)

func InvestAmount(amount domain.Cents) Action {
	return func(s sim.Sim) func() {
		return func() {
			s.Balance().Investment(s.MarketYield(), amount, "Loan")
		}
	}
}
