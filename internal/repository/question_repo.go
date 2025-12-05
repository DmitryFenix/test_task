package repository

import (
	"qa-api/internal/database"
	"qa-api/internal/models"
)

// QuestionRepository handles database operations for questions
type QuestionRepository struct{}

// NewQuestionRepository creates a new QuestionRepository
func NewQuestionRepository() *QuestionRepository {
	return &QuestionRepository{}
}

// Create creates a new question
func (r *QuestionRepository) Create(question *models.Question) error {
	return database.GetDB().Create(question).Error
}

// GetAll retrieves all questions
func (r *QuestionRepository) GetAll() ([]models.Question, error) {
	var questions []models.Question
	err := database.GetDB().Find(&questions).Error
	return questions, err
}

// GetByID retrieves a question by ID with its answers
func (r *QuestionRepository) GetByID(id int) (*models.Question, error) {
	var question models.Question
	err := database.GetDB().Preload("Answers").First(&question, id).Error
	return &question, err
}

// Delete deletes a question by ID (cascade delete will handle answers)
func (r *QuestionRepository) Delete(id int) error {
	return database.GetDB().Delete(&models.Question{}, id).Error
}

// Exists checks if a question exists
func (r *QuestionRepository) Exists(id int) (bool, error) {
	var count int64
	err := database.GetDB().Model(&models.Question{}).Where("id = ?", id).Count(&count).Error
	return count > 0, err
}





