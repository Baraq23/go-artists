package main

import (
	"encoding/json"
	"fmt"
	"log"
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
	LocationsItems []LocationsContent `json:"index"`
}

type LocationsContent struct {
	Id        int    `json:"id"`
	Locations []string `json:"locations"`
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
	dates, err := getDates()
	if err != nil {
		return
	}

	artists, err := getAtists()
	if err != nil {
		return
	}

	locations, err := getLocations()
	if err != nil {
		return
	}

	// relations, err := getRelations()
	// if err != nil {
	// 	return
	// }


	fmt.Sprintln(dates)
	fmt.Println("Dates Printed successfully")
	fmt.Sprintln(artists)
	fmt.Println("Artists Printed successfully")

	fmt.Println(locations)
	fmt.Println("Locations Printed successfully")

	// fmt.Sprintln(relations)
	// fmt.Println("Relations Printed successfully")

}

func getRelations() (RELATIONS, error) {
	url := "https://groupietrackers.herokuapp.com/api/relations"
	var relations RELATIONS

	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return relations, err
	}

	defer res.Body.Close()

	decode := json.NewDecoder(res.Body)
	err = decode.Decode(&relations)
	if err != nil {
		log.Println(err)
		return relations, err
	}

	return relations, nil
}



func getLocations() (LOCATIONS, error) {
	url := "https://groupietrackers.herokuapp.com/api/locations"
	var locations LOCATIONS

	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return locations, err
	}

	defer res.Body.Close()

	decode := json.NewDecoder(res.Body)
	err  = decode.Decode(&locations)
	if err != nil {
		log.Println(err)
		return locations, err
	}
	return locations, nil
}

func getAtists() ([]ARTISTS, error) {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	var artists []ARTISTS

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return artists, err
	}
	defer res.Body.Close()

	decode := json.NewDecoder(res.Body)
	err = decode.Decode(&artists)

	if err != nil {
		fmt.Println(err)
		return artists, err
	}

	return artists, nil
}

func getDates() (DateIndex, error) {

	url := "https://groupietrackers.herokuapp.com/api/dates"
	var dates DateIndex

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("error fetching from dates url")
		return dates, err
	}
	defer res.Body.Close()
	
	decode := json.NewDecoder(res.Body)
	err = decode.Decode(&dates)
	if err != nil {
		fmt.Println(err)
		return dates, err
		 
	}
	return dates, nil
}



