package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title string
	Year int
	Color bool
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	// ...
}

func main(){
	data, err := json.MarshalIndent(movies, "", "	")
	if err != nil{
		log.Fatalf("JSON marshaling failed :%s \n", err)

	}
	fmt.Printf("%s\n", data)
}
