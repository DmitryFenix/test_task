package models

import "time"

// Question represents a question in the system
type Question struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Text      string    `gorm:"type:text;not null" json:"text"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	Answers   []Answer  `gorm:"foreignKey:QuestionID;constraint:OnDelete:CASCADE" json:"answers,omitempty"`
}

// TableName specifies the table name for Question
func (Question) TableName() string {
	return "questions"
}





