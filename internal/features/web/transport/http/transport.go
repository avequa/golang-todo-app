package web_transport_http

import (
	"github.com/avequa/golang-todo-app/internal/core/domain"
	core_http_server "github.com/avequa/golang-todo-app/internal/core/transport/http/server"
)

type WebService interface {
	GetMainPage() (domain.File, error)
}

type WebHTTPHandler struct {
	webService WebService
}

func NewWebHTTPHandler(
	webService WebService,
) *WebHTTPHandler {
	return &WebHTTPHandler{
		webService: webService,
	}
}

func (h *WebHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Path:    "/",
			Handler: h.GetMainPage,
		},
	}
}
