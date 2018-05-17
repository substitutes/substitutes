package structs

// Substitutes struct for holding data
type Substitutes struct {
	Date      string `json:"date"`
	Hour      string `json:"hour"`
	Day       string `json:"day"`
	Teacher   string `json:"teacher"`
	Time      string `json:"time"`
	Subject   string `json:"subject"`
	Type      string `json:"type"`
	Notes     string `json:"notes"`
	Classes   string `json:"classes"`
	Room      string `json:"room"`
	After     string `json:"after"`
	Cancelled bool   `json:"cancelled"`
	New       bool   `json:"new"`
	Reason    string `json:"reason"`
	Counter   string `json:"counter"`
}

// Credentials struct for importing credentials
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
}
