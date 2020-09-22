package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Names map[string][]personID
type People map[personID]*personInfo
type Movies map[movieID]*movieInfo

type personID string
type personInfo struct {
	name   string
	birth  string
	movies []movieID
}
type movieID string
type movieInfo struct {
	title string
	year  string
	stars []personID
}

func main() {
	arg := os.Args[1]
	if arg == "small" {
		fmt.Println("Loading the small test dataset...")
	} else if arg == "large" {
		fmt.Println("Loading the large complete dataset...")
	}

	peopleData, err := readFromFile(arg + "/" + "people.csv")
	if err != nil {
		panic(err)
	}
	movieData, err := readFromFile(arg + "/" + "movies.csv")
	if err != nil {
		panic(err)
	}
	starsData, err := readFromFile(arg + "/" + "stars.csv")
	if err != nil {
		panic(err)
	}

	people := make(People)
	movies := make(Movies)
	names := make(Names)
	for i, p := range peopleData {
		if i == 0 {
			continue
		}
		people[personID(p[0])] = &personInfo{name: p[1], birth: p[2]}
		names[p[1]] = append(names[p[1]], personID(p[0]))
	}

	for i, p := range movieData {
		if i == 0 {
			continue
		}
		movies[movieID(p[0])] = &movieInfo{title: p[1], year: p[2]}
	}

	for i, p := range starsData {
		// p = (personID, movieID)
		if i == 0 {
			continue
		}
		if movies[movieID(p[1])] != nil {
			movies[movieID(p[1])].stars = append(movies[movieID(p[1])].stars, personID(p[0]))
		}
		if people[personID(p[0])] != nil {
			people[personID(p[0])].movies = append(people[personID(p[0])].movies, movieID(p[1]))
		}
	}

	for _, v := range people {
		fmt.Println(v.name, v.birth, v.movies)
	}
	for _, v := range movies {
		fmt.Println(v.title, v.year, v.stars)
	}
	for i, v := range names {
		fmt.Println(i, v)
	}
}

func readFromFile(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := csv.NewReader(file)

	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}
