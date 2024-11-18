package service

import (
	"go-api-project-seed/internal/model"
	"go-api-project-seed/internal/repository"
)

// SampleService handles business logic for Sample.
type SampleService struct {
	repo *repository.SampleRepository
}

// NewSampleService initializes a new instance of SampleService.
func NewSampleService(repo *repository.SampleRepository) *SampleService {
	return &SampleService{repo: repo}
}

// GetAll retrieves all entries.
func (s *SampleService) GetAll() ([]model.Sample, error) {
	return s.repo.FindAll()
}

// GetByID retrieves an entry by ID.
func (s *SampleService) GetByID(id uint) (*model.Sample, error) {
	return s.repo.FindByID(id)
}

// Create adds a new entry.
func (s *SampleService) Create(entity *model.Sample) error {
	return s.repo.Create(entity)
}

// Update modifies an existing entry.
func (s *SampleService) Update(entity *model.Sample) error {
	return s.repo.Update(entity)
}

// Delete removes an existing entry.
func (s *SampleService) Delete(id uint) error {
	return s.repo.Delete(id)
}
