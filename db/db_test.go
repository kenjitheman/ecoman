package db

import (
	"context"
	"testing"
)

func TestGetMongoClient(t *testing.T) {
	client, err := GetMongoClient()
	if err != nil {
		t.Errorf("Expected no error when obtaining MongoDB client, got: %v", err)
	}
	defer client.Disconnect(context.Background())

	if client == nil {
		t.Error("Expected a non-nil MongoDB client, got nil")
	}
}

func TestSaveDataToMongoDB(t *testing.T) {
	testCityData := CityData{
		CityName:    "TestCity",
		StationName: "TestStation",
	}

	err := SaveDataToMongoDB(testCityData)
	if err != nil {
		t.Errorf("Expected no error when saving data to MongoDB, got: %v", err)
	}

	retrievedCityData, err := GetCityData("TestCity", "TestStation")
	if err != nil {
		t.Errorf("Expected no error when retrieving saved data from MongoDB, got: %v", err)
	}

	if retrievedCityData.CityName != testCityData.CityName || retrievedCityData.StationName != testCityData.StationName {
		t.Errorf("Saved data does not match retrieved data.")
	}
}
