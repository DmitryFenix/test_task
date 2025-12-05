package service

import "qa-api/internal/models"

// QuestionServiceInterface defines the interface for question service
type QuestionServiceInterface interface {
	CreateQuestion(text string) (*models.Question, error)
	GetAllQuestions() ([]models.Question, error)
	GetQuestionByID(id int) (*models.Question, error)
	DeleteQuestion(id int) error
}

// AnswerServiceInterface defines the interface for answer service
type AnswerServiceInterface interface {
	CreateAnswer(questionID int, userID, text string) (*models.Answer, error)
	GetAnswerByID(id int) (*models.Answer, error)
	DeleteAnswer(id int) error
}





