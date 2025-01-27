package service

import (
	"errors"

	"github.com/baking-bad/bcdhub/internal/models"
	"github.com/baking-bad/bcdhub/internal/models/service"
	"github.com/baking-bad/bcdhub/internal/postgres/core"
	"github.com/go-pg/pg/v10"
)

// Storage -
type Storage struct {
	*core.Postgres
}

// NewStorage -
func NewStorage(pg *core.Postgres) *Storage {
	return &Storage{pg}
}

// Get -
func (s *Storage) Get(name string) (state service.State, err error) {
	err = s.DB.Model(&state).Where("name = ?", name).First()
	if errors.Is(err, pg.ErrNoRows) {
		err = nil
		state.Name = name
	}
	return
}

// Save -
func (s *Storage) Save(state service.State) error {
	return state.Save(s.DB)
}

// Get -
func (s *Storage) All() (state []service.State, err error) {
	err = s.DB.Model().Table(models.DocServices).Select(&state)
	return
}
