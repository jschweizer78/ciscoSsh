package main

func makeSingleLinePhone(dt string, data [][]string) {
	devData := filterDeviceFromData(dt, data)
	// TODO remove line(s) 2-5 from data[][]string
	switch dt {
	case "CTI Port":
		go makeCTIPortCsv(devData)
	case "Jabber":
		go makeJabberCSFCsv(devData)
	case "Analog":
		go makeAnalogPortCsv(devData)
	case "Cisco 8831":
		go makeConfphoneCsv(data)
	}
}

func makeMultiLinePhone(dt string, data [][]string) {
	// devData := filterDeviceFromData(dt, data)
}

func makeCTIPortCsv(data [][]string) {

	// TODO convert CIPT headers to CUCM and add needed headers
}

func makeJabberCSFCsv(data [][]string) {

}

func makeAnalogPortCsv(data [][]string) {

}

func makeConfphoneCsv(data [][]string) {
}

func convertHeaderToCUCM(ciptHeader string) string {
	if isUsedInCUCMHeaders(ciptHeader) {
		return "UNUSED"
	}

	if isUsedInFormulas(ciptHeader) {
		return "FORMULA"
	}
	switch ciptHeader {
	case `Display / Alerting Name`:
		return "Description"
	case `Button Template`:
		return "Phone Button Template"
	case `Softkey Template`:
		return "Softkey Template"
	case `CWID`:
		return "Owner User ID"
	case `Line 1`:
		return "Directory Number 1"
	case `Class of Service`:
		return "Line CSS 1"
	case `Line 1 text label`:
		return "Line Text Label 1"
	case `External Phone Number Mask`:
		return "External Phone Number Mask 1"
	case `Line 2`:
		return "Directory Number 2"
	case `Line 3 text label`:
		return "Line Text Label 3"
	case `Line 4`:
		return "Directory Number 4"
	case `Line 5 text label`:
		return "Line Text Label 5"
	default:
		return "UNKNOWN FIELD NAME"
	}
}

func isUsedInCUCMHeaders(header string) bool {
	switch header {
	case `Comments / Special Routing Instructions`, `Changed post Import`, `Imported - VM`, `Imported - Phone`, `Department`, `Building / Location`, `Supervisor / Data Contact`, `First Name`, `Last Name`, `Title or CR Number`, `DID`, `DN (dialed num)`, `UM`, `VM (no VM if blank)`, `VM Zero Out`, ``:
		return false
	default:
		return true
	}
}

func isUsedInFormulas(header string) bool {
	switch header {
	case `Site`:
		return true
	default:
		return false
	}
}
