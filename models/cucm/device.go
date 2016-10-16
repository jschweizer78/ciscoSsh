package cucm

// DeviceSummary is general device
type DeviceSummary struct {
	Name    string   `json:"name"`
	Type    string   `json:"type"`
	LineIDs []string `json:"lineIDs"`
}

// NewDeviceSumm used to properly create a new DeviceSummary struct (object)
func NewDeviceSumm(name, kind string) DeviceSummary {
	item := DeviceSummary{Name: name, Type: kind}
	item.LineIDs = make([]string, 1)
	return item
}
