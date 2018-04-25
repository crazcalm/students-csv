package main

import (
	"flag"
	"fmt"
	"github.com/artonge/go-csv-tag"
	"github.com/crazcalm/flash-cards/cards"
	"github.com/crazcalm/students-csv/src"
	"log"
	"os"
	"path/filepath"
)

var csvFile = flag.String("f", "", "file: path to csv file")
var shuffle = flag.Bool("s", false, "Shuffle the cards")
var groups = flag.Int("g", 0, "Groups the students")
var randomStudent = flag.Bool("r", false, "Prints random student")
var attendance = flag.Bool("a", false, "Take attendace")
var output = flag.String("o", ".", "The directory where the files will be created")

func main() {
	//Turn on the commandline arguments
	flag.Parse()

	//Basic error handling for flags
	if *csvFile == "" {
		err := fmt.Errorf("No csv file was passed into the program")
		log.Fatalf("%v\n", err)
	}

	//Check to see of that file exists...
	info, err := os.Stat(*csvFile)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	//Make sure that file is not a dir
	if info.IsDir() {
		log.Fatal("This is not a file, but a directory")
	}

	var ss students.Students
	err = csvtag.Load(csvtag.Config{
		Path: *csvFile,
		Dest: &ss.Students,
	})

	//Make sure csvtag did not throw an error
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	cards := flashcards.Cards{ss.GetCards()}

	if *randomStudent {
		fmt.Println(students.RandomStudent(ss))
	} else if *groups != 0 {

		students.PrintGroups(ss, *groups, os.Stdout, *shuffle)
	} else if *attendance {

		//Check to see if the file exists
		info, err = os.Stat(filepath.Clean(*output))
		if err != nil {
			log.Fatalf("%v\n", err)
		}

		//Ensure that output is a directory
		if !info.IsDir() {
			log.Fatal("Output needs to be a directory")
		}

		students.Attendance(ss, *output)
	} else {
		flashcards.FlashcardApp(cards, *shuffle)
	}
}
