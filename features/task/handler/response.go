package handler

type ResponseById struct {
	ProjectID  uint   `json:"project_id"`
	TaskName   string `json:"task_name"`
	StatusTask string `json:"status_task"`
}
