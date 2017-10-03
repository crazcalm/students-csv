package students

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"github.com/crazcalm/go-csv-tag"
	"github.com/crazcalm/flash-cards/cards"
)

const (
	//PRESENT -- marker for being present
	PRESENT = "p"
	//ABSENT --  marker for being absent
	ABSENT = "a"
)

//AbsentStudent use to create the absent csv
type AbsentStudent struct {
	ChineseName string `csv:"chinese_name"`
	Pinyin      string `csv:"pinyin"`
	EnglishName string `csv:"english_name"`
	StudentID   string `csv:"student_id"`
	Excuse 		string `csv:"excuse"`
}

//Attendance -- takes attendance and writes a present and absent csv to an io.Writer
func Attendance(ss Students){
	var present []Student
	var absent []AbsentStudent
	input := bufio.NewScanner(os.Stdin)

	for _, s := range ss.Students {
		fmt.Printf("Student: %v\n\n", s)
		fmt.Printf("(p)resent and (a)bsent: " )

		for input.Scan(){
			userInput := input.Text()

			if strings.Compare(userInput, PRESENT) == 0 {
				fmt.Printf("\n\nYou are here!!!\n\n")
				present = append(present, s)
				break
			} else if strings.Compare(userInput, ABSENT) == 0 {
				fmt.Println("ABSENT!!!!")
				fmt.Println("Why are they absent?:")

				for input.Scan() {
					excuse := input.Text()
					if strings.Compare(excuse, "") != 0 {
						absent = append(absent,AbsentStudent{s.ChineseName, s.Pinyin, s.EnglishName, s.StudentID, excuse})
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

	csvtag.Dump(present, os.Stdout)
	fmt.Println(absent)
	csvtag.Dump(absent, os.Stdout)
}