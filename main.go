package main

import (
	"./app/controllers"
	"./app/models"
	"log"
	"time"
)

func main() {
	s := models.NewSignalEvents()
	df, _ := models.GetAllCandle("BTC_USD", time.Minute, 10)
	c1 := df.Candles[len(df.Candles)-2]
	c2 := df.Candles[len(df.Candles)-1]
	s.Buy("BTC_USD", c1.Time, c1.Close, 1.0, true)
	s.Sell("BTC_USD", c2.Time, c2.Close, 1.0, true)
	controllers.StreamIngestionData()
	log.Println(controllers.StartWebServer())
}
