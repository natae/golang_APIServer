package controller

import (
	"github.com/natae/golang/model"
	"github.com/natae/golang/cache"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"net/http"
)

func PostLogin(ctx echo.Context) error {
	req := new(model.ReqLogin)
	err := json.Unmarshal(ctx.Get("body").([]byte), req)
	if err != nil {
		fmt.Println("error:", err)
	}

	session := model.Session{
		Nickname: req.Nickname,
		AuthToken: strconv.FormatInt(time.Now().Unix(), 10),
		ExpireTime: time.Now().Add(time.Minute * 10),
	}
	jsonSession, _ := json.Marshal(session)

	isSuccess, err := cache.Set(req.Nickname, string(jsonSession) , time.Minute * 10)
	if !isSuccess {
		fmt.Println("error:", err)
		return ctx.JSON(http.StatusOK, model.ResLogin{Result: -1})
	}

	fmt.Println("handler2")
	ctx.Set("OK", "okokok")
	return ctx.JSON(http.StatusOK, model.ResLogin{Result: 0, Session: session})
}
