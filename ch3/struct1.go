package main

import (
	"fmt"
	"time"
)

// Bootcamp Bootcamp type
type Bootcamp struct {
	Lat  float64
	Lon  float64
	Date time.Time
}

func main() {
	event := Bootcamp{
		Lat: 34.1234,
		Lon: -118.1234,
	}
	event.Date = time.Now()
	fmt.Printf("Event on %s at %f, %f\n", event.Date, event.Lat, event.Lon)
}
