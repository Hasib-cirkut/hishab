package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"hishab.com/main/config"
	"hishab.com/main/dtos"
	"hishab.com/main/models"
)

func AddCategory(c *gin.Context) {
	var categoryJson dtos.AddCategoryDto

	if err := c.ShouldBindJSON(&categoryJson); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if categoryJson.Type != models.EARNING && categoryJson.Type != models.EXPENDITURE {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Category type must of earning or expenditure"})

		return
	}

	var existingCategory models.Category

	if err := config.DB.Where("name = ? AND type = ?", categoryJson.Name, categoryJson.Type).First(&existingCategory).Error; err == nil {

		c.JSON(http.StatusConflict, gin.H{
			"error": fmt.Sprintf("Category with name '%s' already exists for type '%s'", categoryJson.Name, categoryJson.Type),
		})

		return
	}

	var category models.Category

	category.Name = categoryJson.Name
	category.Type = categoryJson.Type

	if err := config.DB.Create(&category).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusCreated, category)
}

func FetchAllCategory(c *gin.Context) {
	var categories []models.Category

	err := config.DB.Find(&categories).Error

	if err != nil {
		fmt.Print("Something went wrong while fetching categories")

		return
	}

	c.JSON(http.StatusOK, categories)
}
