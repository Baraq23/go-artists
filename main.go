package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ARTISTS struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConsertDates string   `json:"consertDate"`
	Relations    string   `json:"relations"`
}

type LOCATIONS struct {
	LocationsItems []LocationsContent
}

type LocationsContent struct {
	Id        int    `json:"id"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
}

type DATES struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type DateIndex struct {
	Indexs []DATES `json:"index"`
}

type RELATIONS struct {
	Id             int             `json:"id"`
	RelationStruct []DateLocations `json:"dateLocations"`
}

type DateLocations struct {
	DateL []string
}

func main() {
	// API URL: https://groupietrackers.herokuapp.com/api
	// relations: https://groupietrackers.herokuapp.com/api/relation

	url := "https://groupietrackers.herokuapp.com/api/dates"

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("error fetching from url")
		return
	}
	defer res.Body.Close()

	// var relations RELATIONS
	var artists DateIndex
	decode := json.NewDecoder(res.Body)
	// err = decode.Decode(&relations)
	err = decode.Decode(&artists)
	if err != nil {
		fmt.Println("error decording data")
		return
	}

	// fmt.Println(relations)
	fmt.Println(artists)

}
