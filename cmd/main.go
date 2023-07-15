package main

import (
  api "main.go/api"
  "fmt"
)

func main()  {
  api.FetchAndSaveData()
  fmt.Println(api.GetCityDataFromMongoDB("Kyiv"))
}
