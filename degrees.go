package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Names map[string][]personID
type People map[personID]personInfo
type Movies map[movieID]movieInfo

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

	records, err := readFromFile(arg + "/" + "people.csv")
	if err != nil {
		panic(err)
	}

	people := make(People)
	for i, p := range records {
		if i == 0 {
			continue
		}
		people[personID(p[0])] = personInfo{name: p[1], birth: p[2]}
	}

	fmt.Println(people)
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
