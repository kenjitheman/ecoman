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
			fmt.Println("City Name:", cityData.CityName)
			fmt.Println("Station Name:", cityData.StationName)
			fmt.Println("Local Name:", cityData.LocalName)

			for _, pollutant := range cityData.Pollutants {
				fmt.Println("Pollutant:", pollutant.Pol)
				fmt.Println("Unit:", pollutant.Unit)
				fmt.Println("Time:", pollutant.Time)
				fmt.Println("Value:", pollutant.Value)
				fmt.Println("Averaging:", pollutant.Averaging)
			}

			// Add your logic to process the data for the desired city
			// ...

			// Break the loop since we have found the desired city
			break
		}
	}
}
