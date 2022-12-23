package validator

import (
	"github.com/go-playground/validator/v10"
	_companyRepository "github.com/jamemyjamess/go-clean-architecture-demo/module/company/repository"
	_userRepository "github.com/jamemyjamess/go-clean-architecture-demo/module/user/repository"
	_userValidation "github.com/jamemyjamess/go-clean-architecture-demo/pkg/validator/user"
)

type GoPlayGroundValidator struct {
	validate    *validator.Validate
	companyRepo _companyRepository.CompanyRepository
	userRepo    _userRepository.UserRepository
}

func New(companyRepo _companyRepository.CompanyRepository, userRepo _userRepository.UserRepository) (v *GoPlayGroundValidator) {
	v = &GoPlayGroundValidator{
		validate:    validator.New(),
		companyRepo: companyRepo,
		userRepo:    userRepo,
	}

	userValidation := _userValidation.NewUserValidator(v.validate, v.userRepo)
	userValidation.Init()
	return v
}

func (v *GoPlayGroundValidator) Validate(item interface{}) (err error) {
	return v.validate.Struct(item)
}
