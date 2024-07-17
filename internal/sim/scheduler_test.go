package sim

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestScheduler(t *testing.T) {
	base := time.Date(2024, 4, 11, 0, 0, 0, 0, time.UTC)
	s := newScheduler()

	events := []string{}
	s.Schedule(base.AddDate(0, 0, 1), func() {
		events = append(events, "1")
		s.Schedule(base.AddDate(0, 0, 3), func() {
			events = append(events, "3")
		})
		s.Schedule(base.AddDate(0, 0, 2), func() {
			events = append(events, "2")
		})
	})
	s.Schedule(base.AddDate(0, 0, 2), func() {
		events = append(events, "2")
	})
	s.Schedule(base.AddDate(0, 1, 0), func() {
		events = append(events, "3")
	})

	s.RunUntil(base)
	require.Equal(t, []string{}, events)

	s.RunUntil(base.AddDate(0, 0, 2))
	require.Equal(t, []string{"1", "2", "2"}, events)

	s.RunUntil(base.AddDate(0, 0, 2))
	require.Equal(t, []string{"1", "2", "2"}, events)

	s.RunUntil(base.AddDate(0, 1, 0))
	require.Equal(t, []string{"1", "2", "2", "3", "3"}, events)
}
