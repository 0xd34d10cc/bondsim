package actions

import (
	"bondsim/internal/domain"
	"bondsim/internal/sim"
)

func InvestCoupons() Action {
	return func(s sim.Sim) func() {
		return func() {
			investments := s.Balance().Investments()
			coupons := float64(0)
			for yield, invested := range investments {
				percents := float64(yield) / 100
				y := (percents / 100) / 2 // assuming semi-yearly coupons
				coupons += y * float64(invested)
			}

			s.Balance().Investment(s.MarketYield(), domain.Cents(coupons), "InvestCoupons")
			next := s.Now().AddDate(0, 6, 0)
			s.Schedule(next, InvestCoupons()(s))
		}
	}
}
