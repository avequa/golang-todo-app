package tasks_service

type TasksRepository interface {
	//
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
