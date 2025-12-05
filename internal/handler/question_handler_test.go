package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"qa-api/internal/models"
	"qa-api/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockQuestionService is a mock implementation of QuestionServiceInterface
type MockQuestionService struct {
	mock.Mock
}

// Ensure MockQuestionService implements QuestionServiceInterface
var _ service.QuestionServiceInterface = (*MockQuestionService)(nil)

func (m *MockQuestionService) CreateQuestion(text string) (*models.Question, error) {
	args := m.Called(text)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Question), args.Error(1)
}

func (m *MockQuestionService) GetAllQuestions() ([]models.Question, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Question), args.Error(1)
}

func (m *MockQuestionService) GetQuestionByID(id int) (*models.Question, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Question), args.Error(1)
}

func (m *MockQuestionService) DeleteQuestion(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestQuestionHandler_CreateQuestion(t *testing.T) {
	mockService := new(MockQuestionService)
	handler := NewQuestionHandler(mockService)

	t.Run("successful creation", func(t *testing.T) {
		expectedQuestion := &models.Question{
			ID:   1,
			Text: "Test question",
		}
		mockService.On("CreateQuestion", "Test question").Return(expectedQuestion, nil)

		reqBody := CreateQuestionRequest{Text: "Test question"}
		jsonBody, _ := json.Marshal(reqBody)
		req := httptest.NewRequest("POST", "/questions/", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		handler.CreateQuestion(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("empty text", func(t *testing.T) {
		mockService.On("CreateQuestion", "").Return(nil, assert.AnError)

		reqBody := CreateQuestionRequest{Text: ""}
		jsonBody, _ := json.Marshal(reqBody)
		req := httptest.NewRequest("POST", "/questions/", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		handler.CreateQuestion(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestQuestionHandler_GetQuestions(t *testing.T) {
	mockService := new(MockQuestionService)
	handler := NewQuestionHandler(mockService)

	t.Run("successful get all", func(t *testing.T) {
		expectedQuestions := []models.Question{
			{ID: 1, Text: "Question 1"},
			{ID: 2, Text: "Question 2"},
		}
		mockService.On("GetAllQuestions").Return(expectedQuestions, nil)

		req := httptest.NewRequest("GET", "/questions/", nil)
		w := httptest.NewRecorder()

		handler.GetQuestions(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockService.AssertExpectations(t)
	})
}

