package actions

import "bondsim/internal/sim"

func InvestAllCash() Action {
	return func(s sim.Sim) func() {
		return func() {
			amount := s.Balance().Cash(0, "InvestFreeCash")
			if amount > 0 {
				y := s.MarketYield()
				s.Balance().Investment(y, amount, "InvestFreeCash")
				s.Balance().Cash(-amount, "InvestFreeCash")
			}
		}
	}
}
