package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

// https://docs.wiseoldman.net/names-api/name-type-definitions#object-name-change
type NameChange struct {
	Id       int    `json:"id"`
	PlayerId int    `json:"playerId"`
	OldName  string `json:"oldName"`
	NewName  string `json:"newName"`
	Status   string `json:"status"`
}

// Fetches latest 20 items
var endpoint = "https://api.wiseoldman.net/v2/groups/%s/name-changes"

func GetNameChanges(g string) ([]NameChange, error) {
	url := fmt.Sprintf(endpoint, g)
	var result []NameChange

	response, err := http.Get(url)
	if err != nil {
		return result, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Unexpected status code:", response.StatusCode)
		return result, errors.New("Unexpected status code:" + strconv.Itoa(response.StatusCode))
	}

	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return result, err
	}

	result = removeDuplicates(result)

	if len(result) == 0 {
		fmt.Println("Name changes not found")
		return result, errors.New("Name changes not found")
	}

	return result, nil
}
