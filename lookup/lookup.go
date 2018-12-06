package lookup

import (
	"encoding/csv"
	"github.com/lunny/log"
	"os"
	"strings"
)

type Lookup struct {
	RawRecords [][]string
	Lookup     map[string]string
	path       string
}

//TeacherLookup is the deafult Lookup table, created on New()
var TeacherLookup *Lookup

func (l *Lookup) ReadFile() {
	fileInfo, err := os.Stat("./lookup/teachers.csv")
	if err != nil {
		if os.IsNotExist(err) {
			log.Warn("Teachers lookup table does not exist - ignoring!")
			return
		} else {
			log.Fatal("Failed to read teachers lookup file!")
			return
		}
	}
	log.Debug("Attempting to read: ", fileInfo.Name())
	file, err := os.Open("./lookup/teachers.csv") // get absolute path
	if err != nil {
		log.Fatal("Failed to read file: ", err)
		return
	}
	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Failed to read file - invalid CSV? ", err)
	}
	l.RawRecords = records
	for _, record := range records {
		l.Lookup[removeSpaces(record[2])] = removeSpaces(record[1])
	}
	log.Info("Loaded teachers")
}

func New() *Lookup {
	l := &Lookup{Lookup: map[string]string{}}
	TeacherLookup = l
	return TeacherLookup
}

func removeSpaces(s string) string {
	return strings.Replace(s, " ", "", -1)
}

func (l *Lookup) Get(s string) string {
	if len(s) >= 2 && len(s) <= 3 {
		l.GetRaw(s)
	}
	// Split
	if strings.Contains(s, "=>") {
		return l.GetRaw(strings.Split(s, " => ")[0]) + " => " + l.GetRaw(strings.Split(s, " => ")[1])
	} else {
		return l.GetRaw(s)
	}
}

func (l *Lookup) GetRaw(s string) string {
	if l.Lookup[s] == "" {
		return s
	}
	return l.Lookup[s]
}
