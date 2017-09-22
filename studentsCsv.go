package main

import (
	"github.com/artonge/go-csv-tag"
	"github.com/crazcalm/flash-cards/cards"
	"flag"
	"github.com/crazcalm/students-csv/src"
	"fmt"
)

var csvFile = flag.String("f", "", "file: path to csv file")
var shuffle = flag.Bool("s", false, "Shuffle the cards")
var groups = flag.Int("g", 0, "Groups the students")
var randomStudent = flag.Bool("r", false, "Prints random student")

func main(){
	//Turn on the commandline arguments
	flag.Parse()

	var ss students.Students
	csvtag.Load(csvtag.Config{
		Path: *csvFile,
		Dest: &ss.Students,
	})

	//fmt.Println(RandomStudent(&students))

	cards := flashcards.Cards{ss.GetCards()}

	if *randomStudent == true {
		fmt.Println(students.RandomStudent(ss))
	} else if *groups != 0 {
		students.PrintGroups(ss, *groups)
	} else {
		flashcards.FlashcardApp(cards, *shuffle)
	}
}
