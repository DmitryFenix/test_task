package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"qa-api/internal/service"
	"strconv"

	"github.com/gorilla/mux"
)

// AnswerHandler handles HTTP requests for answers
type AnswerHandler struct {
	answerService service.AnswerServiceInterface
}

// NewAnswerHandler creates a new AnswerHandler
func NewAnswerHandler(answerService service.AnswerServiceInterface) *AnswerHandler {
	return &AnswerHandler{
		answerService: answerService,
	}
}

// CreateAnswerRequest represents the request body for creating an answer
type CreateAnswerRequest struct {
	UserID string `json:"user_id"`
	Text   string `json:"text"`
}

// CreateAnswer handles POST /questions/{id}/answers/
func (h *AnswerHandler) CreateAnswer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	questionID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid question ID", http.StatusBadRequest)
		return
	}

	var req CreateAnswerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	answer, err := h.answerService.CreateAnswer(questionID, req.UserID, req.Text)
	if err != nil {
		log.Printf("Error creating answer: %v", err)
		if err.Error() == "question not found" {
			http.Error(w, "Question not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}

// GetAnswer handles GET /answers/{id}
func (h *AnswerHandler) GetAnswer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid answer ID", http.StatusBadRequest)
		return
	}

	answer, err := h.answerService.GetAnswerByID(id)
	if err != nil {
		log.Printf("Error getting answer: %v", err)
		http.Error(w, "Answer not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(answer)
}

// DeleteAnswer handles DELETE /answers/{id}
func (h *AnswerHandler) DeleteAnswer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid answer ID", http.StatusBadRequest)
		return
	}

	if err := h.answerService.DeleteAnswer(id); err != nil {
		log.Printf("Error deleting answer: %v", err)
		if err.Error() == "answer not found" {
			http.Error(w, "Answer not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

