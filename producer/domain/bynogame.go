package domain

type ByNoGameCsResponse struct {
	Status      bool   `json:"status"`
	Code        int    `json:"code"`
	Description string `json:"description"`
	Response    struct {
		Last int64            `json:"last"`
		Data []ByNoGameCsItem `json:"data"`
	} `json:"response"`
	ValidationErrors []interface{} `json:"validationErrors"`
}
