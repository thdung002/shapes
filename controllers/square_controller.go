package controllers

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"shape/entities"
	"shape/entities/response"
)

func apiGetSquare(s *Controller) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		var resp response.Response
		return c.JSON(http.StatusOK, resp.Success(map[string]interface{}{
			"square_info": s.square,
		}))
	})
}

func apiGetSquareArea(s *Controller) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		var resp response.Response
		return c.JSON(http.StatusOK, resp.Success(map[string]interface{}{
			"square_area": s.square.Area(),
		}))
	})
}

func apiGetSquarePerimeter(s *Controller) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		var resp response.Response
		return c.JSON(http.StatusOK, resp.Success(map[string]interface{}{
			"square_perimeter": s.square.Perimeter(),
		}))
	})
}

func apiPostSquare(s *Controller) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		var resp response.Response
		body, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			log.Error("Error on read body request", err)
			return c.JSON(http.StatusBadRequest, resp.Error(err, 400))
		}
		square := entities.Square{}
		err = json.Unmarshal(body, &square)
		if err != nil {
			log.Error("Error on unmarshal body to struct", err)
			return c.JSON(http.StatusInternalServerError, resp.Error(err, 500))
		}
		s.square = entities.Square{Length: square.Length}
		return c.JSON(http.StatusOK, resp.Success("Create square data success"))
	})
}

func apiPutSquare(s *Controller) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		var resp response.Response
		body, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			log.Error("Error on read body request", err)
			return c.JSON(http.StatusBadRequest, resp.Error(err, 400))
		}
		square := entities.Square{}
		err = json.Unmarshal(body, &square)
		if err != nil {
			log.Error("Error on unmarshal body to struct", err)
			return c.JSON(http.StatusInternalServerError, resp.Error(err, 500))
		}
		s.square = entities.Square{Length: square.Length}
		return c.JSON(http.StatusOK, resp.Success("Update square data success"))
	})
}

func apiDeleteSquare(s *Controller) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		var resp response.Response
		s.square = entities.Square{}
		return c.JSON(http.StatusOK, resp.Success("Delete square data success"))
	})
}
