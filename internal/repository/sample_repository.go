package repository

import (
	"go-api-project-seed/internal/model"

	"gorm.io/gorm"
)

// SampleRepository handles database operations for Sample.
type SampleRepository struct {
	db *gorm.DB
}

// NewSampleRepository initializes a new instance of SampleRepository.
func NewSampleRepository(db *gorm.DB) *SampleRepository {
	return &SampleRepository{db: db}
}

// FindAll retrieves all records.
func (r *SampleRepository) FindAll() ([]model.Sample, error) {
	var entities []model.Sample
	err := r.db.Find(&entities).Error
	return entities, err
}

// FindByID retrieves a record by ID.
func (r *SampleRepository) FindByID(id uint) (*model.Sample, error) {
	var entity model.Sample
	err := r.db.First(&entity, id).Error
	return &entity, err
}

// Create adds a new record.
func (r *SampleRepository) Create(entity *model.Sample) error {
	return r.db.Create(entity).Error
}

// Update modifies an existing record.
func (r *SampleRepository) Update(entity *model.Sample) error {
	return r.db.Save(entity).Error
}

// Delete removes a record by ID.
func (r *SampleRepository) Delete(id uint) error {
	return r.db.Delete(&model.Sample{}, id).Error
}
