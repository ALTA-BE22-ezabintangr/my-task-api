package data

import (
	"errors"
	"myTaskApp/features/project/data"
	"myTaskApp/features/task"

	"gorm.io/gorm"
)

type taskQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) task.DataInterface {
	return &taskQuery{
		db: db,
	}
}

// Insert implements task.DataInterface.
func (t *taskQuery) Insert(input task.Core) error {
	var projectForTask ProjectBowl
	tx := t.db.Model(&data.Project{}).Where("id = ?", input.ProjectID).First(&projectForTask)
	if tx.Error != nil {
		return tx.Error
	}
	if projectForTask.UserID != input.UserID {
		return errors.New("id project yang diinput bukan milik anda")
	}

	var taskGorm Task
	taskGorm = Task{
		Model:           gorm.Model{},
		ProjectID:       input.ProjectID,
		TaskName:        input.TaskName,
		DescriptionTask: input.DescriptionTask,
		StatusTask:      input.StatusTask,
	}

	tx2 := t.db.Create(&taskGorm)
	if tx2.Error != nil {
		return tx2.Error
	}

	return nil
}

// GetTaskbyUserId implements task.DataInterface.
func (t *taskQuery) GetTaskById(id uint, idUser uint) (task.Core, error) {
	var allTaskCurrent Task
	tx2 := t.db.First(&allTaskCurrent, id)
	if tx2.Error != nil {
		return task.Core{}, tx2.Error
	}

	var projectForTask ProjectBowl
	tx := t.db.Model(&data.Project{}).Where("id = ?", allTaskCurrent.ProjectID).First(&projectForTask)
	if tx.Error != nil {
		return task.Core{}, tx.Error
	}

	if projectForTask.UserID != idUser {
		return task.Core{}, errors.New("id task yang diinput bukan milik anda")
	}

	taskCurrentCore := task.Core{
		ID:              allTaskCurrent.ID,
		ProjectID:       allTaskCurrent.ProjectID,
		TaskName:        allTaskCurrent.TaskName,
		DescriptionTask: allTaskCurrent.DescriptionTask,
		StatusTask:      allTaskCurrent.StatusTask,
		CreatedAt:       allTaskCurrent.CreatedAt,
		UpdatedAt:       allTaskCurrent.UpdatedAt,
	}

	return taskCurrentCore, nil

}

// Update implements task.DataInterface.
func (t *taskQuery) Update(id uint, idUser uint, input task.Core) error {
	var allTaskCurrent Task
	tx := t.db.First(&allTaskCurrent, id)
	if tx.Error != nil {
		return tx.Error
	}

	var projectForTask ProjectBowl
	tx2 := t.db.Model(&data.Project{}).Where("id = ?", allTaskCurrent.ProjectID).First(&projectForTask)
	if tx2.Error != nil {
		return tx2.Error
	}

	if projectForTask.UserID != idUser {
		return errors.New("id task yang diinput bukan milik anda")
	}

	tx3 := t.db.Model(&Task{}).Where("id = ?", id).Updates(input)
	if tx3.Error != nil {
		return tx3.Error
	}
	return nil
}

// Delete implements task.DataInterface.
func (t *taskQuery) Delete(id uint, idUser uint) error {
	var allTaskCurrent Task
	tx := t.db.First(&allTaskCurrent, id)
	if tx.Error != nil {
		return tx.Error
	}

	var projectForTask ProjectBowl
	tx2 := t.db.Model(&data.Project{}).Where("id = ?", allTaskCurrent.ProjectID).First(&projectForTask)
	if tx2.Error != nil {
		return tx2.Error
	}

	if projectForTask.UserID != idUser {
		return errors.New("id task yang diinput bukan milik anda")
	}

	tx3 := t.db.Delete(&Task{}, id)
	if tx3.Error != nil {
		return tx3.Error
	}
	return nil
}
