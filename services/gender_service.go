package services

import "github.com/Friske2/pokemon-api/domain"

type genderService struct {
	genderRepo domain.GenderRepository
}

func NewGenderService(genderRepo domain.GenderRepository) domain.GenderService {
	return &genderService{
		genderRepo: genderRepo,
	}
}

func (s *genderService) FindAll() ([]domain.Gender, error) {
	result := []domain.Gender{}
	res := s.genderRepo.FindAll(&result)
	if res != nil {
		return nil, res
	}
	return result, nil
}
func (s *genderService) GetById(id int) (domain.Gender, error) {
	var result domain.Gender
	err := s.genderRepo.GetById(id, &result)
	if err != nil {
		return result, err
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
	res, err := s.GetById(id)
	if err != nil {
		return err
	}
	if res.ID == 0 {
		return domain.ErrNotFound
	}
	err = s.genderRepo.Update(id, name)
	if err != nil {
		return err
	}
	return nil
}
func (s *genderService) Delete(id int) error {
	res, err := s.GetById(id)
	if err != nil {
		return err
	}
	if res.ID == 0 {
		return domain.ErrNotFound
	}
	err = s.genderRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
