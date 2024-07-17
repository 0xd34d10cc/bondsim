package sim

import (
	"bondsim/internal/domain"
	"fmt"
	"io"
	"time"
)

type Sim interface {
	// task management
	Scheduler

	// portfolio
	Balance() Balance

	// environment
	MarketYield() domain.Yield
	Now() time.Time
}

type SimRunner interface {
	Run(until time.Time)
	PrintReport(w io.Writer) error
	Sim
}

type Config struct {
	Name          string
	Curve         domain.YieldCurve
	YieldDiscount domain.Yield
}

type sim struct {
	sched   *scheduler
	balance *balance
	conf    *Config
	now     time.Time
}

func New(start time.Time, config Config) SimRunner {
	return &sim{
		sched:   newScheduler(),
		balance: newBalance(),
		conf:    &config,
		now:     start,
	}
}

func (s *sim) Run(end time.Time) {
	for s.now.Before(end) {
		s.sched.RunUntil(s.now)
		// TODO: skip directly to Min(r.s.NextTime(), end)
		s.now = s.now.AddDate(0, 0, 1)
	}
}

func (s *sim) PrintReport(w io.Writer) error {
	fmt.Fprintf(w, "=== %v ===\n", s.conf.Name)
	fmt.Fprintf(w, "Net worth: %v\n", s.balance.NetWorth())
	fmt.Fprintf(w, "Coupons: %v\n", s.balance.Coupons())
	fmt.Fprintf(w, "Yield: %v\n", s.balance.Yield())
	fmt.Fprintf(w, "Investments:\n")
	for yield, invested := range s.balance.investments {
		fmt.Fprintf(w, "\t%v @ %v\n", invested, yield)
	}
	fmt.Fprintf(w, "Investments by source:\n")
	for source, invested := range s.balance.investmentsBySource {
		fmt.Fprintf(w, "\t%v from %v\n", invested, source)
	}
	return nil
}

func (s *sim) Schedule(t time.Time, fn func()) {
	s.sched.Schedule(t, fn)
}

func (s *sim) Balance() Balance {
	return s.balance
}

func (s *sim) Now() time.Time {
	return s.now
}

func (s *sim) MarketYield() domain.Yield {
	curve := s.conf.Curve
	for _, yieldAtTime := range curve {
		if s.now.After(yieldAtTime.T) {
			return yieldAtTime.Y - s.conf.YieldDiscount
		}
	}

	return domain.NeutralYield - s.conf.YieldDiscount
}
