package db

import (
	"context"
	"fmt"
	// "os"
	"time"
  // "log"

	// "github.com/joho/godotenv"
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
	PlatformName string `json:"platformname"`
}

func getMongoClient() (*mongo.Client, error) {
	// err := godotenv.Load("../.env")
	// if err != nil {
 //    fmt.Printf("Error loading .env file: %v", err) 
	// 	log.Fatal("Error loading .env file")
	// }
	// username := os.Getenv("USERNAME")
	// password := os.Getenv("PASSWORD")
 //  
 //  serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	// opts := options.Client().
	// 	ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@cluster0.7havayh.mongodb.net/?retryWrites=true&w=majority", username, password)).
	// 	SetServerAPIOptions(serverAPI)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().
		ApplyURI("mongodb+srv://username:password@cluster0.7havayh.mongodb.net/?retryWrites=true&w=majority").
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

	collection := client.Database("weatherdata").Collection("dataman")

	_, err = collection.InsertOne(context.Background(), data)
	if err != nil {
		return fmt.Errorf("[ERROR] failed to insert data into MongoDB: %v", err)
	}

	fmt.Println("[SUCCESS] data saved to MongoDB!")
	return nil
}

func GetCityData(cityName string) (CityData, error) {
	client, err := getMongoClient()
	if err != nil {
		return CityData{}, err
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("weatherdata").Collection("dataman")

	filter := bson.M{"cityname": cityName}
	var cityData CityData
	err = collection.FindOne(context.Background(), filter).Decode(&cityData)
	if err != nil {
		return CityData{}, fmt.Errorf("[ERROR] failed to fetch city data from MongoDB: %v", err)
	}

	return cityData, nil
}

func FetchDataFromMongoDB() ([]CityData, error) {
	client, err := getMongoClient()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("weatherdata").Collection("dataman")

	cursor, err := collection.Find(context.Background(), nil)
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

func AsyncSaveDataToMongoDB(data []CityData, done chan<- bool) {
	client, err := getMongoClient()
	if err != nil {
		fmt.Printf("[ERROR] failed to connect to MongoDB: %v\n", err)
		done <- true
		return
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("weatherdata").Collection("dataman")

	// Delete all existing data before saving new data
	_, err = collection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		fmt.Printf("[ERROR] failed to delete existing data from MongoDB: %v\n", err)
		done <- true
		return
	}

	// Insert new data
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
