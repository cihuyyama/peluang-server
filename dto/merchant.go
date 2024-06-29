package dto

type MerchantRequest struct {
	Name          string `json:"name"`
	Desc          string `json:"desc"`
	Category      string `json:"category"`
	BusinessModel string `json:"business_model"`
	ImgUrl        string `json:"img_url"`
}

type MerchantImageRequest struct {
	ImgURL string `json:"img_url"`
	Key    string `json:"key"`
}
