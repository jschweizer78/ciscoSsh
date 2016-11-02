package main

// Ring is a model of ring info
type Ring struct {
	// DisplayName is the label of the file
	DisplayName string `xml:"DisplayName"`
	// FileName is the name of the file
	FileName string `xml:"FileName"`
}

// CiscoIPPhoneRingList is a list of Rings in te Ringlist.xml file
type CiscoIPPhoneRingList struct {
	Ring []Ring
}

type tInPut struct {
	kind int
	ver  bool
}
