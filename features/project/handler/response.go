package handler

type ProjectResponse struct {
	ID          uint
	ProjectName string
}

type ProjectResponseById struct {
	ID          uint
	ProjectName string
	Description string
}
