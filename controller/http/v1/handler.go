package handler

import (
	"github.com/jamemyjamess/go-clean-architecture-demo/pkg/database"
	"github.com/labstack/echo/v4"

	_companyHttp "github.com/jamemyjamess/go-clean-architecture-demo/module/company/controller"
	_companyRepository "github.com/jamemyjamess/go-clean-architecture-demo/module/company/repository/postgres"
	_companyUsecase "github.com/jamemyjamess/go-clean-architecture-demo/module/company/usecase"
	_userHttp "github.com/jamemyjamess/go-clean-architecture-demo/module/user/controller"
	_userRepository "github.com/jamemyjamess/go-clean-architecture-demo/module/user/repository/postgres"
	_userUsecase "github.com/jamemyjamess/go-clean-architecture-demo/module/user/usecase"
)

type Handler struct {
	// Echo *echo.Echo
	// Cfg *configs.Configs
	// Db *gorm.DB
	company _companyHttp.CompanyController
	user    _userHttp.UserController
}

func NewHandler() *Handler {

	companyRepository := _companyRepository.NewCompanyRepository(database.PostgresSql)
	companyUsecase := _companyUsecase.NewCompanyUsecase(companyRepository)

	userRepository := _userRepository.NewUserRepository(database.PostgresSql)
	userUsecase := _userUsecase.NewUserUsecase(userRepository, companyUsecase)

	return &Handler{
		company: _companyHttp.NewCompanyHandler(companyUsecase),
		user:    _userHttp.NewUserHandler(userUsecase),
	}
}

func (handler Handler) NewRouter(e *echo.Echo) {
	e.Static("public", "assets/public") // to config
	v1 := e.Group("/api/v1")
	handler.mapUserRouteHandler(v1)
	handler.mapCompanyRouteHandler(v1)
}
