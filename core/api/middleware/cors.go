package middleware

import "github.com/labstack/echo/v4"

// CORS is a middleware for echo framework to handle CORS requests.

var CORS = func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		return next(c)
	}
}
