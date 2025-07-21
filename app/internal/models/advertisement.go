package models

import "time"

type Advertisement struct {
	//header, description,url,price
	OwnerID     int    `json:"-"`
	Header      string `json:"ad_header"`
	Description string `json:"ad_description"`
	ImageUrl    string `json:"ad_image"`
	Price       int    `json:"price"`
	Created_at  time.Time
}

type AdsDTO struct {
	Header      string `json:"header"`
	Description string `json:"description"`
	ImageUrl    string `json:"image"`
	Price       int    `json:"price"`
}

type AdFeed struct {
	ID          int       `json:"id"`
	Header      string    `json:"header"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"image"`
	Price       int       `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	OwnerEmail  string    `json:"owner_email"`
	IsOwner     bool      `json:"is_owner"`
}
