package main

import (
	"github.com/artonge/go-csv-tag"
	"github.com/crazcalm/flash-cards/cards"
	"flag"
	"github.com/crazcalm/students-csv/src"
	"fmt"
	"log"
)

var csvFile = flag.String("f", "", "file: path to csv file")
var shuffle = flag.Bool("s", false, "Shuffle the cards")
var groups = flag.Int("g", 0, "Groups the students")
var randomStudent = flag.Bool("r", false, "Prints random student")

func main(){
	//Turn on the commandline arguments
	flag.Parse()

	//Basic error handling for flags
	if *csvFile == "" {
		err := fmt.Errorf("No csv file was passed into the program")
		log.Fatalf("%v\n", err)
	}

	var ss students.Students
	err := csvtag.Load(csvtag.Config{
			Path: *csvFile,
			Dest: &ss.Students,
		})
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	cards := flashcards.Cards{ss.GetCards()}

	if *randomStudent == true {
		fmt.Println(students.RandomStudent(ss))
	} else if *groups != 0 {
		students.PrintGroups(ss, *groups)
	} else {
		flashcards.FlashcardApp(cards, *shuffle)
	}
}
