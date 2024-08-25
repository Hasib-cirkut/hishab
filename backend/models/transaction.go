package models

import "time"

type Transaction struct {
	ID              uint      `gorm:"primaryKey;autoIncrement"`
	UserID          uint      `gorm:"not null"`
	CategoryID      uint      `gorm:"not null"`
	Amount          float64   `gorm:"type:decimal(10,2);not null"`
	TransactionDate time.Time `gorm:"type:date;not null"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Description     string    `gorm:"type:text"`

	User     User     `gorm:"foreignKey:UserID"`
	Category Category `gorm:"foreignKey:CategoryID"`
}
