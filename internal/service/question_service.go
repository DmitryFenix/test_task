package service

import (
	"errors"
	"qa-api/internal/models"
	"qa-api/internal/repository"
	"strings"
)

// QuestionService handles business logic for questions
type QuestionService struct {
	questionRepo *repository.QuestionRepository
}

// NewQuestionService creates a new QuestionService
func NewQuestionService(questionRepo *repository.QuestionRepository) *QuestionService {
	return &QuestionService{
		questionRepo: questionRepo,
	}
}

// CreateQuestion creates a new question
func (s *QuestionService) CreateQuestion(text string) (*models.Question, error) {
	text = strings.TrimSpace(text)
	if text == "" {
		return nil, errors.New("question text cannot be empty")
	}

	question := &models.Question{
		Text: text,
	}

	if err := s.questionRepo.Create(question); err != nil {
		return nil, err
	}

	return question, nil
}

// GetAllQuestions retrieves all questions
func (s *QuestionService) GetAllQuestions() ([]models.Question, error) {
	return s.questionRepo.GetAll()
}

// GetQuestionByID retrieves a question by ID with its answers
func (s *QuestionService) GetQuestionByID(id int) (*models.Question, error) {
	question, err := s.questionRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return question, nil
}

// DeleteQuestion deletes a question by ID
func (s *QuestionService) DeleteQuestion(id int) error {
	exists, err := s.questionRepo.Exists(id)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("question not found")
	}

	return s.questionRepo.Delete(id)
}





