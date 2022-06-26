package controllers

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"shape/entities"
	"shape/entities/request"
	"shape/entities/response"
)

var (
	ErrMissingUsername = errors.New("Username is missing.")
)

func apiLogin(s *Controller) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		var resp response.Response
		//Read body to get json
		body, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			log.Error("Error on read body request", err)
			return c.JSON(http.StatusBadRequest, resp.Error(err, 400))
		}
		var user request.UserRequest
		err = json.Unmarshal(body, &user)
		if err != nil {
			log.Error("Error on unmarshal body to struct", err)
			return c.JSON(http.StatusInternalServerError, resp.Error(err, 500))
		}
		if client, ok := s.defender.Client(c.RealIP()); ok && !client.Banned() {
			u, err := s.auth.AuthenticateJWT(user.Username, user.Password)
			if err != nil {
				log.Error("Error on authorize jwt", err)
				if s.defender.Inc(c.RealIP()) {
					log.Error("IP ", c.RealIP(), " has been warning!")
					return c.JSON(http.StatusUnauthorized, resp.Error(entities.ErrOnLogin, 401))
				}
				return c.JSON(http.StatusUnauthorized, resp.Error(err, 401))
			}
			log.Info(user.Username, " login success", ". Request: ", c.Request().Referer())
			return c.JSON(http.StatusOK, resp.Success(u))

		}
		return c.JSON(http.StatusUnauthorized, resp.Error(entities.ErrOnLogin, 401))
	})
}
func apiNewUser(s *Controller) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		var resp response.Response

		body, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			log.Error("Error on read body request", err)

			return c.JSON(http.StatusBadRequest, resp.Error(err, 400))
		}
		var user request.UserRequest
		err = json.Unmarshal(body, &user)
		if err != nil {
			log.Error("Error on unmarshal body to struct", err)
			return c.JSON(http.StatusInternalServerError, resp.Error(err, 500))
		}
		//validate struct
		validate := validator.New()
		err = validate.Struct(user)
		if err != nil {
			if _, ok := err.(*validator.InvalidValidationError); ok {
				log.Error("Error on validate struct", err)
				return c.JSON(http.StatusInternalServerError, resp.Error(err, 500))
			}
		}

		if user.Username != "" {
			newuser, err := s.auth.DB.NewUser(user)
			if err != nil {
				log.Error("Error on create new auth user", err)
				return c.JSON(http.StatusUnprocessableEntity, resp.Error(err, 422))
			}
			log.Info("Create new auth user success")
			return c.JSON(http.StatusOK, resp.Success(newuser))
		}
		log.Error("Error on create new auth user")
		return c.JSON(http.StatusBadRequest, resp.Error(ErrMissingUsername, 422))
	})
}

func apiDeleteUser(s *Controller) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		var resp response.Response

		id := c.Param("id")
		currentUser, err := s.auth.DB.DisabledUser(id)
		if err != nil {
			log.Error("Error on delete user ", err)

			return c.JSON(http.StatusUnprocessableEntity, resp.Error(err, 422))
		}
		log.Info("Delete auth user success")

		return c.JSON(http.StatusOK, resp.Success(currentUser))
	})
}

func apiGetUsers(s *Controller) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		var resp response.Response
		if c.Param("id") == "" {
			username := c.QueryParam("username")
			if username == "" {
				users, err := s.auth.DB.GetsPublic()
				if err != nil {
					log.Error("Error on get user", err)
					return c.JSON(http.StatusUnprocessableEntity, resp.Error(err, 422))
				}
				log.Info("Get user success", users)
				return c.JSON(http.StatusOK, resp.Success(users))
			}
			user, err := s.auth.DB.GetPublic(username)
			if err != nil {
				log.Error("Error on get user", err)
				return c.JSON(http.StatusUnprocessableEntity, resp.Error(err, 422))
			}
			log.Info("Get user success", user)
			return c.JSON(http.StatusOK, resp.Success(user))
		}
		id := c.Param("id")
		user, err := s.auth.DB.GetUserByID(id)
		user.Password = ""
		if err != nil {
			log.Error("Error on get user", err)
			return c.JSON(http.StatusUnprocessableEntity, resp.Error(err, 422))
		}
		log.Info("Get user success")
		return c.JSON(http.StatusOK, resp.Success(user))
	})
}

func apiUpdateUser(s *Controller) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		var resp response.Response

		body, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			log.Error("Error on read body request", err)
			return c.JSON(http.StatusBadRequest, resp.Error(err, 400))
		}
		var user request.UserRequest
		err = json.Unmarshal(body, &user)
		if err != nil {
			log.Error("Error on unmarshal body to struct", err)
			return c.JSON(http.StatusInternalServerError, resp.Error(err, 500))
		}
		id := c.Param("id")
		err = s.auth.DB.UpdateUser(id, user)
		if err != nil {
			log.Error("Error on update user", err)
			return c.JSON(http.StatusUnprocessableEntity, resp.Error(err, 422))
		}
		log.Info("Update user success")
		return c.JSON(http.StatusOK, resp.Success("Update user success"))
	})
}
