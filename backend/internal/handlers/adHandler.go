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

// CreateAdvertisement godoc
// @Summary Создание объявления
// @Security ApiKeyAuth
// @Tags advertisements
// @Accept json
// @Produce json
// @Param input body models.Advertisement true "Новое объявление"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/advertisement [post]
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

	advertisement, err := adh.adService.CreateAd(ads, usrEmail)
	fmt.Println("handler ads")
	fmt.Println(advertisement)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to create advertisement"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ads was succesfully created!",
		"advertisement": advertisement})
}

// GetAdvertisements godoc
// @Summary Получить список объявлений
// @Tags advertisements
// @Produce json
// @Param limit query int false "Лимит"
// @Param offset query int false "Смещение"
// @Param order query string false "По убыванию или возрастанию (desc, asc)"
// @Param sort query string false "По какому полю будет сортировка (price, created_at)"
// @Success 200 {array} models.Advertisement
// @Router /api/advertisements [get]
func (adh *AdsHandler) GetAds(c *gin.Context) {
	sortField := c.DefaultQuery("sort", "created_at")
	order := c.DefaultQuery("order", "desc")
	limit := c.DefaultQuery("limit", "5")
	offset := c.DefaultQuery("offset", "0")

	userEmail, _ := utils.GetSubFromToken(c.GetHeader("Authorization"))

	ads, err := adh.adService.GetAds(sortField, order, limit, offset, userEmail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch ads feed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ads": ads})
}
