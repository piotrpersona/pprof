package scrape

import (
	"context"
	"sync"
	"time"

	"github.com/google/pprof/profile"
	"github.com/google/uuid"
)

type ProfileDump struct {
	prof        *profile.Profile
	Timestamp   time.Time `json:"timestamp"`
	TimestampNs uint64    `json:"timestampNs"`
	ID          string    `json:"id"`
}

func NewProfileDump(prof *profile.Profile) *ProfileDump {
	now := time.Now().UTC()
	return &ProfileDump{
		prof:        prof,
		Timestamp:   now,
		TimestampNs: uint64(now.UnixNano()),
		ID:          uuid.NewString(),
	}
}

func (p *ProfileDump) Prof() *profile.Profile {
	return p.prof
}

type Storage interface {
	SaveProfile(ctx context.Context, prof *ProfileDump) (id string, err error)
	GetProfiles(ctx context.Context) (profiles map[string]*ProfileDump)
	GetProfile(ctx context.Context, id string) (prof *ProfileDump)
}

type memStorage struct {
	sync.RWMutex
	mem map[string]*ProfileDump
}

func NewStorage() Storage {
	return &memStorage{mem: make(map[string]*ProfileDump)}
}

func (m *memStorage) SaveProfile(ctx context.Context, prof *ProfileDump) (id string, err error) {
	m.Lock()
	defer m.Unlock()

	id = prof.ID
	m.mem[id] = prof
	return
}

func (m *memStorage) GetProfiles(ctx context.Context) (profiles map[string]*ProfileDump) {
	m.RLock()
	defer m.RUnlock()

	profiles = m.mem
	return
}

func (m *memStorage) GetProfile(ctx context.Context, id string) (prof *ProfileDump) {
	m.RLock()
	defer m.RUnlock()

	prof, ok := m.mem[id]
	if !ok {
		return
	}
	return
}
