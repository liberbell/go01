package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
)

type Tour struct {
	Name, Price, Tourid string
}

func main() {
	url := "http://services.explorecalifornia.org/json/tours.php"
	content := contentFromServer(url)

	// fmt.Println(content)

	tours := toursFromJson(content)
	// fmt.Println(tours)

	for _, tour := range tours {
		price, _, _ := big.ParseFloat(tour.Price, 10, 2, big.ToZero)
		tourid, _, _ := big.ParseFloat(tour.Tourid, 10, 0, big.ToZero)
		fmt.Printf("Tour id : %v Tour title:%v Price:($%.2f)\n", tourid, tour.Name, price)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func contentFromServer(url string) string {
	resp, err := http.Get(url)
	checkError(err)

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	return string(bytes)
}

func toursFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20)

	decorder := json.NewDecoder(strings.NewReader(content))
	_, err := decorder.Token()
	checkError(err)

	var tour Tour
	for decorder.More() {
		err := decorder.Decode(&tour)
		checkError(err)

		tours = append(tours, tour)
	}

	return tours
}
