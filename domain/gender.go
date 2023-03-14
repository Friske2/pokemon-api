package domain

import "github.com/Friske2/pokemon-api/model"

type GenderRepository interface {
	FindAll(m *[]model.Gender) error
	GetById(id int, m *model.Gender) error
	Insert(name string) error
	Update(id int, name string) error
	Delete(id int) error
}

type GenderService interface {
	FindAll() ([]model.Gender, error)
	GetById(id int) (model.Gender, error)
	Insert(name string) error
	Update(id int, name string) error
	Delete(id int) error
}
