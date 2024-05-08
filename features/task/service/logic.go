package service

import (
	"errors"
	"myTaskApp/features/task"
)

type taskService struct {
	taskData task.DataInterface
}

func New(td task.DataInterface) task.ServiceInterface {
	return &taskService{
		taskData: td,
	}
}

// Create implements task.ServiceInterface.
func (t *taskService) Create(input task.Core) error {
	if input.UserID == 0 || input.ProjectID == 0 || input.TaskName == "" {
		return errors.New("id user/project dan nama task tidak boleh kosong")
	}
	err := t.taskData.Insert(input)
	if err != nil {
		return err
	}
	return nil
}

// GetTaskbyId implements task.ServiceInterface.
func (t *taskService) GetTaskbyId(id uint) ([]task.Core, error) {
	return t.taskData.GetTaskbyUserId(id)
}

// Delete implements task.ServiceInterface.
func (t *taskService) Delete(id uint) error {
	panic("unimplemented")
}

// Update implements task.ServiceInterface.
func (t *taskService) Update(id uint, input task.Core) error {
	panic("unimplemented")
}