package domain

type CsItem struct {
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
	SteamStats       SteamResponse   `json:"steamStats"`
}
