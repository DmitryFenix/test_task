package main

import (
	"fmt"
	"log"
	"net/http"
	"qa-api/internal/config"
	"qa-api/internal/database"
	"qa-api/internal/handler"
	"qa-api/internal/repository"
	"qa-api/internal/service"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/pressly/goose/v3"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize database
	if err := database.Init(cfg); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Run migrations
	if err := runMigrations(cfg.DatabaseURL); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Initialize repositories
	questionRepo := repository.NewQuestionRepository()
	answerRepo := repository.NewAnswerRepository()

	// Initialize services
	questionService := service.NewQuestionService(questionRepo)
	answerService := service.NewAnswerService(answerRepo, questionRepo)

	// Initialize handlers
	questionHandler := handler.NewQuestionHandler(questionService)
	answerHandler := handler.NewAnswerHandler(answerService)

	// Setup routes
	router := mux.NewRouter()

	// Question routes
	router.HandleFunc("/questions/", questionHandler.GetQuestions).Methods("GET")
	router.HandleFunc("/questions/", questionHandler.CreateQuestion).Methods("POST")
	router.HandleFunc("/questions/{id}", questionHandler.GetQuestion).Methods("GET")
	router.HandleFunc("/questions/{id}", questionHandler.DeleteQuestion).Methods("DELETE")

	// Answer routes
	router.HandleFunc("/questions/{id}/answers/", answerHandler.CreateAnswer).Methods("POST")
	router.HandleFunc("/answers/{id}", answerHandler.GetAnswer).Methods("GET")
	router.HandleFunc("/answers/{id}", answerHandler.DeleteAnswer).Methods("DELETE")

	// Health check
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	// Start server
	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Server starting on %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func runMigrations(databaseURL string) error {
	db, err := goose.OpenDBWithDriver("postgres", databaseURL)
	if err != nil {
		return fmt.Errorf("failed to open database for migrations: %w", err)
	}
	defer db.Close()

	if err := goose.Up(db, "migrations"); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	log.Println("Migrations completed successfully")
	return nil
}

