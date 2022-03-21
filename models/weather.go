package models

import (
	"math/rand"
	"os"
	"strconv"
	"text/tabwriter"
	"time"
)

type Weather struct {
	Water Water
	Wind  Wind
}

type Water struct {
	Status  string
	Message string
}

type Wind struct {
	Status  string
	Message string
}

func GetStatus() Weather {

	var (
		weather Weather
	)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	defer w.Flush()

	waterRand := r.Intn(100)
	windRand := r.Intn(100)

	if waterRand < 5 {
		weather.Water.Message = "aman"
	} else if waterRand > 8 {
		weather.Water.Message = "bahaya"
	} else {
		weather.Water.Message = "siaga"
	}

	if windRand < 6 {
		weather.Wind.Message = "aman"
	} else if windRand > 15 {
		weather.Wind.Message = "bahaya"
	} else {
		weather.Wind.Message = "siaga"
	}

	waterString := strconv.Itoa(waterRand)
	windString := strconv.Itoa(windRand)

	weather.Water.Status = waterString + " m"
	weather.Wind.Status = windString + " m/s"

	return weather
}
