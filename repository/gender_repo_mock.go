package repository

import (
	"github.com/Friske2/pokemon-api/domain"
	"github.com/stretchr/testify/mock"
)

type genderRepoMock struct {
	mock.Mock
}

func NewGenderRepoMock() *genderRepoMock {
	return &genderRepoMock{}
}

func (m *genderRepoMock) FindAll(gender *[]domain.Gender) error {
	args := m.Called()
	*gender = args.Get(0).([]domain.Gender)
	return args.Error(1)
}

func (m *genderRepoMock) GetById(id int, gender *domain.Gender) error {
	args := m.Called()
	*gender = args.Get(0).(domain.Gender)
	return args.Error(1)
}
func (m *genderRepoMock) Insert(name string) error {
	args := m.Called()
	return args.Error(0)
}
func (m *genderRepoMock) Update(id int, name string) error {
	args := m.Called()
	return args.Error(0)
}
func (m *genderRepoMock) Delete(id int) error {
	args := m.Called()
	return args.Error(0)
}
