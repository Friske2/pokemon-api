package domain

import (
	"time"

	"github.com/Friske2/pokemon-api/dto"
)

type Monters struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	CreatedDate time.Time `json:"created_date"`
	CreatedBy   string    `json:"created_by"`
	Enabled     bool      `json:"enabled"`
}

type MontersRepository interface {
	FindAll(m *[]Monters) error
	GetById(id int, m *Monters) error
	Insert(m *Monters) error
	Update(id int, body *map[string]interface{}) error
	Delete(id int) error
}

type MontersService interface {
	FindAll() ([]Monters, error)
	GetById(id int) (Monters, error)
	Insert(monters Monters) (int, error)
	Update(id int, body dto.MonterValue) error
	Delete(id int) error
}
