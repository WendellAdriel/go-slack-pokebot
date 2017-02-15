package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

const API_PATH string = "https://pokeapi.co/api/v2/"

func GetPokemonInfo(searchValue string) string {
	url := fmt.Sprintf(API_PATH+"%s/%s", "pokemon", searchValue)

	response, err := makeRequest(url)
	if err != nil {
		return err.Error()
	}

	finalText := buildResponseText(response)
	return finalText
}

func makeRequest(url string) (map[string]interface{}, error) {
	netClient := &http.Client{
		Timeout: time.Second * 10,
	}

	response, err := netClient.Get(url)
	checkError(err)

	defer response.Body.Close()

	var data map[string]interface{}
	parseError := json.NewDecoder(response.Body).Decode(&data)
	checkError(parseError)

	if response.StatusCode == 200 {
		return data, nil
	}
	return data, errors.New("Pokemon not found")
}

func buildResponseText(data map[string]interface{}) string {
	floatId := data["id"].(float64)
	name := data["name"].(string)
	weight := data["weight"].(float64)
	id := int(floatId)

	var responseBuffer bytes.Buffer

	responseBuffer.WriteString(fmt.Sprintf("Name: %s\n", strings.ToUpper(name)))
	responseBuffer.WriteString(fmt.Sprintf("Number: #%d\n", id))
	responseBuffer.WriteString(fmt.Sprintf("Weight: %.2f\n", weight))

	return responseBuffer.String()
}

func checkError(err error) {
	if err != nil {
		log.Fatal("Error: ", err)
	}
}
