package users_service

import (
	"context"

	"github.com/avequa/golang-todo-app/internal/core/domain"
)

func (s *UsersService) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	// 1.user.Validate()
	// 2.repo.Save()

	//return user
}