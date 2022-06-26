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

func apiGetDiamond(s *Controller) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		var resp response.Response
		return c.JSON(http.StatusOK, resp.Success(map[string]interface{}{
			"diamond_info": s.diamond,
		}))
	})
}
func apiGetDiamondArea(s *Controller) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		var resp response.Response
		return c.JSON(http.StatusOK, resp.Success(map[string]interface{}{
			"diamond_area": s.diamond.Area(),
		}))
	})
}
func apiGetDiamondPerimeter(s *Controller) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		var resp response.Response
		return c.JSON(http.StatusOK, resp.Success(map[string]interface{}{
			"diamond_perimeter": s.diamond.Perimeter(),
		}))
	})
}

func apiPostDiamond(s *Controller) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		var resp response.Response
		body, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			log.Error("Error on read body request", err)
			return c.JSON(http.StatusBadRequest, resp.Error(err, 400))
		}
		diamond := entities.Diamond{}
		err = json.Unmarshal(body, &diamond)
		if err != nil {
			log.Error("Error on unmarshal body to struct", err)
			return c.JSON(http.StatusInternalServerError, resp.Error(err, 500))
		}
		s.diamond = entities.Diamond{Length: diamond.Length, Height: diamond.Height, Side: diamond.Side}
		return c.JSON(http.StatusOK, resp.Success("Create diamond data success"))
	})
}

func apiPutDiamond(s *Controller) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		var resp response.Response
		body, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			log.Error("Error on read body request", err)
			return c.JSON(http.StatusBadRequest, resp.Error(err, 400))
		}
		diamond := entities.Diamond{}
		err = json.Unmarshal(body, &diamond)
		if err != nil {
			log.Error("Error on unmarshal body to struct", err)
			return c.JSON(http.StatusInternalServerError, resp.Error(err, 500))
		}
		s.diamond = entities.Diamond{Length: diamond.Length, Height: diamond.Height, Side: diamond.Side}
		return c.JSON(http.StatusOK, resp.Success("Update diamond data success"))

	})
}
func apiDeleteDiamond(s *Controller) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		var resp response.Response
		s.diamond = entities.Diamond{}
		return c.JSON(http.StatusOK, resp.Success("Delete diamond data success"))

	})
}
