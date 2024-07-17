package sim

import "bondsim/internal/domain"

type Balance interface {
	Cash(change domain.Cents, source string) domain.Cents
	Investment(yield domain.Yield, change domain.Cents, source string) domain.Cents
	Investments() map[domain.Yield]domain.Cents
}

type balance struct {
	cash                domain.Cents
	investments         map[domain.Yield]domain.Cents
	investmentsBySource map[string]domain.Cents
}

func newBalance() *balance {
	return &balance{
		cash:                0,
		investments:         map[domain.Yield]domain.Cents{},
		investmentsBySource: map[string]domain.Cents{},
	}
}

func (b *balance) Cash(change domain.Cents, source string) domain.Cents {
	b.cash += change
	return b.cash
}

func (b *balance) Investment(yield domain.Yield, change domain.Cents, source string) domain.Cents {
	if change != 0 {
		b.investments[yield] += change
		b.investmentsBySource[source] += change
	}
	return b.investments[yield]
}

func (b *balance) Investments() map[domain.Yield]domain.Cents {
	return b.investments
}

func (b *balance) NetWorth() domain.Cents {
	s := b.cash
	for _, invested := range b.investments {
		s += invested
	}
	return s
}

func (b *balance) Coupons() domain.Cents {
	c := float64(0)
	for yield, invested := range b.investments {
		percents := float64(yield) / 100
		c += (percents / 100) * float64(invested)
	}
	return domain.Cents(c)
}

func (b *balance) Yield() domain.Yield {
	netWorth := b.NetWorth()
	coupons := b.Coupons()
	yield := float64(coupons) / float64(netWorth)
	percents := yield * 100
	return domain.Yield(percents * 100)
}
