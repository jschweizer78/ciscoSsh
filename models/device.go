package models

// Device is an object repersenting a network device i.e a switch or router
type Device struct {
	ID         string `json:"id"`
	Host       string `json:"host"`
	Priv       int    `json:"priv"`
	SSHEnabled bool   `json:"ssh"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	EnablePass string `json:"enablePass"`
}

// Switch is a model of a switch info for login
type Switch struct {
	IDF          string
	IPAddress    string
	Room         string
	Hostname     string
	Model        string
	Username     string
	Password     string
	IsSSHEnabled bool
	Site         string
}
