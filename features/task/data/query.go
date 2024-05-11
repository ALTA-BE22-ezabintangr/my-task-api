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

	tx2 := t.db.Create(&taskGorm)
	if tx2.Error != nil {
		return tx2.Error
	}

	return nil
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
	tx3 := t.db.Delete(&Task{}, id)
	if tx3.Error != nil {
		return tx3.Error
	}
	return nil
}

func (t *taskQuery) GetTaskById(id uint) (task.Core, error) {
	var allTaskCurrent Task
	tx := t.db.First(&allTaskCurrent, id)
	if tx.Error != nil {
		return task.Core{}, tx.Error
	}
	taskCurrentCore := task.Core{
		ID:        id,
		ProjectID: allTaskCurrent.ProjectID,
	}

	return taskCurrentCore, nil
}
