package controllers

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

var (
	ErrNoAuth           = errors.New("Authorization header is missing")
	ErrNoToken          = errors.New("Token is missing")
	ErrInvalidTokenType = errors.New("Authorization Bearer is required")
)

func Authorization(s *Controller) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth := c.Request().Header.Get("Authorization")
			if len(auth) < 1 {
				log.Error("Can not get header or header is nil")
				return c.JSON(http.StatusUnauthorized, echo.Map{"message": ErrNoAuth.Error()})
			}
			if strings.HasPrefix(auth, "Bearer ") {
				if parts := strings.Split(auth, "Bearer "); len(parts) > 1 {
					token := parts[1]
					t, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
						return s.secret, nil
					})
					if err != nil {
						log.Error("Can not parse header", err)
						return c.JSON(http.StatusUnauthorized, echo.Map{"message": err.Error()})
					}
					c.Set("info", t.Claims.(jwt.MapClaims))
					log.Info("Get token success")
					c.Response().Header().Set("Keep-Alive", "Timeout = 10, max = 1000")

					return next(c)
				}
				log.Error("Error on invalid token type")

				return c.JSON(http.StatusUnauthorized, echo.Map{"message": ErrInvalidTokenType.Error()})
			}
			log.Error("Error on no token")
			return c.JSON(http.StatusUnauthorized, echo.Map{"message": ErrNoToken.Error()})
		}
	}
}
