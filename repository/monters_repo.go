package repository

import (
	"github.com/Friske2/pokemon-api/domain"
	"github.com/Friske2/pokemon-api/model"
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

func (m *montersRepo) FindAll(monters *[]model.Monter) error {
	// find all
	res := m.DB.Find(&monters)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (m *montersRepo) GetById(id int32, monters *model.Monter) error {
	// find by id
	res := m.DB.First(&monters, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (m *montersRepo) Insert(monters *model.Monter) error {
	// insert
	res := m.DB.Create(&monters)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (m *montersRepo) Update(id int32, body *map[string]interface{}) error {
	// update
	res := m.DB.Model(model.Monter{}).Where("id = ?", id).Updates(&body)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (m *montersRepo) Delete(id int32) error {
	// delete
	res := m.DB.Delete(model.Monter{}, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
