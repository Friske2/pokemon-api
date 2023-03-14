package domain

import (
	"github.com/Friske2/pokemon-api/dto"
	"github.com/Friske2/pokemon-api/model"
)

type MontersRepository interface {
	FindAll(m *[]model.Monter) error
	GetById(id int32, m *model.Monter) error
	Insert(m *model.Monter) error
	Update(id int32, body *map[string]interface{}) error
	Delete(id int32) error
}

type MontersService interface {
	FindAll() ([]model.Monter, error)
	GetById(id int32) (model.Monter, error)
	Insert(monters model.Monter) (int32, error)
	Update(id int32, body dto.MonterValue) error
	Delete(id int32) error
}
