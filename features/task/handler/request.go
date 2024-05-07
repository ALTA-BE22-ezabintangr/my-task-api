package handler

type TaskRequest struct {
	UserID          uint   `json:"user_id" form:"user_id"`
	ProjectID       uint   `json:"project_id" form:"project_id"`
	TaskName        string `json:"task_name" form:"task_name"`
	DescriptionTask string `json:"description_task" form:"description_task"`
	StatusTask      string `json:"status_task" form:"status_task"`
}
