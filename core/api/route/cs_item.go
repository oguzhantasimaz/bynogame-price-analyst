package route

import (
	"github.com/labstack/echo/v4"
	"github.com/oguzhantasimaz/bynogame-price-analyst/core/api/controller"
	"github.com/oguzhantasimaz/bynogame-price-analyst/core/domain"
	"github.com/oguzhantasimaz/bynogame-price-analyst/core/repository"
	"github.com/oguzhantasimaz/bynogame-price-analyst/core/usecase"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewCsItemRouter(db *mongo.Database, group *echo.Group) {
	cr := repository.NewCsItemRepository(db, domain.CollectionCsItem)
	cu := &controller.CsItemController{
		CsItemUseCase: usecase.NewCsItemUseCase(cr),
	}

	group.GET("/cs", cu.GetCsItems)
	group.POST("/cs", cu.PostCsItem)
}
