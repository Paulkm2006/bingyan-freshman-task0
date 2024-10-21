package controller

import (
	"bingyan-freshman-task0/internal/config"
	"bingyan-freshman-task0/internal/controller/param"
	"bingyan-freshman-task0/internal/model"
	"bingyan-freshman-task0/internal/utils"
	"crypto/md5"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func UserLogin(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return echo.ErrBadRequest
	}
	result, err := model.GetUser(user.Username)
	if err != nil {
		return echo.ErrUnauthorized
	}
	if result.Password != fmt.Sprintf("%x", md5.Sum([]byte(user.Password))) {
		return echo.ErrUnauthorized
	}
	claims := utils.JWTClaims{
		UID:   result.ID,
		Admin: true,
		Exp:   config.Config.Jwt.Expire + jwt.TimeFunc().Unix(),
	}
	token, err := utils.GenerateToken(claims)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(200, &param.Resp{
		Success: true,
		Data: &param.TokenResponse{
			Token:   token,
			Expires: int(config.Config.Jwt.Expire),
		},
	})
}

func UserRegister(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return echo.ErrBadRequest
	}
	user.Password = fmt.Sprintf("%x", md5.Sum([]byte(user.Password)))
	if err := model.AddUser(&user); err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(200, &param.Resp{
		Success: true,
	})
}
