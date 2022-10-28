package scrape

import (
	"context"

	"github.com/google/pprof/profile"
)

type Storage interface {
	SaveProfile(ctx context.Context, prof *profile.Profile) (err error)
	GetProfiles(ctx context.Context) (profiles map[string]*profile.Profile)
	GetProfile(ctx context.Context, id string) (prof *profile.Profile)
}

type memStorage struct {
	mem map[string]*profile.Profile
}

func NewStorage() Storage {
	return &memStorage{mem: make(map[string]*profile.Profile)}
}

func (m *memStorage) SaveProfile(ctx context.Context, prof *profile.Profile) (err error) {
	return
}

func (m *memStorage) GetProfiles(ctx context.Context) (profiles map[string]*profile.Profile) {
	return
}

func (m *memStorage) GetProfile(ctx context.Context, id string) (prof *profile.Profile) {
	return
}
