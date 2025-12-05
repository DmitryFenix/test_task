package service

import (
	"errors"
	"qa-api/internal/models"
	"qa-api/internal/repository"
	"strings"
)

// AnswerService handles business logic for answers
type AnswerService struct {
	answerRepo   *repository.AnswerRepository
	questionRepo *repository.QuestionRepository
}

// NewAnswerService creates a new AnswerService
func NewAnswerService(answerRepo *repository.AnswerRepository, questionRepo *repository.QuestionRepository) *AnswerService {
	return &AnswerService{
		answerRepo:   answerRepo,
		questionRepo: questionRepo,
	}
}

// CreateAnswer creates a new answer for a question
func (s *AnswerService) CreateAnswer(questionID int, userID, text string) (*models.Answer, error) {
	text = strings.TrimSpace(text)
	if text == "" {
		return nil, errors.New("answer text cannot be empty")
	}

	userID = strings.TrimSpace(userID)
	if userID == "" {
		return nil, errors.New("user_id cannot be empty")
	}

	// Check if question exists
	exists, err := s.questionRepo.Exists(questionID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("question not found")
	}

	answer := &models.Answer{
		QuestionID: questionID,
		UserID:     userID,
		Text:       text,
	}

	if err := s.answerRepo.Create(answer); err != nil {
		return nil, err
	}

	return answer, nil
}

// GetAnswerByID retrieves an answer by ID
func (s *AnswerService) GetAnswerByID(id int) (*models.Answer, error) {
	answer, err := s.answerRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return answer, nil
}

// DeleteAnswer deletes an answer by ID
func (s *AnswerService) DeleteAnswer(id int) error {
	answer, err := s.answerRepo.GetByID(id)
	if err != nil {
		return errors.New("answer not found")
	}
	if answer == nil {
		return errors.New("answer not found")
	}

	return s.answerRepo.Delete(id)
}





