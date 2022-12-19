package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type TaskRepository interface {
	GetTasks(ctx context.Context, id int) ([]entity.Task, error)
	StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error)
	GetTaskByID(ctx context.Context, id int) (entity.Task, error)
	GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error)
	UpdateTask(ctx context.Context, task *entity.Task) error
	DeleteTask(ctx context.Context, id int) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) GetTasks(ctx context.Context, id int) ([]entity.Task, error) {
	result := []entity.Task{}
	err := r.db.Where("user_id = ?", id).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil // TODO: replace this
}

func (r *taskRepository) StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error) {
	error := r.db.Create(&task)
	if error != nil {
		return 0, nil
	}
	return task.ID, nil // TODO: replace this
}

func (r *taskRepository) GetTaskByID(ctx context.Context, id int) (entity.Task, error) {
	result := entity.Task{}
	err := r.db.Where("id = ?", id).Find(&result).Error
	if err != nil {
		return entity.Task{}, err
	}

	return result, nil // TODO: replace this
}

func (r *taskRepository) GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error) {
	result := []entity.Task{}
	err := r.db.Where("category_id = ?", catId).Find(&result).Error
	if err != nil {
		return []entity.Task{}, err
	}

	return result, nil // TODO: replace this
}

func (r *taskRepository) UpdateTask(ctx context.Context, task *entity.Task) error {
	result := entity.Task{}
	err := r.db.Model(&result).Where("id = ?", task.ID).Updates(&result).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (r *taskRepository) DeleteTask(ctx context.Context, id int) error {
	result := entity.Task{}
	err := r.db.Where("id = ?", id).Delete(&result)
	if err != nil {
		return err.Error
	}
	return nil // TODO: replace this
}
