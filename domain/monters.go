package domain

import "time"

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
	Update(m *Monters) error
	Delete(m *Monters) error
}

type MontersService interface {
	FindAll() ([]Monters, error)
	GetById(id int) (Monters, error)
	Insert(monters Monters) (int, error)
	Update(monters Monters) error
	Delete(id int) error
}
