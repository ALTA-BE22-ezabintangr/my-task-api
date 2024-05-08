package data

import (
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
	var taskGorm Task

	taskGorm = Task{
		Model:           gorm.Model{},
		ProjectID:       input.ProjectID,
		TaskName:        input.TaskName,
		DescriptionTask: input.DescriptionTask,
		StatusTask:      input.StatusTask,
	}

	tx := t.db.Create(&taskGorm)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// GetTaskbyUserId implements task.DataInterface.
func (t *taskQuery) GetTaskById(id uint) (task.Core, error) {
	var allTaskCurrent Task
	tx := t.db.First(&allTaskCurrent, id)
	if tx.Error != nil {
		return task.Core{}, tx.Error
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
func (t *taskQuery) Update(id uint, input task.Core) error {
	tx := t.db.Model(&Task{}).Where("id = ?", id).Updates(input)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Delete implements task.DataInterface.
func (t *taskQuery) Delete(id uint) error {
	tx := t.db.Delete(&Task{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
