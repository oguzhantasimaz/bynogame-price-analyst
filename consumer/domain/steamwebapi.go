package domain

import "time"

type SteamWebApi struct {
	Id              string      `json:"id"`
	Markethashname  string      `json:"markethashname"`
	Marketname      string      `json:"marketname"`
	Hashid          string      `json:"hashid"`
	Type            string      `json:"type"`
	Rarity          interface{} `json:"rarity"`
	Quality         string      `json:"quality"`
	Pricelatest     float64     `json:"pricelatest"`
	Pricemedian     float64     `json:"pricemedian"`
	Pricesafe       float64     `json:"pricesafe"`
	Priceavg        float64     `json:"priceavg"`
	Pricemin        float64     `json:"pricemin"`
	Pricemax        float64     `json:"pricemax"`
	Pricesafe24H    float64     `json:"pricesafe24h"`
	Pricesafe7D     float64     `json:"pricesafe7d"`
	Pricesafe30D    float64     `json:"pricesafe30d"`
	Pricesafe90D    float64     `json:"pricesafe90d"`
	Pricereal       float64     `json:"pricereal"`
	Pricemix        float64     `json:"pricemix"`
	Sold24H         int         `json:"sold24h"`
	Sold7D          int         `json:"sold7d"`
	Sold30D         int         `json:"sold30d"`
	Sold90D         int         `json:"sold90d"`
	Slug            string      `json:"slug"`
	Image           string      `json:"image"`
	Wear            string      `json:"wear"`
	Createdat       time.Time   `json:"createdat"`
	Stattrack       bool        `json:"stattrack"`
	Dailysoldvolume int         `json:"dailysoldvolume"`
	Iscase          bool        `json:"iscase"`
	Iskey           bool        `json:"iskey"`
	Itemgroup       string      `json:"itemgroup"`
	Isstar          bool        `json:"isstar"`
}
