package main

import (
	"context"
	"encoding/json"
	"github.com/oguzhantasimaz/bynogame-price-analyst/consumer/bootstrap"
	"github.com/oguzhantasimaz/bynogame-price-analyst/consumer/domain"
	"github.com/oguzhantasimaz/bynogame-price-analyst/consumer/pkg"
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

		//log.Info("Message received: ", string(message.Value))

		var bgCsItem *domain.ByNoGameCsItem
		if err := json.Unmarshal(message.Value, &bgCsItem); err != nil {
			log.Error(err)
			continue
		}

		log.Info("Item: ", bgCsItem)

		hashedName := pkg.HashItemName(bgCsItem.TypeInfoSteam.Hash)

		sr, err := pkg.GetItemSteamMarketPrice(hashedName)
		if err != nil {
			log.Error(err)
			continue
		}

		log.Println("SteamResponse: ", sr)

		// Merge SteamResponse and ByNoGameCsItem to CsItem
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
			CreatedAt:        bgCsItem.CreatedAt,
			DateTimeSold:     bgCsItem.DateTimeSold,
			SteamStats:       *sr,
		}

		log.Println("CsItem: ", csItem)

		//if err := pkg.PostToCoreService(csItem); err != nil {
		//	log.Error(err)
		//	continue
		//}
	}

	//hashedItem := pkg.HashItemName("★ Gut Knife | Scorched (Field-Tested)")
	//println(hashedItem)
	//
	//sr, err := pkg.GetItemSteamMarketPrice(hashedItem)
	//if err != nil {
	//	log.Error(err)
	//}
	//
	//log.Println("SteamResponse: ", sr)
}
