package domain

import (
	"encoding/json"
	"time"
)

type ByNoGameCsItem struct {
	SellerMarketName string          `json:"sellerMarketName"`
	Description      string          `json:"description"`
	ViewCounter      int             `json:"viewCounter"`
	SellerId         string          `json:"sellerId"`
	Price            float64         `json:"price"`
	Image            string          `json:"image"`
	NameSlug         string          `json:"nameSlug"`
	SellerLogo       string          `json:"sellerLogo"`
	Bngid            int             `json:"bngid"`
	CommentCount     int             `json:"commentCount"`
	TypeInfoSteam    CsItemInfoSteam `json:"typeInfoSteam"`
	GameCode         string          `json:"gameCode"`
	SellerOnline     bool            `json:"sellerOnline"`
	Status           string          `json:"status"`
	CreatedAt        CustomTime      `json:"createdAt"`
	DateTimeSold     CustomTime      `json:"dateTimeSold"`
}

type CsItemInfoSteam struct {
	App          int     `json:"app"`
	MarketPrice  float64 `json:"marketPrice"`
	StatTrak     bool    `json:"statTrak"`
	Exterior     string  `json:"exterior"`
	HashShort    string  `json:"hashShort"`
	ItemSteamId  string  `json:"ItemSteamId"`
	Stickers     string  `json:"stickers"`
	Type         string  `json:"type"`
	HashBare     string  `json:"hashBare"`
	Hash         string  `json:"hash"`
	StickersText string  `json:"stickersText"`
}

type CustomTime struct {
	time.Time
}

func (u *CustomTime) UnmarshalJSON(b []byte) error {
	var timestamp int64
	err := json.Unmarshal(b, &timestamp)
	if err == nil {
		u.Time = time.UnixMilli(timestamp)
		return nil
	} else {
		var timestamp string
		err = json.Unmarshal(b, &timestamp)
		if err != nil {
			return err
		}

		u.Time, err = time.Parse(time.RFC3339, timestamp)
		if err != nil {
			return err
		}
	}

	return nil
}
