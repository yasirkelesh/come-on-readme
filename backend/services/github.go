package services

import (
	"fmt"
	"io"
	"net/http"
)


func FetchReadme(owner, repo string) (string, error) {
	url := "https://api.github.com/repos/yasirkelesh/INCEPTION/contents"
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
