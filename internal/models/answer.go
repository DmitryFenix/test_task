package models

import "time"

// Answer represents an answer to a question
type Answer struct {
	ID         int       `gorm:"primaryKey;autoIncrement" json:"id"`
	QuestionID int       `gorm:"not null;index" json:"question_id"`
	UserID     string    `gorm:"type:varchar(255);not null;index" json:"user_id"`
	Text       string    `gorm:"type:text;not null" json:"text"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	Question   Question  `gorm:"foreignKey:QuestionID" json:"question,omitempty"`
}

// TableName specifies the table name for Answer
func (Answer) TableName() string {
	return "answers"
}





