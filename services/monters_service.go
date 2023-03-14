package services

import (
	"github.com/Friske2/pokemon-api/domain"
	"github.com/Friske2/pokemon-api/dto"
	"github.com/Friske2/pokemon-api/model"
)

type monterService struct {
	monterRepo domain.MontersRepository
}

func NewMontersService(m domain.MontersRepository) domain.MontersService {
	return &monterService{
		monterRepo: m,
	}
}

func (s *monterService) FindAll() ([]model.Monter, error) {
	monters := []model.Monter{}
	err := s.monterRepo.FindAll(&monters)
	if err != nil {
		return nil, err
	}
	return monters, nil
}

func (s *monterService) GetById(id int32) (model.Monter, error) {
	monters := model.Monter{}
	err := s.monterRepo.GetById(id, &monters)
	if err != nil {
		return monters, err
	}
	return monters, nil
}

func (s *monterService) Insert(monters model.Monter) (int32, error) {
	body := model.Monter{
		Name: monters.Name,
	}
	err := s.monterRepo.Insert(&body)
	if err != nil {
		return 0, err
	}
	return body.ID, nil
}

func (s *monterService) Update(id int32, body dto.MonterValue) error {
	monters := model.Monter{
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

func (s *monterService) Delete(id int32) error {
	err := s.monterRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
