package structs

// Substitutes struct for holding data
type Substitutes struct {
	Class     string `json:"class"`
	Hour      string `json:"hour"`
	Teacher   string `json:"teacher"`
	Subject   string `json:"subject"`
	Room      string `json:"room"`
	Type      string `json:"type"`
	Notes     string `json:"notes"`
	Cancelled bool   `json:"cancelled"`
}

// Credentials struct for importing credentials
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
}
