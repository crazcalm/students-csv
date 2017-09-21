package students

import (
	"math/rand"
	"time"
)

func init() {
	// Seedning the random generator
	rand.Seed(time.Now().UnixNano())
}

//Shuffle Randomizes the ordering of students
func Shuffle(s Students){
	numOfStudents := len(s.Students)
	var tempt Student
	var swapIndex int

	for index := range s.Students {
		swapIndex = rand.Intn(numOfStudents)
		tempt = s.Students[index]
		s.Students[index] = s.Students[swapIndex]
		s.Students[swapIndex] = tempt
	}
}