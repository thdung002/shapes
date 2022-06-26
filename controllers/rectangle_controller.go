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

func apiGetRectangle(s *Controller) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		var resp response.Response
		return c.JSON(http.StatusOK, resp.Success(map[string]interface{}{
			"rectangle_info": s.rectangle,
		}))
	})
}
func apiGetRectangleArea(s *Controller) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		var resp response.Response
		return c.JSON(http.StatusOK, resp.Success(map[string]interface{}{
			"rectangle_area": s.rectangle.Area(),
		}))
	})
}
func apiGetRectanglePerimeter(s *Controller) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		var resp response.Response
		return c.JSON(http.StatusOK, resp.Success(map[string]interface{}{
			"rectangle_perimeter": s.rectangle.Perimeter(),
		}))
	})
}

func apiPostRectangle(s *Controller) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		var resp response.Response
		body, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			log.Error("Error on read body request", err)
			return c.JSON(http.StatusBadRequest, resp.Error(err, 400))
		}
		rectangle := entities.Rectangle{}
		err = json.Unmarshal(body, &rectangle)
		if err != nil {
			log.Error("Error on unmarshal body to struct", err)
			return c.JSON(http.StatusInternalServerError, resp.Error(err, 500))
		}
		s.rectangle = entities.Rectangle{Length: rectangle.Length, Width: rectangle.Width}
		return c.JSON(http.StatusOK, resp.Success("Create rectangle data success"))
	})
}

func apiPutRectangle(s *Controller) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		var resp response.Response
		body, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			log.Error("Error on read body request", err)
			return c.JSON(http.StatusBadRequest, resp.Error(err, 400))
		}
		rectangle := entities.Rectangle{}
		err = json.Unmarshal(body, &rectangle)
		if err != nil {
			log.Error("Error on unmarshal body to struct", err)
			return c.JSON(http.StatusInternalServerError, resp.Error(err, 500))
		}
		s.rectangle = entities.Rectangle{Length: rectangle.Length, Width: rectangle.Width}
		return c.JSON(http.StatusOK, resp.Success("Update rectangle data success"))

	})
}

func apiDeleteRectangle(s *Controller) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		var resp response.Response
		s.rectangle = entities.Rectangle{}
		return c.JSON(http.StatusOK, resp.Success("Delete rectangle data success"))

	})
}
