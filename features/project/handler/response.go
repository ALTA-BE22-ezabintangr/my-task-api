package handler

type ProjectResponse struct {
	ID          uint
	UserID      uint
	ProjectName string
}

type ProjectResponseById struct {
	ID          uint
	UserID      uint
	ProjectName string
	Description string
}
