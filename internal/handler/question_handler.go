package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"qa-api/internal/service"
	"strconv"

	"github.com/gorilla/mux"
)

// QuestionHandler handles HTTP requests for questions
type QuestionHandler struct {
	questionService service.QuestionServiceInterface
}

// NewQuestionHandler creates a new QuestionHandler
func NewQuestionHandler(questionService service.QuestionServiceInterface) *QuestionHandler {
	return &QuestionHandler{
		questionService: questionService,
	}
}

// CreateQuestionRequest represents the request body for creating a question
type CreateQuestionRequest struct {
	Text string `json:"text"`
}

// GetQuestions handles GET /questions/
func (h *QuestionHandler) GetQuestions(w http.ResponseWriter, r *http.Request) {
	questions, err := h.questionService.GetAllQuestions()
	if err != nil {
		log.Printf("Error getting questions: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questions)
}

// CreateQuestion handles POST /questions/
func (h *QuestionHandler) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	var req CreateQuestionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	question, err := h.questionService.CreateQuestion(req.Text)
	if err != nil {
		log.Printf("Error creating question: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(question)
}

// GetQuestion handles GET /questions/{id}
func (h *QuestionHandler) GetQuestion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid question ID", http.StatusBadRequest)
		return
	}

	question, err := h.questionService.GetQuestionByID(id)
	if err != nil {
		log.Printf("Error getting question: %v", err)
		http.Error(w, "Question not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(question)
}

// DeleteQuestion handles DELETE /questions/{id}
func (h *QuestionHandler) DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid question ID", http.StatusBadRequest)
		return
	}

	if err := h.questionService.DeleteQuestion(id); err != nil {
		log.Printf("Error deleting question: %v", err)
		if err.Error() == "question not found" {
			http.Error(w, "Question not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

