package bot

import (
	"fmt"
	"github.com/enescakir/emoji"
	"github.com/kenjitheman/ecoman/api"
)

func Datafetch(cityname, stationName string) string {
	cityData, err := api.GetCityDataFromMongoDB(cityname, stationName)
	if err != nil {
		fmt.Printf("[ERROR] error fetching city data from MongoDB: %v", err)
		return "error fetching city data! \ncity is incorrect!"
	} else {
		result := fmt.Sprintf("%v City: %s\n", emoji.Cityscape, cityData.CityName)
		result += fmt.Sprintf("%v Station: %s\n", emoji.House, cityData.StationName)
		result += fmt.Sprintf("%v Latitude: %s\n%v Longitude: %s\n%v Timezone: %s\n",
			emoji.Compass, cityData.Latitude,
			emoji.Compass, cityData.Longitude,
			emoji.ThreeOClock, cityData.Timezone)
		for _, pollutant := range cityData.Pollutants {
			pollutantInfo := fmt.Sprintf("%v Pollutant: %s\n   + Units: %s\n   + Time: %v\n   + Value: %v\n   + Average: %v\n",
				emoji.GemStone, pollutant.Pol, pollutant.Unit, pollutant.Time, pollutant.Value, pollutant.Averaging)
			result += pollutantInfo
		}
		return result
	}
}
