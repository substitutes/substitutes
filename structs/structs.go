package structs

import "time"

// Substitute struct for holding data
type Substitute struct {
	Date            time.Time `json:"date"`
	Hour            int       `json:"hour"`
	Teacher         string    `json:"teacher"`
	TeacherInitials string    `json:"initials"`
	Subject         string    `json:"subject"`
	Type            string    `json:"type"`
	Notes           string    `json:"notes"`
	Classes         string    `json:"classes"`
	Room            string    `json:"room"`
	Cancelled       bool      `json:"cancelled"`
	New             bool      `json:"new"`
}

// Version struct for displaying current application versoin
type Version struct {
	Dirty   bool   `json:"dirty"`
	Hash    string `json:"hash"`
	Message string `json:"message"`
}

type SubstituteMeta struct {
	Date     time.Time `json:"date"`
	Class    string    `json:"class"`
	Extended bool      `json:"extended"`
	Updated  time.Time `json:"updated"`
}

type SubstituteResponse struct {
	Substitutes []Substitute   `json:"substitutes"`
	Meta        SubstituteMeta `json:"meta"`
}
