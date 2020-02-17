package vanilla

import (
	"encoding/json"
	"time"
)

type duration struct {
	time.Duration
}

func (d *duration) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' {
		sd := string(b[1 : len(b)-1])
		d.Duration, err = time.ParseDuration(sd)
		return
	}

	var id int64
	id, err = json.Number(string(b)).Int64()
	d.Duration = time.Duration(id)
	return
}

func (d duration) MarshalJSON() (b []byte, err error) {
	return []byte(d.String()), nil
}

type timely struct {
	Interval duration `json:"interval,omitempty"`
	Accrual  uint     `json:"accrual,omitempty"`
}

type config struct {
	Timely timely `json:"timely,omitempty"`
}
