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

func apiGetTriangle(s *Controller) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		var resp response.Response
		return c.JSON(http.StatusOK, resp.Success(map[string]interface{}{
			"triangle_info": s.triangle,
		}))
	})
}
func apiGetTriangleArea(s *Controller) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		var resp response.Response
		return c.JSON(http.StatusOK, resp.Success(map[string]interface{}{
			"triangle_area": s.triangle.Area(),
		}))
	})
}
func apiGetTrianglePerimeter(s *Controller) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		var resp response.Response
		return c.JSON(http.StatusOK, resp.Success(map[string]interface{}{
			"triangle_perimeter": s.triangle.Perimeter(),
		}))
	})
}

func apiPostTriangle(s *Controller) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		var resp response.Response
		body, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			log.Error("Error on read body request", err)
			return c.JSON(http.StatusBadRequest, resp.Error(err, 400))
		}
		triangle := entities.Triangle{}
		err = json.Unmarshal(body, &triangle)
		if err != nil {
			log.Error("Error on unmarshal body to struct", err)
			return c.JSON(http.StatusInternalServerError, resp.Error(err, 500))
		}
		s.triangle = entities.Triangle{FirstSide: triangle.FirstSide, SecondSide: triangle.SecondSide, ThirdSide: triangle.ThirdSide}
		return c.JSON(http.StatusOK, resp.Success("Create triangle data success"))
	})
}

func apiPutTriangle(s *Controller) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		var resp response.Response
		body, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			log.Error("Error on read body request", err)
			return c.JSON(http.StatusBadRequest, resp.Error(err, 400))
		}
		triangle := entities.Triangle{}
		err = json.Unmarshal(body, &triangle)
		if err != nil {
			log.Error("Error on unmarshal body to struct", err)
			return c.JSON(http.StatusInternalServerError, resp.Error(err, 500))
		}
		s.triangle = entities.Triangle{FirstSide: triangle.FirstSide, SecondSide: triangle.SecondSide, ThirdSide: triangle.ThirdSide}
		return c.JSON(http.StatusOK, resp.Success("Update triangle data success"))

	})
}

func apiDeleteTriangle(s *Controller) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		var resp response.Response
		s.triangle = entities.Triangle{}
		return c.JSON(http.StatusOK, resp.Success("Delete triangle data success"))

	})
}
