package utils

import (
	"fmt"
	"mp-service/internal/models"
	"net/url"
	"regexp"
	"unicode/utf8"
)

var imageRegexp = regexp.MustCompile(`^(https?:\/\/.*\.(?:png|jpe?g))$`)

func IsDTOValid(dto *models.AdsDTO) (bool, error) {
	fmt.Println(dto.Description, dto.Header)
	err := IsHeaderValid(dto.Header)
	if err != nil {
		return false, err
	}

	err = IsDescriptionValid(dto.Description)
	if err != nil {
		return false, err
	}

	err = IsPriceValid(dto.Price)
	if err != nil {
		return false, err
	}

	err = IsImageURLValid(dto.ImageUrl)
	if err != nil {
		return false, err
	}

	return true, nil
}

func IsHeaderValid(s string) error {
	if len(s) == 0 {
		return fmt.Errorf("header can`t be empty!")
	}
	if utf8.RuneCountInString(s) > 128 {
		return fmt.Errorf("Header size should be least than 128 characters")
	}
	return nil
}

func IsDescriptionValid(s string) error {
	if len(s) == 0 {
		return fmt.Errorf("Description can`t be empty!")
	}
	if len(s) > 512 {
		return fmt.Errorf("Header size should be least than 512 characters")
	}
	return nil
}

func IsPriceValid(p int) error {
	if p < 0 {
		return fmt.Errorf("Price can`t be negative or equal 0")
	}
	return nil
}

func IsImageURLValid(u string) error {
	if !imageRegexp.MatchString(u) {
		return fmt.Errorf("URL does not match the image format (expected .png, .jpg, or .jpeg)")
	}

	parsedURL, err := url.ParseRequestURI(u)
	if err != nil {
		return fmt.Errorf("invalid URL format")
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return fmt.Errorf("URL must use HTTP or HTTPS")
	}

	return nil
}
