package repository

import (
	"github.com/Friske2/pokemon-api/domain"
	"gorm.io/gorm"
)

type montersRepo struct {
	DB *gorm.DB
}

func NewMontersRepo(db *gorm.DB) domain.MontersRepository {
	return &montersRepo{
		DB: db,
	}
}

func (m *montersRepo) FindAll(monters *[]domain.Monters) error {
	// find all
	return nil
}

func (m *montersRepo) GetById(id int, monters *domain.Monters) error {
	// find by id
	return nil
}

func (m *montersRepo) Insert(monters *domain.Monters) error {
	// insert
	return nil
}

func (m *montersRepo) Update(monters *domain.Monters) error {
	// update
	return nil
}

func (m *montersRepo) Delete(monters *domain.Monters) error {
	// delete
	return nil
}
