package usecase

import (
	"context"
	"github.com/oguzhantasimaz/bynogame-price-analyst/core/domain"
	"github.com/oguzhantasimaz/bynogame-price-analyst/core/repository"
	log "github.com/sirupsen/logrus"
)

type csItemUseCase struct {
	csItemRepository repository.CsItemRepository
}

func NewCsItemUseCase(csItemRepository repository.CsItemRepository) domain.CsItemUseCase {
	return &csItemUseCase{
		csItemRepository: csItemRepository,
	}
}

func (cu *csItemUseCase) GetCsItems(ctx context.Context, limit int) ([]*domain.CsItem, error) {
	filter := map[string]interface{}{
		"limit": limit,
	}

	items, err := cu.csItemRepository.GetCsItems(ctx, filter)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	log.Info("ITEMS IN USECASE: ", items)

	return items, nil
}

func (cu *csItemUseCase) PostCsItem(ctx context.Context, item *domain.CsItem) error {
	err := cu.csItemRepository.PostCsItem(ctx, item)
	if err != nil {
		return err
	}

	return nil
}
