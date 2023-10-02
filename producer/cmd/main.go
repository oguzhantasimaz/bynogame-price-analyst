package main

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"github.com/oguzhantasimaz/bynogame-price-analyst/producer/bootstrap"
	"github.com/oguzhantasimaz/bynogame-price-analyst/producer/pkg"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Println("Producer started")
	env := bootstrap.NewEnv()
	log.Println("KAFKA_USERNAME: ", env.KafkaUsername)
	kwCS := pkg.NewKafkaWriter("cs", env.KafkaUsername, env.KafkaPassword, env.KafkaBrokerUrl)
	defer kwCS.Close()

	cronClient := pkg.NewCronClient()
	cronClient.Schedule("5m", func() {
		ProcessCsItems(kwCS)
	})

	cronClient.Start()
	defer cronClient.Stop()

	select {}

}

func ProcessCsItems(kw *pkg.KafkaWriter) {
	listings, err := pkg.GetBynoGameCsListings()
	if err != nil {
		log.Error(err)
		return
	}

	if len(*listings) == 0 || listings == nil {
		log.Error("No listings found")
		return
	}

	for _, listing := range *listings {
		bs := make([]byte, 4)
		binary.LittleEndian.PutUint32(bs, uint32(listing.Bngid))

		listingAsBytes, err := json.Marshal(listing)
		if err != nil {
			log.Error(err)
		}

		err = kw.WriteMessages(context.Background(), kafka.Message{
			Key:   bs,
			Value: listingAsBytes,
		})
		if err != nil {
			log.Error(err)
		}
		log.Info("Message sent to Kafka")
	}
}
