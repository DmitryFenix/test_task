package service

import (
	"errors"
	"qa-api/internal/models"
	"qa-api/internal/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockQuestionRepository is a mock implementation of QuestionRepository
type MockQuestionRepository struct {
	mock.Mock
}

func (m *MockQuestionRepository) Create(question *models.Question) error {
	args := m.Called(question)
	return args.Error(0)
}

func (m *MockQuestionRepository) GetAll() ([]models.Question, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Question), args.Error(1)
}

func (m *MockQuestionRepository) GetByID(id int) (*models.Question, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Question), args.Error(1)
}

func (m *MockQuestionRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockQuestionRepository) Exists(id int) (bool, error) {
	args := m.Called(id)
	return args.Bool(0), args.Error(1)
}

func TestQuestionService_CreateQuestion(t *testing.T) {
	mockRepo := new(MockQuestionRepository)
	service := NewQuestionService(mockRepo)

	t.Run("successful creation", func(t *testing.T) {
		question := &models.Question{
			ID:        1,
			Text:      "Test question",
			CreatedAt: time.Now(),
		}
		mockRepo.On("Create", mock.AnythingOfType("*models.Question")).Return(nil).Run(func(args mock.Arguments) {
			q := args.Get(0).(*models.Question)
			q.ID = 1
			q.CreatedAt = time.Now()
		})

		result, err := service.CreateQuestion("Test question")

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "Test question", result.Text)
		mockRepo.AssertExpectations(t)
	})

	t.Run("empty text", func(t *testing.T) {
		result, err := service.CreateQuestion("")

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "question text cannot be empty", err.Error())
	})

	t.Run("whitespace only", func(t *testing.T) {
		result, err := service.CreateQuestion("   ")

		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestQuestionService_DeleteQuestion(t *testing.T) {
	mockRepo := new(MockQuestionRepository)
	service := NewQuestionService(mockRepo)

	t.Run("successful deletion", func(t *testing.T) {
		mockRepo.On("Exists", 1).Return(true, nil)
		mockRepo.On("Delete", 1).Return(nil)

		err := service.DeleteQuestion(1)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("question not found", func(t *testing.T) {
		mockRepo.On("Exists", 999).Return(false, nil)

		err := service.DeleteQuestion(999)

		assert.Error(t, err)
		assert.Equal(t, "question not found", err.Error())
	})
}





