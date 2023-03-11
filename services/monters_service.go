package services

import (
	"github.com/Friske2/pokemon-api/domain"
	"github.com/Friske2/pokemon-api/dto"
)

type MonterService struct {
	monterRepo domain.MontersRepository
}

func NewMontersService(m domain.MontersRepository) domain.MontersService {
	return &MonterService{
		monterRepo: m,
	}
}

func (s *MonterService) FindAll() ([]domain.Monters, error) {
	monters := []domain.Monters{}
	err := s.monterRepo.FindAll(&monters)
	if err != nil {
		return nil, err
	}
	return monters, nil
}

func (s *MonterService) GetById(id int) (domain.Monters, error) {
	monters := domain.Monters{}
	err := s.monterRepo.GetById(id, &monters)
	if err != nil {
		return monters, err
	}
	return monters, nil
}

func (s *MonterService) Insert(monters domain.Monters) (int, error) {
	body := domain.Monters{
		Name: monters.Name,
	}
	err := s.monterRepo.Insert(&body)
	if err != nil {
		return 0, err
	}
	return body.ID, nil
}

func (s *MonterService) Update(id int, body dto.MonterValue) error {
	monters := domain.Monters{
		ID: id,
	}
	err := s.monterRepo.GetById(id, &monters)
	if err != nil {
		return err
	}
	var updateBody = map[string]interface{}{
		"Name": body.Name,
	}
	if body.Enabled != nil {
		updateBody["Enabled"] = *body.Enabled
	}
	err = s.monterRepo.Update(id, &updateBody)
	if err != nil {
		return err
	}
	return nil
}

func (s *MonterService) Delete(id int) error {
	err := s.monterRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
