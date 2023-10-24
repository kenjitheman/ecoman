package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/kenjitheman/ecoman/db"
	"github.com/kenjitheman/ecoman/vars"
)

func FetchAndSaveData() ([]db.CityData, error) {
	resp, err := http.Get(vars.DataUrl)
	if err != nil {
		return nil, fmt.Errorf("Error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response body: %v", err)
	}

	var citiesData []db.CityData
	err = json.Unmarshal(body, &citiesData)
	if err != nil {
		return nil, fmt.Errorf("Error parsing JSON: %v", err)
	}

	done := make(chan bool)
	go db.AsyncSaveDataToMongoDB(citiesData, done)

	select {
	case <-done:
		return citiesData, nil
	case <-time.After(5 * time.Minute):
		return nil, fmt.Errorf("Error: data saving took too long")
	}
}

func GetCityDataFromMongoDB(cityName, stationName string) (db.CityData, error) {
	return db.GetCityData(cityName, stationName)
}
