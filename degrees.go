package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

func NewData() *Data {
	return &Data{}
}

type Data struct {
	Names  Names
	People People
	Movies Movies
}

// Names is a map with keys of actors' names that points to a slice of personIDs
type Names map[string][]personID

// People is a map with keys of personIDs pointing to a personInfo struct
type People map[personID]*personInfo

// Movies is a map with keys of movieIDs pointing to a movieInfo struct
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
	data := Data{
		Names:  names,
		People: people,
		Movies: movies,
	}
	for i, p := range peopleData {
		if i == 0 {
			continue
		}
		data.People[personID(p[0])] = &personInfo{name: p[1], birth: p[2]}
		data.Names[p[1]] = append(data.Names[p[1]], personID(p[0]))
	}

	for i, p := range movieData {
		if i == 0 {
			continue
		}
		data.Movies[movieID(p[0])] = &movieInfo{title: p[1], year: p[2]}
	}

	for i, p := range starsData {
		// p = (personID, movieID)
		if i == 0 {
			continue
		}
		if data.Movies[movieID(p[1])] != nil {
			data.Movies[movieID(p[1])].stars = append(data.Movies[movieID(p[1])].stars, personID(p[0]))
		}
		if data.People[personID(p[0])] != nil {
			data.People[personID(p[0])].movies = append(data.People[personID(p[0])].movies, movieID(p[1]))
		}
	}

	// for _, v := range data.People {
	// 	fmt.Println(v.name, v.birth, v.movies)
	// }
	// for _, v := range data.Movies {
	// 	fmt.Println(v.title, v.year, v.stars)
	// }
	// for i, v := range data.Names {
	// 	fmt.Println(i, v)
	// }

	fmt.Println("Done loading data.")
	var actor1 string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Name of the first actor: ")
	if scanner.Scan() {
		actor1 = scanner.Text()
	}
	// fmt.Print("Name of the second actor: ")
	// if scanner.Scan() {
	// 	actor2 = scanner.Text()
	// }
	id, err := data.personIDfromName(actor1)
	if err != nil {
		fmt.Println("Error encountered: ", err)
		return
	}
	fmt.Println(id)
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

func (d Data) personIDfromName(name string) (personID, error) {
	var id personID
	if d.Names[name] == nil {
		return "", errors.New("Name not found: " + name)
	} else if len(d.Names[name]) > 1 {
		fmt.Printf("There are %d actors named %s.\n", len(d.Names[name]), name)
		for i, v := range d.Names[name] {
			fmt.Printf("%d) ID: %s Birth Year: %s Movies: ", i+1, v, d.People[v].birth)
			mID := d.People[v].movies
			for _, m := range mID {
				fmt.Printf("'%s' ", d.Movies[m].title)
			}
			fmt.Println("")

		}
		count := 0
		for count > len(d.Names[name]) || count <= 0 {
			fmt.Printf("Enter the number: ")
			fmt.Scanf("%d\n", &count)
		}
		id = d.Names[name][count-1]
	} else {
		id = d.Names[name][0]
	}
	return id, nil
}
