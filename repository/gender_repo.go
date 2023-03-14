package repository

import (
	"github.com/Friske2/pokemon-api/domain"
	"gorm.io/gorm"
)

type genderRepo struct {
	DB *gorm.DB
}

func NewGenderRepo(db *gorm.DB) domain.GenderRepository {
	return &genderRepo{
		DB: db,
	}
}

// Find All
func (r *genderRepo) FindAll(genders *[]domain.Gender) error {
	res := r.DB.Find(&genders)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// Find ById
func (r *genderRepo) GetById(id int, gender *domain.Gender) error {
	res := r.DB.First(&gender, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// Insert Gender
func (r *genderRepo) Insert(name string) error {
	res := r.DB.Create(&domain.Gender{Name: name})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// Update Gender
func (r *genderRepo) Update(id int, name string) error {
	body := map[string]interface{}{
		"name": name,
	}
	res := r.DB.Model(&domain.Gender{}).Where("id = ?", id).Updates(body)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// Delete Gender
func (r *genderRepo) Delete(id int) error {
	res := r.DB.Where("id = ?", id).Delete(&domain.Gender{})
	if res.Error != nil {
		return res.Error
	}
	return nil
}
