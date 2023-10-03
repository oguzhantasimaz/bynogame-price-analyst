package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/oguzhantasimaz/bynogame-price-analyst/producer/domain"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

const (
	byNoGameUrl          = "https://integration.bynogame.com/api/listing/last/"
	byNoGameGameNameCs   = "csgo"
	byNoGameGameNameDota = "dota2"
	batchSize            = 5
	date                 = 0
)

func GetBynoGameCsListings() (*[]domain.ByNoGameCsItem, error) {
	var bynogameResp *domain.ByNoGameCsResponse

	url := fmt.Sprintf("%s%s?%s=%d&%s=%d", byNoGameUrl, byNoGameGameNameCs, "size", batchSize, "date", date)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	var body []byte
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	if err := json.Unmarshal(body, &bynogameResp); err != nil {
		log.Error(err)
		return nil, err
	}

	return &bynogameResp.Response.Data, nil
}
