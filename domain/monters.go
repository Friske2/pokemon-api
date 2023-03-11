package domain

import (
	"github.com/Friske2/pokemon-api/dto"
)

type Monters struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	TypeId   int    `json:"typeId"`
	GenderId int    `json:"genderId"`
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
