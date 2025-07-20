package service

import "mp-service/internal/repository/ad"

type AdService struct {
	repository *ad.AdRepository
}

func NewAdService(repo *ad.AdRepository) *AdService {
	return &AdService{repository: repo}
}
