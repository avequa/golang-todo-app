package tasks_transport

import core_http_server "github.com/avequa/golang-todo-app/internal/core/transport/http/server"

type TasksService interface {
}

type TasksHTTPHandler struct {
	tasksService TasksService
}

func NewTasksHTTPHandler(
	tasksService TasksService,
) *TasksHTTPHandler {
	return &TasksHTTPHandler{
		tasksService: tasksService,
	}
}

func (h *TasksHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{}
}
