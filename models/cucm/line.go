package cucm

import "fmt"

// LineDevSummary is join of Line and Device data
type LineDevSummary struct {
	ID        string   `json:"id"`
	DirNum    string   `json:"dirNum" binding:"required"`
	Partition string   `json:"partition" binding:"required"`
	Model     string   `json:"model" binding:"required"`
	DeviceIDs []string `json:"deviceIDs"` // NOTE for analog device name is the full namte of gateway and port num
}

// NewLineSummary is used to create new line/parttion and assassocatiated devices
func NewLineSummary(dn, part string) LineDevSummary {
	item := LineDevSummary{DirNum: dn, Partition: part}
	item.setID()
	item.DeviceIDs = make([]string, 5)
	return item
}

func (lds LineDevSummary) setID() {
	lds.ID = fmt.Sprintf("%s %s", lds.DirNum, lds.Partition)
}
