package services_test

import (
	"testing"

	"github.com/Friske2/pokemon-api/domain"
	"github.com/Friske2/pokemon-api/repository"
	"github.com/Friske2/pokemon-api/services"
	"github.com/stretchr/testify/assert"
)

func TestGenderRepo(t *testing.T) {
	t.Run("FindAll", func(t *testing.T) {
		// Arrange
		genderRepo := repository.NewGenderRepoMock()
		var payload []domain.Gender
		mock := []domain.Gender{
			{
				ID:   1,
				Name: "Male",
			},
			{
				ID:   2,
				Name: "Female",
			},
		}
		genderRepo.On("FindAll").Return(mock, nil)
		genderService := services.NewGenderService(genderRepo)
		// Act
		payload, _ = genderService.FindAll()
		expect := mock
		// Assert
		assert.Equal(t, expect, payload)
	})

	t.Run("GetById", func(t *testing.T) {
		// Arrange
		genderRepo := repository.NewGenderRepoMock()
		var payload domain.Gender
		mock := domain.Gender{
			ID:   1,
			Name: "Male",
		}
		genderRepo.On("GetById").Return(mock, nil)
		genderService := services.NewGenderService(genderRepo)
		// Act
		payload, _ = genderService.GetById(1)
		expect := mock
		// Assert
		assert.Equal(t, expect, payload)
	})

	t.Run("GetByIdNotFound", func(t *testing.T) {
		// Arrange
		genderRepo := repository.NewGenderRepoMock()
		genderRepo.On("GetById").Return(domain.Gender{}, domain.ErrNotFound)
		// Act
		genderService := services.NewGenderService(genderRepo)
		// Assert
		_, err := genderService.GetById(1)
		assert.Equal(t, domain.ErrNotFound, err)
	})

	t.Run("Insert", func(t *testing.T) {
		// Arrange
		genderRepo := repository.NewGenderRepoMock()
		genderRepo.On("Insert").Return(nil)
		genderService := services.NewGenderService(genderRepo)
		// Act
		err := genderService.Insert("Female")
		// Assert
		assert.Equal(t, nil, err)
	})

	t.Run("Update", func(t *testing.T) {
		// Arrange
		genderRepo := repository.NewGenderRepoMock()
		mock := domain.Gender{
			ID:   1,
			Name: "Male",
		}
		genderRepo.On("GetById").Return(mock, nil)
		genderRepo.On("Update").Return(nil)
		genderService := services.NewGenderService(genderRepo)
		// Act
		err := genderService.Update(1, "Female")
		// Assert
		assert.Equal(t, nil, err)
	})

	t.Run("Delete", func(t *testing.T) {
		// Arrange
		genderRepo := repository.NewGenderRepoMock()
		mock := domain.Gender{
			ID:   1,
			Name: "Male",
		}
		genderRepo.On("GetById").Return(mock, nil)
		genderRepo.On("Delete").Return(nil)
		genderService := services.NewGenderService(genderRepo)
		// Act
		err := genderService.Delete(1)
		// Assert
		assert.Equal(t, nil, err)
	})
}
