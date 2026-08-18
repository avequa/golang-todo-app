package tasks_service

import (
	"context"
	"fmt"

	"github.com/avequa/golang-todo-app/internal/core/domain"
	core_errors "github.com/avequa/golang-todo-app/internal/core/errors"
)

func (s *TasksService) GetTasks(
	ctx context.Context,
	userID *int,
	limit *int,
	offset *int,
) ([]domain.Task, error) {

	// 1. validate()
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf(
			"limit must be non-negative",
			core_errors.ErrInvalidArgument,
		)
	}

	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf(
			"offset must be non-negative",
			core_errors.ErrInvalidArgument,
		)
	}

	// 2. get_tasks()
	tasks, err := s.tasksRepository.GetTasks(ctx, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get tasks from repo: %w", err)

	}

	return tasks, nil
}
