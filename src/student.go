package students

import (
	"fmt"
)

//Student struct to hold student information
type Student struct {
	ChineseName string `csv:"chinese_name"`
	Pinyin      string `csv:"pinyin"`
	EnglishName string `csv:"english_name"`
	StudentID   string `csv:"student_id"`
}

//Stringer used for stadnard print function
func (s Student) String() string {
	var result string
	if s.EnglishName == "" {
		result = fmt.Sprintf("%s (%s)", s.ChineseName, s.Pinyin)
	} else {
		result = fmt.Sprintf("%s (%s) -- %s", s.ChineseName, s.Pinyin, s.EnglishName)
	}
	return result
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
