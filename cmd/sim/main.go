package main

import (
	"bondsim/internal/actions"
	"bondsim/internal/domain"
	"bondsim/internal/sim"
	"fmt"
	"os"
	"time"
)

func date(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

var yieldRealistic = domain.YieldCurve{
	{T: date(2026, 12, 1), Y: domain.NeutralYield},
	{T: date(2026, 6, 1), Y: 9_00},
	{T: date(2025, 12, 1), Y: 10_00},
	{T: date(2025, 6, 1), Y: 12_00},
	{T: date(2024, 12, 1), Y: 16_00},
	{T: date(2024, 7, 26), Y: 18_00},
	{T: date(2023, 12, 1), Y: 16_00},
}

type investment struct {
	yield  domain.Yield
	amount domain.Cents
}

type Scenario struct {
	init    investment
	conf    sim.Config
	actions []actions.Action
}

func main() {
	init := investment{
		// yield:  12_00,
		// amount: 5_100_000_00,
	}
	curve := yieldRealistic
	discount := domain.Yield(4_00)
	scenarios := []Scenario{
		{
			init: init,
			conf: sim.Config{
				Name:          "baseline",
				Curve:         curve,
				YieldDiscount: discount,
			},
			actions: []actions.Action{
				actions.Monthly(actions.Salary(350_000_00)),
				actions.Daily(actions.InvestAllCash()),
				actions.InvestCoupons(),
			},
		},
		{
			init: init,
			conf: sim.Config{
				Name:          "loan 3kk",
				Curve:         curve,
				YieldDiscount: discount,
			},
			actions: []actions.Action{
				actions.Monthly(actions.Salary(350_000_00)),
				actions.Daily(actions.InvestAllCash()),
				actions.InvestCoupons(),
				actions.StartAt(date(2024, 8, 1), actions.All(
					actions.InvestAmount(3_000_000_00),
					actions.Monthly(actions.Spend(80_000_00)),
				)),
			},
		},
		{
			init: init,
			conf: sim.Config{
				Name:          "loan 5kk",
				Curve:         curve,
				YieldDiscount: discount,
			},
			actions: []actions.Action{
				actions.Monthly(actions.Salary(350_000_00)),
				actions.Daily(actions.InvestAllCash()),
				actions.InvestCoupons(),
				actions.StartAt(date(2024, 8, 1), actions.All(
					actions.InvestAmount(5_000_000_00),
					actions.Monthly(actions.Spend(135_000_00)),
				)),
			},
		},
	}

	from := date(2024, 8, 1)
	to := date(2029, 8, 1)
	for _, s := range scenarios {
		r := sim.New(from, s.conf)
		r.Balance().Investment(s.init.yield, s.init.amount, "InitialInvestment")
		for _, act := range s.actions {
			r.Schedule(r.Now(), act(r))
		}
		r.Run(to)

		err := r.PrintReport(os.Stdout)
		if err != nil {
			fmt.Println("Report:", err)
		}
	}
}
