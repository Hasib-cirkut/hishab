package dtos

type AddCategoryDto struct {
	Name string `json:"name" binding:"required"`
	Type string `json:"type" binding:"required"`
}
