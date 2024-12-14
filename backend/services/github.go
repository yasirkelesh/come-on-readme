package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)



// GetSomething, API'den verileri alır ve yalnızca istenen alanları döndürür
func GetContents(owner, repo string) (Content, error) {

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents", owner, repo)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch repo info: %v", err)
	}
	defer resp.Body.Close()

	var responseBody Content

	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		return nil, fmt.Errorf("Failed to fetch repo info: %v", err)
	}
	return responseBody, nil
}

func FetchReadme(owner, repo string) (string, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents", owner, repo)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("Cookie", "_octo=GH1.1.2110837635.1733941375; logged_in=no")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(body), nil
}
