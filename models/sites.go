package models

// Site is an struct repersenting a customer site
type Site struct {
	Name      string
	Address   PhysAddres
	ShortName string
	Owner     Customer
	Contacts  []*Contact
}

// PhysAddres is an struct repersenting a physical address
type PhysAddres struct {
	Street1 string
	Street2 string
	City    string
	State   string
	Zip     string
}
