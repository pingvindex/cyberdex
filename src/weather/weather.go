package weather

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// CoordType is structure of coordinates
type CoordType struct {
	Lat int `json:"lat"`
	Lgt int `json:"lgt"`
}

// Info ...
type Info struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// MainStruct contains measurement's settings
type MainStruct struct {
	Temp     float64 `json:"temp"`
	Pressure int     `json:"pressure"`
	Humidity int     `json:"humidity"`
	TempMin  float64 `json:"temp_min"`
	TempMax  float64 `json:"temp_max"`
}

// WindStruct ...
type WindStruct struct {
	Speed float64 `json:"speed"`
	Def   int     `json:"deg"`
}

// CloudStruct ...
type CloudStruct struct {
	All int `json:"all"`
}

// SysStruct ...
type SysStruct struct {
	Type    int     `json:"type"`
	ID      int     `json:"id"`
	Message float64 `json:"message"`
	Country string  `json:"country"`
	Sunrise int     `json:"sunrise"`
	Sunset  int     `json:"sunset"`
}

// RespParam contains information about weather
type RespParam struct {
	Coord      CoordType   `json:"coord"`
	Weather    Info        `json:"weather"`
	Base       string      `json:"base"`
	Main       MainStruct  `json:"main"`
	Visibility int         `json:"visibility"`
	Wind       WindStruct  `json:"wind"`
	Clouds     CloudStruct `json:"clouds"`
	DT         int         `json:"dt"`
	Sys        SysStruct   `json:"sys"`
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	COD        int         `json:"cod"`
}

// GetWeather returns info about current weather in Moscow
func GetWeather() string {
	c := http.Client{}
	resp, err := c.Get("http://api.openweathermap.org/data/2.5/weather?q=Moscow&units=metric&lang=ru&APPID=7a3937709a28279ddeca2d281dec984f")
	if err != nil {
		log.Println("Данные о погоде недоступны")
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	weather := RespParam{}

	err = json.Unmarshal(body, &weather)
	if err != nil {
		return "Данные о погоде недоступны"
	}
	result := "" + weather.Weather.Description + ", температура"
	if weather.Main.TempMin >= 0 {
		result += " +" + strconv.FormatFloat(weather.Main.TempMin, 'E', -1, 64)
	} else {
		result += " " + strconv.FormatFloat(weather.Main.TempMin, 'E', -1, 64)
	}
	if weather.Main.TempMax >= 0 {
		result += " +" + strconv.FormatFloat(weather.Main.TempMax, 'E', -1, 64)
	} else {
		result += " " + strconv.FormatFloat(weather.Main.TempMax, 'E', -1, 64)
	}
	return result
}