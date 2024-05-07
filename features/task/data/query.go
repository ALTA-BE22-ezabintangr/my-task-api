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
		UserID:          input.UserID,
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

// SelectAll implements task.DataInterface.
func (t *taskQuery) SelectAll() ([]task.Core, error) {
	panic("unimplemented")
}

// Update implements task.DataInterface.
func (t *taskQuery) Update(id uint, input task.Core) error {
	panic("unimplemented")
}

// Delete implements task.DataInterface.
func (t *taskQuery) Delete(id uint) error {
	panic("unimplemented")
}
