package students

import (
	"bufio"
	"fmt"
	"github.com/artonge/go-csv-tag"
	"github.com/crazcalm/flash-cards/cards"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	//Present -- marker for being present
	Present = "p"
	//Absent --  marker for being absent
	Absent = "a"
	//PresentFileName -- csv file name for present students
	PresentFileName = "present.csv"
	//AbsentFileName -- csv file name for absent students
	AbsentFileName = "absent.csv"
)

//AbsentStudent use to create the absent csv
type AbsentStudent struct {
	ChineseName string `csv:"chinese_name"`
	Pinyin      string `csv:"pinyin"`
	EnglishName string `csv:"english_name"`
	StudentID   string `csv:"student_id"`
	Excuse      string `csv:"excuse"`
}

func createFile(fileName string) *os.File {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func createFileName(suffix string) string {
	date := time.Now()
	return fmt.Sprintf("%d-%d-%d_%s", date.Year(), date.Month(), date.Day(), suffix)
}

//Attendance -- takes attendance and writes a present and absent csv to an io.Writer
//@param ss -- The instance of students being acted on
//@param dir -- The path to the directory where the created csv files will be placed
func Attendance(ss Students, dir string) {
	var present []Student
	var absent []AbsentStudent
	input := bufio.NewScanner(os.Stdin)

	for _, s := range ss.Students {
		fmt.Printf("Student: %v\n\n", s)
		fmt.Printf("(p)resent and (a)bsent: ")

		for input.Scan() {
			userInput := input.Text()

			if strings.Compare(userInput, Present) == 0 {
				fmt.Printf("\n\nYou are here!!!\n\n")
				present = append(present, s)
				break
			} else if strings.Compare(userInput, Absent) == 0 {
				fmt.Println("ABSENT!!!!")
				fmt.Println("Why are they absent?:")

				for input.Scan() {
					excuse := input.Text()
					if strings.Compare(excuse, "") != 0 {
						absent = append(absent, AbsentStudent{s.ChineseName, s.Pinyin, s.EnglishName, s.StudentID, excuse})
						break
					} else {
						fmt.Printf("\n\nYou must add an excuse.\n")
						fmt.Printf("Try again: ")
					}
				}
				//Breaks out of the main loop so that we can go to the next person
				break
			} else {
				fmt.Printf("\n\nThat is not a valid answer...\n")
				fmt.Print("(p)resent and (a)bsent: ")
			}
		}
		flashcards.Clear()
	}

	//Csv Files
	presentFileName := filepath.Join(dir, createFileName(PresentFileName))
	presentFile := createFile(presentFileName)

	absentFileName := filepath.Join(dir, createFileName(AbsentFileName))
	absentFile := createFile(absentFileName)

	//Writing to csv files
	csvtag.Dump(present, presentFile)
	csvtag.Dump(absent, absentFile)

	//Let the user know that the program has finished running
	fmt.Println("The program has finished running")
}
