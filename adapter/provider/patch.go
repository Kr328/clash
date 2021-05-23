package provider

import "time"

var (
	suspended bool
)

type UpdatableProvider interface {
	UpdatedAt() time.Time
}

func (f *fetcher) UpdatedAt() time.Time {
	return f.updatedAt
}

func Suspend(s bool) {
	suspended = s
}
