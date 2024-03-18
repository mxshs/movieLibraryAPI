package utils

import "time"

type Date struct {
	time.Time
}

func (d *Date) UnmarshalJSON(b []byte) error {
	date, err := time.Parse(`"02.01.2006"`, string(b))
	if err != nil {
		return err
	}

	d.Time = date
	return nil
}
