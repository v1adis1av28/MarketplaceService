package service

import (
	"mp-service/internal/models"
	"mp-service/internal/repository/ad"
)

type AdService struct {
	repository *ad.AdRepository
}

func (a *AdService) CreateAd(ads models.AdsDTO, email string) error {
	return a.repository.NewAd(ads, email)
}

func NewAdService(repo *ad.AdRepository) *AdService {
	return &AdService{repository: repo}
}
