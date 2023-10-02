package domain

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
	Float        float64 `json:"float"`
}

type CustomTime struct {
	UnixTime int64
}
