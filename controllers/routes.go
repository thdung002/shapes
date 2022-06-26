package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func (ctrl *Controller) loadMux() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(1000)))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:          middleware.DefaultSkipper,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))

	apiv1 := e.Group("/api")
	shapev1 := apiv1.Group("/shape/v1")
	shapev1.POST("/login", apiLogin(ctrl))
	shapev1.POST("/register", apiNewUser(ctrl))

	shapev1.Use(Authorization(ctrl))
	{
		user := shapev1.Group("/user")
		{
			user.GET("", apiGetUsers(ctrl))
			user.GET("/:id", apiGetUsers(ctrl))
			user.PUT("/:id", apiUpdateUser(ctrl))
			user.DELETE("/:id", apiDeleteUser(ctrl))
		}
	}

	{
		triangle := shapev1.Group("/triangle")
		{
			triangle.GET("", apiGetTriangle(ctrl))
			triangle.GET("/area", apiGetTriangleArea(ctrl))
			triangle.GET("/perimeter", apiGetTrianglePerimeter(ctrl))
			triangle.PUT("", apiPutTriangle(ctrl))
			triangle.POST("", apiPostTriangle(ctrl))
			triangle.DELETE("", apiDeleteTriangle(ctrl))

		}
	}
	{
		square := shapev1.Group("/square")
		{
			square.GET("", apiGetSquare(ctrl))
			square.GET("/area", apiGetSquareArea(ctrl))
			square.GET("/perimeter", apiGetSquarePerimeter(ctrl))
			square.PUT("", apiPutSquare(ctrl))
			square.POST("", apiPostSquare(ctrl))
			square.DELETE("", apiDeleteSquare(ctrl))

		}
	}
	{
		rectangle := shapev1.Group("/rectangle")
		{
			rectangle.GET("", apiGetRectangle(ctrl))
			rectangle.GET("/area", apiGetRectangleArea(ctrl))
			rectangle.GET("/perimeter", apiGetRectanglePerimeter(ctrl))
			rectangle.PUT("", apiPutRectangle(ctrl))
			rectangle.POST("", apiPostRectangle(ctrl))
			rectangle.DELETE("", apiDeleteRectangle(ctrl))

		}
	}
	{
		diamond := shapev1.Group("/diamond")
		{
			diamond.GET("", apiGetDiamond(ctrl))
			diamond.GET("/area", apiGetDiamondArea(ctrl))
			diamond.GET("/perimeter", apiGetDiamondPerimeter(ctrl))
			diamond.PUT("", apiPutDiamond(ctrl))
			diamond.POST("", apiPostDiamond(ctrl))
			diamond.DELETE("", apiDeleteDiamond(ctrl))
		}
	}

	ctrl.e = e

}
