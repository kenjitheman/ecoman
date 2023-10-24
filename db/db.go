package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kenjitheman/ecoman/vars"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CityData struct {
	ID          string `json:"id"`
	CityName    string `json:"cityname"`
	StationName string `json:"stationname"`
	LocalName   string `json:"localname"`
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
	Stations     []string `json:"stations"`
	PlatformName string   `json:"platformname"`
}

func getMongoClient() (*mongo.Client, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Printf("Error loading .env file: %v", err)
		log.Panic(err)
	}
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().
		ApplyURI(os.Getenv("MONGO_URI")).
		SetServerAPIOptions(serverAPI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to ping MongoDB: %v", err)
	}

	fmt.Println("[SUCCESS] connected to MongoDB!")
	return client, nil
}

func SaveDataToMongoDB(data CityData) error {
	client, err := getMongoClient()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(vars.DbName).Collection(vars.CollectionName)

	_, err = collection.InsertOne(context.Background(), data)
	if err != nil {
		return fmt.Errorf("[ERROR] failed to insert data into MongoDB: %v", err)
	}

	fmt.Println("[SUCCESS] data saved to MongoDB!")
	return nil
}

func GetCityData(cityName, stationName string) (CityData, error) {
	client, err := getMongoClient()
	if err != nil {
		return CityData{}, err
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(vars.DbName).Collection(vars.CollectionName)

	filter := bson.M{
		"cityname":    cityName,
		"stationname": stationName,
	}
	var cityData CityData
	err = collection.FindOne(context.Background(), filter).Decode(&cityData)
	if err != nil {
		return CityData{}, fmt.Errorf("[ERROR] failed to fetch city data from MongoDB: %v", err)
	}

	return cityData, nil
}

func FetchDataFromMongoDB(cityName string) ([]CityData, error) {
	client, err := getMongoClient()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(vars.DbName).Collection(vars.CollectionName)

	filter := bson.M{
		"cityname": cityName,
	}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var citiesData []CityData
	err = cursor.All(context.Background(), &citiesData)
	if err != nil {
		return nil, err
	}

	return citiesData, nil
}

func FetchAllCityNamesFromMongoDB() ([]string, error) {
	client, err := getMongoClient()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(vars.DbName).Collection(vars.CollectionName)

	filter := bson.M{}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	citySet := make(map[string]bool)
	for cursor.Next(context.Background()) {
		var cityData CityData
		if err := cursor.Decode(&cityData); err != nil {
			return nil, err
		}
		citySet[cityData.CityName] = true
	}

	uniqueCityNames := make([]string, 0, len(citySet))
	for cityName := range citySet {
		uniqueCityNames = append(uniqueCityNames, cityName)
	}

	return uniqueCityNames, nil
}

func AsyncSaveDataToMongoDB(data []CityData, done chan<- bool) {
	client, err := getMongoClient()
	if err != nil {
		fmt.Printf("[ERROR] failed to connect to MongoDB: %v\n", err)
		done <- true
		return
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(vars.DbName).Collection(vars.CollectionName)

	_, err = collection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		fmt.Printf("[ERROR] failed to delete existing data from MongoDB: %v\n", err)
		done <- true
		return
	}

	var documents []interface{}
	for _, cityData := range data {
		documents = append(documents, cityData)
	}

	_, err = collection.InsertMany(context.Background(), documents)
	if err != nil {
		fmt.Printf("[ERROR] failed to insert data into MongoDB: %v\n", err)
	}

	fmt.Println("[SUCCESS] data saved to MongoDB!")
	done <- true
}
