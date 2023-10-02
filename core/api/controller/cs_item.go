package controller

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/oguzhantasimaz/bynogame-price-analyst/core/domain"
	log "github.com/sirupsen/logrus"
	"strconv"
)

type CsItemController struct {
	CsItemUseCase domain.CsItemUseCase
}

func (cc *CsItemController) GetCsItems(c echo.Context) error {
	ctx := context.Background()
	var err error

	limit := 2
	if c.QueryParam("limit") != "" {
		limit, err = strconv.Atoi(c.QueryParam("limit"))
		if err != nil {
			log.Error(err)
			return c.JSON(400, domain.ErrorResponse{Message: err.Error()})
		}
	}

	items, err := cc.CsItemUseCase.GetCsItems(ctx, limit)
	if err != nil {
		log.Error(err)
		return c.JSON(400, domain.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(200, items)

}

func (cc *CsItemController) PostCsItem(c echo.Context) error {
	ctx := context.Background()

	var item *domain.CsItem
	if err := c.Bind(&item); err != nil {
		log.Error(err)
		return c.JSON(400, err)
	}

	log.Info(item)

	err := cc.CsItemUseCase.PostCsItem(ctx, item)
	if err != nil {
		log.Error(err)
		return c.JSON(400, domain.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(200, "Success")
}
