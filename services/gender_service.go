package services

import (
	"github.com/Friske2/pokemon-api/domain"
	"github.com/Friske2/pokemon-api/model"
)

type genderService struct {
	genderRepo domain.GenderRepository
}

func NewGenderService(genderRepo domain.GenderRepository) domain.GenderService {
	return &genderService{
		genderRepo: genderRepo,
	}
}

func (s *genderService) FindAll() ([]model.Gender, error) {
	result := []model.Gender{}
	res := s.genderRepo.FindAll(&result)
	if res != nil {
		return nil, res
	}
	return result, nil
}
func (s *genderService) GetById(id int) (model.Gender, error) {
	var result model.Gender
	err := s.genderRepo.GetById(id, &result)
	if err != nil {
		return result, err
	}
	if result.ID == 0 {
		return result, domain.ErrNotFound
	}
	return result, nil
}
func (s *genderService) Insert(name string) error {
	err := s.genderRepo.Insert(name)
	if err != nil {
		return err
	}
	return err
}
func (s *genderService) Update(id int, name string) error {
	_, err := s.GetById(id)
	if err != nil {
		return err
	}
	err = s.genderRepo.Update(id, name)
	if err != nil {
		return err
	}
	return nil
}
func (s *genderService) Delete(id int) error {
	_, err := s.GetById(id)
	if err != nil {
		return err
	}
	err = s.genderRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
