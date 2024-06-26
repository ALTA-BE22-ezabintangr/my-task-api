package service

import (
	"errors"
	"myTaskApp/features/project"
	"myTaskApp/features/task"
)

type taskService struct {
	taskData    task.DataInterface
	projectData project.DataInterface
}

func New(td task.DataInterface, pd project.DataInterface) task.ServiceInterface {
	return &taskService{
		taskData:    td,
		projectData: pd,
	}
}

// Create implements task.ServiceInterface.
func (t *taskService) Create(input task.Core) error {
	result, err := t.projectData.GetUserByProjectId(input.ProjectID)
	if err != nil {
		return err
	}
	if result.UserID != input.UserID {
		return errors.New("id project bukan milik anda")
	}

	if input.ProjectID == 0 || input.TaskName == "" {
		return errors.New("id project/nama task tidak boleh kosong")
	}

	err2 := t.taskData.Insert(input)
	if err2 != nil {
		return err2
	}
	return nil
}

// Delete implements task.ServiceInterface.
func (t *taskService) Delete(id uint, idUser uint) error {
	result, err := t.taskData.GetTaskById(id)
	if err != nil {
		return err
	}
	result2, err2 := t.projectData.GetUserByProjectId(result.ProjectID)
	if err2 != nil {
		return err2
	}
	if result2.UserID != idUser {
		return errors.New("id task bukan milik anda")
	}
	return t.taskData.Delete(id)
}

// Update implements task.ServiceInterface.
func (t *taskService) Update(id uint, idUser uint, input task.Core) error {
	result, err := t.taskData.GetTaskById(id)
	if err != nil {
		return err
	}
	result2, err2 := t.projectData.GetUserByProjectId(result.ProjectID)
	if err2 != nil {
		return err2
	}
	if result2.UserID != idUser {
		return errors.New("id task bukan milik anda")
	}

	return t.taskData.Update(id, input)
}
