package route

import (
	"github.com/labstack/echo/v4"
	"github.com/oguzhantasimaz/bynogame-price-analyst/core/api/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(db *mongo.Database, echo *echo.Echo) {
	route := echo.Group("/api")

	route.Use(middleware.CORS)

	NewCsItemRouter(db, route)
}
