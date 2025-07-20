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
	//получаем дтошку с данными по объявлению
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
	//После парсинга надо проверить валидность всех данных  в объявлении
	// наличие заголовка и описания(их размерность), изображение проверить на регулярке(обязательно использованеи .jpeg .png)
	// проверка цена просто что не отрицательная

}
