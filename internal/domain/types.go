package domain

import (
	"fmt"
	"strconv"
	"time"
)

type Cents int64

func (m Cents) String() string {
	v := float64(m / 100)
	if v > 1_000 {
		return fmt.Sprintf("%.1fk", v/1_000)
	}

	return strconv.FormatFloat(v, 'f', 1, 64)
}

type Yield int // percent points

func (y Yield) String() string {
	return strconv.FormatFloat(float64(y)/100, 'f', 1, 64) + "%"
}

const NeutralYield = Yield(8_00)

type YieldAtTime struct {
	T time.Time
	Y Yield
}

type YieldCurve = []YieldAtTime
