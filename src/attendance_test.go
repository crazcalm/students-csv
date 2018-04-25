package students

import (
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestCreateFileName(t *testing.T) {
	//Obtaining todays date
	today := time.Now()

	//suffic used for file
	suffic := "test.csv"

	testData := []string{strconv.Itoa(today.Year()), strconv.Itoa(int(today.Month())), strconv.Itoa(today.Day()), suffic}

	result := createFileName(suffic)
	for _, item := range testData {
		if !strings.Contains(result, item) {
			t.Errorf("Expected %s to be in %s, but it was not.", item, result)
		}
	}
}
