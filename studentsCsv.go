package main

import (
	"github.com/artonge/go-csv-tag"
	"github.com/crazcalm/flash-cards/cards"
	"flag"
	"github.com/crazcalm/students-csv/src"
)

var csvFile = flag.String("f", "", "file: path to csv file")
var shuffle = flag.Bool("s", false, "Shuffle the cards")

func main(){
	//Turn on the commandline arguments
	flag.Parse()

	var students students.Students
	csvtag.Load(csvtag.Config{
		Path: *csvFile,
		Dest: &students.Students,
	})

	//PrintGroups(students.Students, 4)
	//fmt.Println(RandomStudent(&students))

	cards := flashcards.Cards{students.GetCards()}

	flashcards.FlashcardApp(cards, *shuffle) //true == shuffle cards

}
