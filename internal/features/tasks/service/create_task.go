package tasks_service

import (
	"context"
	"fmt"

	"github.com/avequa/golang-todo-app/internal/core/domain"
)

func (s *TasksService) CreateTask(
	ctx context.Context,
	task domain.Task,
) (domain.Task, error) {

	// 1. validate()
	if err := task.Validate(); err != nil {
		return domain.Task{}, fmt.Errorf("validate task domain: %w", err)
	}

	// 2. new := repo.Save(task)
	task, err := s.tasksRepository.CreateTask(ctx, task);
	if err != nil {
		return domain.Task{}, fmt.Errorf("create task: %w", err)
	}

	// 3. return new
	return task, nil
}
