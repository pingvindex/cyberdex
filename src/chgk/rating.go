package chgk

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// RespParam Структура ответа от api рейтинга
type RespParam struct {
	ID      int    `json:"idplayer"`
	Release int    `json:"idrelease"`
	Rating  int    `json:"rating"`
	Pos     int    `json:"rating_position"`
	Date    string `json:"date"`
	Tyear   int    `json:"tournaments_in_year"`
	Ttotal  int    `json:"tournament_count_total"`
}

// GetInfo returns rating.chgk info about me
func GetInfo() string {
	c := http.Client{}
	resp, err := c.Get("http://rating.chgk.info/api/players/54035/rating/last")

	if err != nil {
		log.Println("Запрос не удался\n" + err.Error())
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	data := RespParam{}

	err = json.Unmarshal(body, &data)

	result := "id: " + strconv.Itoa(data.ID) + "\nрейтинг: " + strconv.Itoa(data.Rating) +
		"\nместо: " + strconv.Itoa(data.Pos) + "\nтурниров за год: " + strconv.Itoa(data.Tyear)

	return result
}
