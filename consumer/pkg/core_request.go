package pkg

import (
	"bytes"
	"encoding/json"
	"github.com/oguzhantasimaz/bynogame-price-analyst/consumer/domain"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const (
	coreServiceUrl      = "http://localhost:8080"
	coreServiceEndpoint = "/api/v1/cs"
)

func PostToCoreService(csItem domain.CsItem) error {
	url := coreServiceUrl + coreServiceEndpoint

	payload, err := json.Marshal(csItem)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		log.Error(err)
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return err
	}
	defer resp.Body.Close()

	log.Info("response Status:", resp.Status)
	
	return nil
}
