package main

type vertretung struct {
	Class     string
	Std       string
	Teacher   string
	Subject   string
	Room      string
	Type      string
	Notes     string
	Cancelled bool
}

// Credentials struct for importing credentials
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
}
