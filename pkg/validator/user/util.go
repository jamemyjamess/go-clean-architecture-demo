package user

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/jamemyjamess/go-clean-architecture-demo/module/user/entity/domain"
	_userRepository "github.com/jamemyjamess/go-clean-architecture-demo/module/user/repository"
)

func (u *userValidator) CheckUserIDUnique(ctx context.Context, structLV validator.StructLevel, repo _userRepository.UserRepository, ID string) (user *domain.User) {
	user, err := repo.Find(ctx, ID)
	if err == nil && user.ID == ID {
		structLV.ReportError(ID, "id", "ID", "unique", "")
	}
	return user
}
