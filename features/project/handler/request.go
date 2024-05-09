package handler

type ProjectRequest struct {
	ProjectName string `json:"project_name" form:"project_name"`
	Description string `json:"description" form:"description"`
}

type UpdateRequest struct {
	ProjectName string `json:"project_name" form:"project_name"`
	Description string `json:"description" form:"description"`
}
