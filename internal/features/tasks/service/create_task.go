package tasks_service

import (
	"context"

	"github.com/avequa/golang-todo-app/internal/core/domain"
)

func (s *TasksService) CreateTask(
	ctx context.Context,
	task domain.Task,
) (domain.Task, error) {
	// 1. task.Validate()
	// 2. new := repo.Save(task)
	// 3. return new
}
