package models

// GFile is an object repersenting a Google Drive file
type GFile struct {
	ID   string
	GID  string
	Name string
	URL  string
}

// GSheet is an struct repersenting a Google SpreadSheet file
type GSheet struct {
	GFile
	Sheets []Sheet
}

// Sheet is a struct repersenting a sheet within a spreadsheet
type Sheet struct {
	Name string
}

// GDoc is an object repersenting a Google Document file
type GDoc struct {
	GFile
	Title string
}

// GetGID returns a Google Drive file ID as a string
func (gf *GFile) GetGID() string {
	// TODO add function to parse URL to set this if nil
	return gf.GID
}

// GetName returns a Google Drive file Name as a string
func (gf *GFile) GetName() string {

	return gf.Name
}
