package tasks_service

import (
	"context"

	"github.com/avequa/golang-todo-app/internal/core/domain"
)

type TasksRepository interface {
	CreateTask(
		ctx context.Context,
		task domain.Task,
	) (domain.Task, error)
}

type TasksService struct {
	tasksRepository TasksRepository
}

func NewTasksService(
	tasksRepository TasksRepository,
) *TasksService {
	return &TasksService{
		tasksRepository: tasksRepository,
	}
}
