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
	res := m.DB.Where("enabled = ?", true).Find(&monters)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (m *montersRepo) GetById(id int, monters *domain.Monters) error {
	// find by id
	res := m.DB.First(&monters, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (m *montersRepo) Insert(monters *domain.Monters) error {
	// insert
	res := m.DB.Create(&monters)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (m *montersRepo) Update(id int, body *map[string]interface{}) error {
	// update
	res := m.DB.Model(domain.Monters{}).Where("id = ?", id).Updates(&body)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (m *montersRepo) Delete(id int) error {
	// delete
	res := m.DB.Delete(domain.Monters{}, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
