package models

type Advertisement struct {
	//header, description,url,price
	OwnerID     int    `json:"-"`
	Header      string `json:"ad_header"`
	Description string `json:"ad_description"`
	ImageUrl    string `json:"ad_image"`
	Price       int    `json:"price"`
}

type AdsDTO struct {
	Header      string `json:"header"`
	Description string `json:"description"`
	ImageUrl    string `json:"image"`
	Price       int    `json:"price"`
}
