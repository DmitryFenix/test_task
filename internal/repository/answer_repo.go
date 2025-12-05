package repository

import (
	"qa-api/internal/database"
	"qa-api/internal/models"
)

// AnswerRepository handles database operations for answers
type AnswerRepository struct{}

// NewAnswerRepository creates a new AnswerRepository
func NewAnswerRepository() *AnswerRepository {
	return &AnswerRepository{}
}

// Create creates a new answer
func (r *AnswerRepository) Create(answer *models.Answer) error {
	return database.GetDB().Create(answer).Error
}

// GetByID retrieves an answer by ID
func (r *AnswerRepository) GetByID(id int) (*models.Answer, error) {
	var answer models.Answer
	err := database.GetDB().First(&answer, id).Error
	return &answer, err
}

// Delete deletes an answer by ID
func (r *AnswerRepository) Delete(id int) error {
	return database.GetDB().Delete(&models.Answer{}, id).Error
}





