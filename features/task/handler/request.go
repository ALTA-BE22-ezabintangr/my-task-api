package handler

type TaskRequest struct {
	ProjectID       uint   `json:"project_id" form:"project_id"`
	TaskName        string `json:"task_name" form:"task_name"`
	DescriptionTask string `json:"description_task" form:"description_task"`
	StatusTask      string `json:"status_task" form:"status_task"`
}

type TaskUpdateRequest struct {
	TaskName        string `json:"task_name" form:"task_name"`
	DescriptionTask string `json:"description_task" form:"description_task"`
	StatusTask      string `json:"status_task" form:"status_task"`
}
