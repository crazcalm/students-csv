package students

import (
	"strings"
	"testing"
)

const (
	ExpectedString   = "ChineseName (Pinyin) -- EnglishName"
	ExpectedGetFront = "ChineseName"
	ExpectedGetBack  = "Pinyin"
	ExpectedGetHint  = "EnglishName"
)

func getTestStudent() Student {
	return Student{
		"ChineseName",
		"Pinyin",
		"EnglishName",
		"StudentID",
	}
}

func TestGetHint(t *testing.T) {
	//Student
	s := getTestStudent()

	if strings.Compare(s.GetHint(), ExpectedGetHint) != 0 {
		t.Fail()
	}
}

func TestGetBack(t *testing.T) {
	//Student
	s := getTestStudent()

	if strings.Compare(s.GetBack(), ExpectedGetBack) != 0 {
		t.Fail()
	}
}

func TestGetFront(t *testing.T) {
	//Student
	s := getTestStudent()

	if strings.Compare(s.GetFront(), ExpectedGetFront) != 0 {
		t.Fail()
	}
}

func TestString(t *testing.T) {
	//Student
	s := getTestStudent()

	if strings.Compare(s.String(), ExpectedString) != 0 {
		t.Fail()
	}
}
