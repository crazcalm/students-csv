package students

import (
	"fmt"
	"io"
	"math"
	"math/rand"
	"time"
)

func init() {
	// Seedning the random generator
	rand.Seed(time.Now().UnixNano())
}

//Shuffle Randomizes the ordering of students
func Shuffle(s Students) {
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

//RandomStudent selects a student at random
func RandomStudent(s Students) Student {
	numOfStudents := len(s.Students)
	randNum := rand.Intn(numOfStudents)
	return s.Students[randNum]
}

//PrintGroups used to organize the students into groups
func PrintGroups(s Students, numOfGroups int, w io.Writer, shuffle bool) {
	//Shuffle the students
	if shuffle {
		Shuffle(s)
	}

	numOfPeopleInGroups := math.Ceil(float64(len(s.Students)) / float64(numOfGroups))

	//count is used for the internal loop
	var count int

	for i := 1; i <= numOfGroups; i++ {
		fmt.Fprint(w, fmt.Sprintf("Group %d:\n", i))
		for count < len(s.Students) {
			fmt.Fprint(w, fmt.Sprintf("%d ", count))
			fmt.Fprintln(w, s.Students[count])
			count++
			if count > 0 && math.Mod(float64(count), numOfPeopleInGroups) == 0 {
				break
			}
		}
	}
}
