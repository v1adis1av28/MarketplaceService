package service

import (
	"mp-service/internal/models"
	"mp-service/internal/repository/ad"
)

type AdService struct {
	repository *ad.AdRepository
}

func (a *AdService) CreateAd(ads models.AdsDTO, email string) (*models.Advertisement, error) {
	return a.repository.NewAd(ads, email)
}

func NewAdService(repo *ad.AdRepository) *AdService {
	return &AdService{repository: repo}
}

func (a *AdService) GetAds(sort, order, limit, offset, email string) ([]models.AdFeed, error) {
	return a.repository.FetchAds(sort, order, limit, offset, email)
}
