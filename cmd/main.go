package main

import (
	"fmt"

  "github.com/enescakir/emoji"
	api "main.go/api"
)

func main() {
	api.FetchAndSaveData()
	cityData, err := api.GetCityDataFromMongoDB("Kyiv")
	if err != nil {
		fmt.Printf("[ERROR] error fetching city data from MongoDB: %v", err)
	} else {
		fmt.Println(emoji.Cityscape, "Місто:", cityData.CityName)
		fmt.Println(emoji.House, "Вулиця:", cityData.StationName)
  	fmt.Println(emoji.Compass, "Широта:", cityData.Latitude)
		fmt.Println(emoji.Compass, "Довгота:", cityData.Longitude)
    fmt.Println(emoji.ThreeOClock, "Часовий пояс:", cityData.Timezone)
		for _, pollutant := range cityData.Pollutants {
			fmt.Println(emoji.GemStone, "Показник:", pollutant.Pol)
			fmt.Println("   + Одиниці: ", pollutant.Unit)
			fmt.Println("   + Час:", pollutant.Time)
			fmt.Println("   + Значення:", pollutant.Value)
			fmt.Println("   + У середньому:", pollutant.Averaging)
		}
	}
}
