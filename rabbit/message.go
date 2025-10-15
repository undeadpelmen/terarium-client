package rabbit

import "time"

type Mesage struct {
	Mac string
	Message string
	Time time.Time
}

func (m *Mesage) JSON() string {
	return `{"mac":"` + m.Mac + `","message":"` + m.Message + `","time":"` + m.Time.Format("2006-01-02 15:04:05") + `"}`
}