package user

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/jamemyjamess/go-clean-architecture-demo/module/user/entity/input"

	_userRepository "github.com/jamemyjamess/go-clean-architecture-demo/module/user/repository"
)

type userValidator struct {
	validate *validator.Validate
	repo     _userRepository.UserRepository
}

func NewUserValidator(validate *validator.Validate, repo _userRepository.UserRepository) *userValidator {
	return &userValidator{
		validate: validate,
		repo:     repo,
	}
}

func (u *userValidator) Init() {
	u.RegisterStructValidation()
}

func (u *userValidator) RegisterStructValidation() {
	u.validate.RegisterStructValidation(u.UserCreateReqStructLevelValidation, &input.UserCreateReq{})
}

func (u *userValidator) UserCreateReqStructLevelValidation(sl validator.StructLevel) {
	ctx := context.Background()
	user := sl.Current().Interface().(input.UserCreateReq)

	// validation step
	u.CheckUserIDUnique(ctx, sl, u.repo, user.ID)
	// func.....
}
