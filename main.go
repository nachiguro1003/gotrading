package main

import (
	// "./app/controllers"
	"./app/models"
	"./bitflyer"
	"./config"
	"./utils"
	"fmt"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	fmt.Println(models.DbConnection)
	tickerChan := make(chan bitflyer.Ticker)
	apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChan)
}
