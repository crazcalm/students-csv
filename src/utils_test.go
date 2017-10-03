package students

import (
	"bytes"
	"github.com/artonge/go-csv-tag"
	"path/filepath"
	"strings"
	"testing"
)

const (
	GROUP5 = "Group 5:\n40 吉斯·霍华德 (ji2si1 huo4hua2de2) -- GEESE HOWARD\n41 比利·凯恩 (bi3li4 kai3en1) -- BILLY KANE\n42 安琪莉娜 (an1qi2li4na4) -- ANGELINA\n43 不知火舞 (bu4zhi1huo3wu3) -- MAI SHIRANUI\n44 麻宫雅典娜 (ma2gong1 ya3dian3na4) -- ATHENA ASAMIYA\n45 琼 (qiong2) -- KING"
)

func getTestData() Students {
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

func TestShuffle(t *testing.T) {
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

func TestPrintGroups(t *testing.T) {
	//Students
	ss := getTestData()

	//Num of groups
	g := 5

	//Shuffle option
	shuffle := false

	//Buffer to hold the data
	b := &bytes.Buffer{}

	PrintGroups(ss, g, b, shuffle)

	if strings.Contains(b.String(), GROUP5) != true {
		t.Fail()
	}
}
