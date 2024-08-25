package models

import "gorm.io/gorm"

type Category struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"unique;not null"`
	Type string `gorm:"not null"`
}

func InsertDefaultCategories(db *gorm.DB) {
	categories := []Category{
		{Name: "Salary", Type: "earning"},
		{Name: "Food", Type: "expenditure"},
		{Name: "Transportation", Type: "expenditure"},
		{Name: "Payment", Type: "expenditure"},
		{Name: "ATM", Type: "expenditure"},
		{Name: "Grocery", Type: "expenditure"},
	}

	for _, category := range categories {
		db.FirstOrCreate(&category, Category{Name: category.Name})
	}
}
