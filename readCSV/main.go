package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type book struct {
	id     string
	title  string
	author string
	price  string
}

func createbookList(data [][]string) []book {
	var mybooks []book
	for i, line := range data {
		if i > 0 { // omit header line
			var rec book
			for j, field := range line {
				if j == 0 {
					rec.id = field
				} else if j == 1 {
					rec.title = field
				} else if j == 2 {
					rec.author = field
				} else {
					rec.price = field
				}
			}
			mybooks = append(mybooks, rec)
		}
	}
	return mybooks
}

func printbookList(mybooks []book) {

	fmt.Printf("ID\tTITLE\t\t\t\tAUTHOR\t\tPRICE\n")

	for _, rec := range mybooks {
		fmt.Printf("%s\t%s\t%s\t%s\n", rec.id, rec.title, rec.author, rec.price)
	}
}

func main() {

	f, err := os.Open("booklist.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	mybooks := createbookList(data)
	printbookList(mybooks)
	
}
