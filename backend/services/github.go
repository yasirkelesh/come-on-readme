package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/come-on-readme/models"
)

// GetSomething, API'den verileri alır ve yalnızca istenen alanları döndürür
func GetGeneralContents(owner, repo string) (*[]models.GeneralContent, error) {

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents", owner, repo)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch repo info: %v", err)
	}
	defer resp.Body.Close()

	var responseBody *[]models.GeneralContent

	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		return nil, fmt.Errorf("failed to fetch repo info: %v", err)
	}
	return responseBody, nil
}

func GetFile(url string) (*models.FileContent, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch file: %v", err)
	}
	defer resp.Body.Close()

	var responseBody models.FileContent

	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		return nil, fmt.Errorf("failed to fetch file: %v", err)
	}
	return &responseBody, nil
}

func GetDir(url string) (*[]models.GeneralContent, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch dir: %v", err)
	}
	defer resp.Body.Close()

	var responseBody *[]models.GeneralContent

	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		return nil, fmt.Errorf("failed to fetch dir: %v", err)
	}
	return responseBody, nil
}
