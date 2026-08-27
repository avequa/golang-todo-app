package web_service

import "github.com/avequa/golang-todo-app/internal/core/domain"

type WebService struct {
	webRepository WebRepository
}

type WebRepository interface {
	GetFile(filePath string) (domain.File, error)
}

func NewWebService(
	webRepository WebRepository,
) *WebService {
	return &WebService{
		webRepository: webRepository,
	}
}
