package input

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jamemyjamess/go-clean-architecture-demo/module/user/entity/domain"
	_userRepository "github.com/jamemyjamess/go-clean-architecture-demo/module/user/repository"
)

type UserCreateReq struct {
	ID        string    `json:"id" param:"id"`
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
} // @Name UserCreateInput

func MakeTestCreateInput() (req *UserCreateReq) {
	return &UserCreateReq{
		ID:   "test",
		Name: "test",
	}
}

func CreateReqToUserDomain(req *UserCreateReq) (user *domain.User) {
	return &domain.User{
		ID:   req.ID,
		Name: req.Name,
	}
}

// validation struct level
func UserCreateReqStructLevelValidation(structLV validator.StructLevel, repo _userRepository.UserRepository) {
	ctx := context.Background()
	user := structLV.Current().Interface().(UserCreateReq)

	CheckUserIDUnique(ctx, structLV, repo, user.ID)
}

func CheckUserIDUnique(ctx context.Context, structLV validator.StructLevel, repo _userRepository.UserRepository, ID string) (user *domain.User) {
	user, err := repo.Find(ctx, ID)
	if err == nil && user.ID == ID {
		structLV.ReportError(ID, "id", "ID", "unique", "")
	}
	return user
}
