package main

import (
	"fmt"
	"github.com/artonge/go-csv-tag"
	"math/rand"
	"time"
	"math"
	"github.com/crazcalm/flash-cards/cards"
)

//Students testing
type Students struct {
	Students []Student
}

//GetCards testing
func (s Students) GetCards() []flashcards.Card {
	var cards []flashcards.Card
	for _, item := range s.Students{
		cards = append(cards, flashcards.Card{item.GetFront(), item.GetBack(), item.GetHint()})
	}
	return cards
}

//Student struct to hold student information
type Student struct {
	ChineseName		string `csv:"chinese_name"`
	Pinyin			string `csv:"pinyin"`
	EnglishName		string `csv:"english_name"`
	StudentID		string `csv:"student_id"`
}

//GetFront used as the front of the flashcard
func (s Student) GetFront() string {
	return s.ChineseName
}

//GetBack used as the Back for the flashcard
func (s Student) GetBack() string {
	return s.Pinyin
}

//GetHint used as the hint for the flashcard
func (s Student) GetHint() string {
	return s.EnglishName
}

func init() {
	// Seedning the random generator
	rand.Seed(time.Now().UnixNano())
}

//Shuffle Randomizes the ordering of students
func Shuffle(s []Student){
	numOfStudents := len(s)
	var tempt Student
	var swapIndex int

	for index := range s {
		swapIndex = rand.Intn(numOfStudents)
		tempt = s[index]
		s[index] = s[swapIndex]
		s[swapIndex] = tempt
	}
}

//PrintGroups used to organize the students into groups
func PrintGroups(s []Student, numOfGroups int){
		numOfPeopleInGroups := math.Ceil(float64(len(s)) / float64(numOfGroups))
	
		//count is used for the internal loop
		var count int
	
		for i := 1; i <= numOfGroups; i++ {
			fmt.Printf("Group %d:\n", i)
			for count < len(s) {
				fmt.Print(count, " ")
				fmt.Println(s[count])
				count++
				if count > 0 && math.Mod(float64(count), float64(numOfPeopleInGroups)) == 0 {
					break
				}
			}
		}
}

//RandomStudent selects a student at random
func RandomStudent(s Students) Student {
	numOfStudents := len(s.Students)
	randNum := rand.Intn(numOfStudents)
	return s.Students[randNum]
}

func main(){
	var students Students
	csvtag.Load(csvtag.Config{
		Path: "/home/crazcalm/Documents/PeaceCorps/School/jiqiren161/names.csv",
		Dest: &students.Students,
	})
	//Shuffle(students)

	//PrintGroups(students, 4)
	fmt.Println(RandomStudent(students))

	flashcards.FlashcardApp(students)

}
