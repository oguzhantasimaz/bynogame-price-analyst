package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/oguzhantasimaz/bynogame-price-analyst/consumer/domain"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

const (
	SteamApiUrl = "https://steamcommunity.com/market/priceoverview"
	CsAppId     = 730
	Currency    = 17 //Turkish Lira

	SteamWebApiUrl      = "https://www.steamwebapi.com/steam/api/item"
	SteamWebApiCurrency = "TRY"
)

func GetItemFromSteamWebApi(hashedName, key string) (*domain.SteamWebApi, error) {
	var steamResp *domain.SteamWebApi
	url := fmt.Sprintf("%s?%s=%s&%s=%s&%s=%s", SteamWebApiUrl, "key", key, "market_hash_name", hashedName, "currency", SteamWebApiCurrency)

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

	if err := json.Unmarshal(body, &steamResp); err != nil {
		log.Error(err)
		return nil, err
	}

	return steamResp, nil
}

//func GetItemSteamMarketPrice(hashedName string) (*domain.SteamResponse, error) {
//	var steamResp *domain.SteamResponse
//
//	url := fmt.Sprintf("%s?%s=%d&%s=%d&%s=%s", SteamApiUrl, "appid", CsAppId, "currency", Currency, "market_hash_name", hashedName)
//
//	req, err := http.NewRequest("GET", url, nil)
//	if err != nil {
//		log.Error(err)
//		return nil, err
//	}
//
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if err != nil {
//		log.Error(err)
//		return nil, err
//	}
//
//	var body []byte
//	body, err = io.ReadAll(resp.Body)
//	if err != nil {
//		log.Error(err)
//		return nil, err
//	}
//
//	if err := json.Unmarshal(body, &steamResp); err != nil {
//		return nil, err
//	}
//
//	return steamResp, nil
//}
