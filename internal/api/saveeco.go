package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CityData struct {
	ID          string `json:"id"`
	CityName    string `json:"cityName"`
	StationName string `json:"stationName"`
	LocalName   string `json:"localName"`
	Timezone    string `json:"timezone"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	Pollutants  []struct {
		Pol       string  `json:"pol"`
		Unit      string  `json:"unit"`
		Time      string  `json:"time"`
		Value     float64 `json:"value"`
		Averaging string  `json:"averaging"`
	} `json:"pollutants"`
	PlatformName string `json:"platformName"`
}

// Specify the city name you want to retrieve data for
func GetData(desiredCity string) {
	url := "https://api.saveecobot.com/output.json"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error making HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading response body:", err)
		return
	}

	var citiesData []CityData
	err = json.Unmarshal(body, &citiesData)
	if err != nil {
		fmt.Println("error parsing JSON:", err)
		return
	}

	// Find the data for the desired city
	for _, cityData := range citiesData {
		if cityData.CityName == desiredCity {
			fmt.Println("Місто:", cityData.CityName)
			fmt.Println("Район:", cityData.StationName)

			for _, pollutant := range cityData.Pollutants {
				fmt.Println("Забруднювач:", pollutant.Pol)
				fmt.Println("Одиниці:", pollutant.Unit)
				fmt.Println("Час:", pollutant.Time)
				fmt.Println("Значення:", pollutant.Value)
				fmt.Println("У середньому:", pollutant.Averaging)
			}

			// Add your logic to process the data for the desired city
			// ...

			// Break the loop since we have found the desired city
			break
		}
	}
}
