package models

// Customer is an object repersenting a contact person
type Customer struct {
	id        string // Database ID first db will be firebase
	SfID      string // SalesForce ID for furture use.
	Name      string // Display Name for Customer. Should be Full registered Name
	DNSDomain string // registered DNS Domain name.
	Sites     []Site
}
