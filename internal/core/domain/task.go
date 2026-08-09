package domain

import "time"

type Task struct {
	ID      int
	Version int

	Title       string
	Description *string

	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time

	AuthorUserID int
}

func NewTask(
	id int,
	version int,

	title string,
	description *string,

	completed bool,
	createdAt time.Time,
	completedAt *time.Time,

	authorUserID int,
) Task {
	return Task{
		ID:      id,
		Version: version,

		Title:       title,
		Description: description,

		Completed:   completed,
		CreatedAt:   createdAt,
		CompletedAt: completedAt,

		AuthorUserID: authorUserID,
	}
}

func NewTaskUninitialized(
	title string,
	description *string,
	authorUserID int,
) Task {
	return NewTask(
		UninitializedID,
		UninitializedVersion,
		title,
		description,
		false,
		time.Now(),
		nil,
		authorUserID,
	)
}
