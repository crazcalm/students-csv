package students

import (
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestcreateFileName(t *testing.T) {
	//Obtaining todays date
	today := time.Now()

	//suffic used for file
	suffic := "test.csv"

	testData := []string{strconv.Itoa(today.Year()), today.Month().String(), strconv.Itoa(today.Day()), suffic}

	result := createFileName(suffic)
	for _, item := range testData {
		if strings.Contains(result, item) == false {
			t.Fail()
		}
	}
}
