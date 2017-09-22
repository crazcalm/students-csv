package main

import (
	"github.com/artonge/go-csv-tag"
	"github.com/crazcalm/flash-cards/cards"
	"flag"
	"github.com/crazcalm/students-csv/src"
)

var csvFile = flag.String("f", "", "file: path to csv file")
var shuffle = flag.Bool("s", false, "Shuffle the cards")
var groups = flag.Int("g", 0, "Groups the students")

func main(){
	//Turn on the commandline arguments
	flag.Parse()

	var ss students.Students
	csvtag.Load(csvtag.Config{
		Path: *csvFile,
		Dest: &ss.Students,
	})

	//PrintGroups(students.Students, 4)
	//fmt.Println(RandomStudent(&students))

	cards := flashcards.Cards{ss.GetCards()}

	if *groups == 0{
		flashcards.FlashcardApp(cards, *shuffle) //true == shuffle cards
	} else {
		students.PrintGroups(ss, *groups)
	}
}
