package web_service

import (
	"fmt"
	"os"
	"path"

	"github.com/avequa/golang-todo-app/internal/core/domain"
)

func (s *WebService) GetMainPage() (domain.File, error) {
	htmlFilePath := path.Join(
		os.Getenv("PROJECT_ROOT"),
		"/public/index.html",
	)

	htmlFile, err := s.webRepository.GetFile(htmlFilePath)
	if err != nil {
		return domain.File{}, fmt.Errorf("get file from repository: %w", err)
	}

	return htmlFile, nil
}
