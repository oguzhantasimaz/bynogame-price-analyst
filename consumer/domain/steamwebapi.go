package domain

import "time"

type SteamWebApi struct {
	Id             string      `json:"id"`
	Markethashname string      `json:"markethashname"`
	Marketname     string      `json:"marketname"`
	Hashid         string      `json:"hashid"`
	ItemType       string      `json:"itemtype"`
	Rarity         interface{} `json:"rarity"`
	Quality        string      `json:"quality"`
	Pricelatest    float64     `json:"pricelatest"`
	Pricemedian    float64     `json:"pricemedian"`
	Priceavg       float64     `json:"priceavg"`
	Slug           string      `json:"slug"`
	Image          string      `json:"image"`
	Wear           string      `json:"wear"`
	CreatedAt      time.Time   `json:"createdAt"`
	IsStattrak     bool        `json:"isstattrak"`
	Isstar         bool        `json:"isstar"`
}
