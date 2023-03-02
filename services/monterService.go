package services

import "github.com/Friske2/pokemon-api/domain"

type MonterService struct {
	monterRepo domain.MontersRepository
}

// func NewMonterService(m domain.MontersRepository) domain.MontersService {
// 	return &MonterService{
// 		monterRepo: m,
// 	}
// }
