package mocking_test

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetRepos(username string) ([]map[string]interface{}, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/repos?sort=created&direction=desc", username)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	m := make([]map[string]interface{}, 0)
	err = json.NewDecoder(resp.Body).Decode(&m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
