package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/oguzhantasimaz/bynogame-price-analyst/consumer/domain"
	"io"
	"net/http"
)

const (
	SteamApiUrl = "https://steamcommunity.com/market/priceoverview/"
	CsAppId     = 730
	Currency    = 17 //Turkish Lira
)

func GetItemSteamMarketPrice(hashedName string) (*domain.SteamResponse, error) {
	var steamResp *domain.SteamResponse

	url := fmt.Sprintf("%s?%s=%d&%s=%d&%s=%s", SteamApiUrl, "appid", CsAppId, "currency", Currency, "market_hash_name", hashedName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var body []byte
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &steamResp); err != nil {
		return nil, err
	}

	return steamResp, nil
}
