package controller

import (
	"fmt"
	"log"
	"time"

	"github.com/jamemyjamess/go-clean-architecture-demo/configs/responseConfig"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func (sessionController *sessionController) SetSession(c echo.Context) error {
	now := time.Now()
	log.Println(fmt.Sprintf("session|%s", c.Param("id")))
	token := ""
	sess, err := session.Get(fmt.Sprintf("session|%s", c.Param("id")), c)
	if err != nil {
		log.Println(err.Error())
		return responseConfig.Handler(c).InternalServerErrorSetError(err)
	}
	if _, ok := sess.Values["token"]; !ok {
		log.Println("does note exist then create sess.Values[\"token\"]")
		sess.Options.MaxAge = now.Add(time.Second * time.Duration(10)).Second() // 10s
		sess.Values["token"] = "mock token"
		sess.Save(c.Request(), c.Response())
	}
	token = sess.Values["token"].(string)
	return responseConfig.Handler(c).Success(map[string]interface{}{
		"token": token,
	})
}

func (sessionController *sessionController) GetSession(c echo.Context) error {
	log.Println(fmt.Sprintf("session|%s", c.Param("id")))
	sess, err := session.Get(fmt.Sprintf("session|%s", c.Param("id")), c)
	if err != nil {
		log.Println(err.Error())
		return responseConfig.Handler(c).InternalServerErrorSetError(err)
	}
	token := ""
	if v, ok := sess.Values["token"]; ok {
		token = v.(string)
	}
	return responseConfig.Handler(c).Success(map[string]interface{}{
		"token": token,
	})
}
