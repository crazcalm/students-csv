package students

import (
	"testing"
	"github.com/artonge/go-csv-tag"
	"strings"
	"path/filepath"
)


func getTestData() Students{
	var ss Students
		csvtag.Load(csvtag.Config{
			Path: filepath.Join("test_data", "names.csv"),
			Dest: &ss.Students,
		})
	return ss
}

func compareStudentsSlice(s1, s2 []Student) bool {
	result := true

	for index, item := range s1 {
		if strings.Compare(item.StudentID, s2[index].StudentID) != 0 {
			result = false
			break
		}
	}
	return result
}

func TestShuffle(t *testing.T){
	//Not Shuffled list
	s1 := getTestData()

	//Shuffled list
	s2 := getTestData() 
	Shuffle(s2)

	if compareStudentsSlice(s1.Students, s2.Students) == true {
		t.Fail()
	}
}

func TestRandomStudent(t *testing.T) {
	//Students
	ss := getTestData()

	// Using 3 students to reduce the chance of false positives
	s1 := RandomStudent(ss)
	s2 := RandomStudent(ss)
	s3 := RandomStudent(ss)

	if strings.Compare(s1.StudentID, s2.StudentID) == 0 && strings.Compare(s2.StudentID, s3.StudentID) == 0 {
		t.Fail()
	}
	
}