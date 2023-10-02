package repository

import (
	"context"
	"github.com/oguzhantasimaz/bynogame-price-analyst/core/domain"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type CsItemRepository interface {
	GetCsItems(ctx context.Context, filter interface{}) ([]*domain.CsItem, error)
	PostCsItem(ctx context.Context, item *domain.CsItem) error
}

type csItemRepository struct {
	*CoreRepository
}

func NewCsItemRepository(mongo *mongo.Database, cn string) CsItemRepository {
	return &csItemRepository{
		NewCoreRepository(mongo, nil, cn),
	}
}

func (r *csItemRepository) GetCsItems(ctx context.Context, filter interface{}) ([]*domain.CsItem, error) {
	cursor, err := r.Find(ctx, filter, nil)
	if err != nil {
		return nil, err
	}

	var csItems []*domain.CsItem

	for cursor.Next(ctx) {
		var csItem *domain.CsItem
		err := cursor.Decode(&csItem)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		csItems = append(csItems, csItem)
	}

	return csItems, nil

}

func (r *csItemRepository) PostCsItem(ctx context.Context, item *domain.CsItem) error {
	_, err := r.InsertOne(ctx, item)
	if err != nil {
		return err
	}

	return nil
}
