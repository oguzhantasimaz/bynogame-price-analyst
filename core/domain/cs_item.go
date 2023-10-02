package domain

import (
	"context"
)

const CollectionCsItem = "cs_item"

type CsItemUseCase interface {
	GetCsItems(ctx context.Context, limit int) ([]*CsItem, error)
	PostCsItem(ctx context.Context, items *CsItem) error
}

type CsItem struct {
	SellerMarketName string          `json:"sellerMarketName" bson:"sellerMarketName"`
	Description      string          `json:"description" bson:"description"`
	ViewCounter      int             `json:"viewCounter" bson:"viewCounter"`
	SellerId         string          `json:"sellerId" bson:"sellerId"`
	Price            float64         `json:"price" bson:"price"`
	Image            string          `json:"image" bson:"image"`
	NameSlug         string          `json:"nameSlug" bson:"nameSlug"`
	SellerLogo       string          `json:"sellerLogo" bson:"sellerLogo"`
	Bngid            int             `json:"bngid" bson:"bngid"`
	CommentCount     int             `json:"commentCount" bson:"commentCount"`
	TypeInfoSteam    CsItemInfoSteam `json:"typeInfoSteam" bson:"typeInfoSteam"`
	GameCode         string          `json:"gameCode" bson:"gameCode"`
	SellerOnline     bool            `json:"sellerOnline" bson:"sellerOnline"`
	Status           string          `json:"status" bson:"status"`
	CreatedAt        int64           `json:"createdAt" bson:"createdAt"`
	DateTimeSold     int64           `json:"dateTimeSold" bson:"dateTimeSold"`
	SteamStats       SteamResponse   `json:"steamStats" bson:"steamStats"`
}

type CsItemInfoSteam struct {
	App          int     `json:"app" bson:"app"`
	MarketPrice  float64 `json:"marketPrice" bson:"marketPrice"`
	StatTrak     bool    `json:"statTrak" bson:"statTrak"`
	Exterior     string  `json:"exterior" bson:"exterior"`
	HashShort    string  `json:"hashShort" bson:"hashShort"`
	ItemSteamId  string  `json:"ItemSteamId" bson:"itemSteamId"`
	Stickers     string  `json:"stickers" bson:"stickers"`
	Type         string  `json:"type" bson:"type"`
	HashBare     string  `json:"hashBare" bson:"hashBare"`
	Hash         string  `json:"hash" bson:"hash"`
	StickersText string  `json:"stickersText" bson:"stickersText"`
}
