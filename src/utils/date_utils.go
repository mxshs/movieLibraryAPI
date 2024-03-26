package utils

import (
	"database/sql/driver"
	"fmt"
	"time"
)

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

func (d *Date) MarshalJSON() ([]byte, error) {
	return []byte(d.Format(`"02.01.2006"`)), nil
}

func (d Date) Value() (driver.Value, error) {
	return d.Time, nil
}

func (d *Date) Scan(v any) error {
	if v == nil {
		d.Time = time.Time{}
		return nil
	}

	if time, ok := v.(time.Time); ok {
		d.Time = time
		return nil
	}

	return fmt.Errorf("cannot convert %v to time", v)
}
