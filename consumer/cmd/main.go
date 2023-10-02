package main

import (
	"context"
	"encoding/json"
	"github.com/oguzhantasimaz/bynogame-price-analyst/consumer/bootstrap"
	"github.com/oguzhantasimaz/bynogame-price-analyst/consumer/domain"
	"github.com/oguzhantasimaz/bynogame-price-analyst/consumer/pkg"
	"github.com/oguzhantasimaz/bynogame-price-analyst/consumer/util"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Println("Starting consumer...")
	env := bootstrap.NewEnv()
	ctx := context.Background()

	consumer := pkg.NewKafkaReader("cs", env.KafkaUsername, env.KafkaPassword, env.KafkaBrokerUrl)

	for {
		message, err := consumer.ReadMessage(ctx)
		if err != nil {
			log.Error(err)
			continue
		}

		log.Info("Message received")

		var bgCsItem *domain.ByNoGameCsItem
		if err := json.Unmarshal(message.Value, &bgCsItem); err != nil {
			log.Error(err)
			continue
		}

		hashedName := util.HashItemName(bgCsItem.TypeInfoSteam.Hash)

		sr, err := pkg.GetItemSteamMarketPrice(hashedName)
		if err != nil {
			log.Error(err)
			continue
		}

		createdAtTime, err := util.ParseTimeAsUnix(bgCsItem.CreatedAt)
		if err != nil {
			log.Error(err)
			continue
		}
		dateTimeSoldTime, err := util.ParseTimeAsUnix(bgCsItem.DateTimeSold)
		if err != nil {
			log.Error(err)
			continue
		}

		csItem := domain.CsItem{
			SellerMarketName: bgCsItem.SellerMarketName,
			Description:      bgCsItem.Description,
			ViewCounter:      bgCsItem.ViewCounter,
			SellerId:         bgCsItem.SellerId,
			Price:            bgCsItem.Price,
			Image:            bgCsItem.Image,
			NameSlug:         bgCsItem.NameSlug,
			SellerLogo:       bgCsItem.SellerLogo,
			Bngid:            bgCsItem.Bngid,
			CommentCount:     bgCsItem.CommentCount,
			TypeInfoSteam:    bgCsItem.TypeInfoSteam,
			GameCode:         bgCsItem.GameCode,
			SellerOnline:     bgCsItem.SellerOnline,
			Status:           bgCsItem.Status,
			CreatedAt:        createdAtTime,
			DateTimeSold:     dateTimeSoldTime,
			SteamStats:       *sr,
		}

		if err := pkg.PostToCoreService(csItem, env.CoreServiceUrl); err != nil {
			log.Error(err)
			continue
		}
	}
}
