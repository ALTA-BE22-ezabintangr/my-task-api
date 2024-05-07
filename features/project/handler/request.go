package handler

type ProjectRequest struct {
	UserID      uint   `gorm:"unique" json:"user_id" form:"user_id"`
	ProjectName string `json:"project_name" form:"project_name"`
	Description string `json:"description" form:"description"`
}
