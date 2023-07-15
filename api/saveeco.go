package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"main.go/db"
)

// Fetches data for all cities, saves it to MongoDB, and returns the saved data
func FetchAndSaveData() ([]db.CityData, error) {
	url := "https://api.saveecobot.com/output.json"

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] error reading response body: %v", err)
	}

	var citiesData []db.CityData
	err = json.Unmarshal(body, &citiesData)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] error parsing JSON: %v", err)
	}

	done := make(chan bool)
	go db.AsyncSaveDataToMongoDB(citiesData, done)

	select {
	case <-done:
		return citiesData, nil
	case <-time.After(3 * time.Hour):
		return nil, fmt.Errorf("[ERROR] data saving took too long")
	}
}

// Searches for a city in the database by its name and returns the data if found
func GetCityDataFromMongoDB(cityName string) (db.CityData, error) {
	return db.GetCityData(cityName)
}

