package util

import "time"

//go:generate mockgen -destination=../mocks/mock_clock.go -package=mocks nagase/util Clock

type Clock interface {
	Now() time.Time
}

type RealClock struct{}

func (RealClock) Now() time.Time { return time.Now() }
