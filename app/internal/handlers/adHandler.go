package handlers

import (
	"fmt"
	"mp-service/internal/models"
	"mp-service/internal/service"
	"mp-service/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdsHandler struct {
	adService *service.AdService
}

func NewAdsHandler(ads *service.AdService) *AdsHandler {
	return &AdsHandler{adService: ads}
}

func (adh *AdsHandler) CreateAd(c *gin.Context) {
	var ads models.AdsDTO
	err := c.ShouldBindBodyWithJSON(&ads)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(ads.Description, ads.Header)
	valid, err := utils.IsDTOValid(&ads)
	if !valid {
		c.JSON(http.StatusExpectationFailed, gin.H{"error": err.Error()})
		return
	}

	usrEmail, err := utils.GetSubFromToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "error with sub in token"})
		return
	}

	err = adh.adService.CreateAd(ads, usrEmail)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to create advertisement"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ads was succesfully created!"})
}
