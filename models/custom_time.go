package models

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

const timeFormat = "2006-01-02 15:04:05"

type CustomTime time.Time

func (ct *CustomTime) UnmarshalJSON(data []byte) error {
    newTime, err := time.Parse(timeFormat, strings.Trim(string(data), "\""))
    if err != nil {
        return err
    }

    *ct = CustomTime(newTime)
    return nil
}

func (ct *CustomTime) MarshalJSON() ([]byte, error) {
    return []byte(fmt.Sprintf("%q", time.Time(*ct).Format(timeFormat))), nil
}

func (ct CustomTime) Value() (driver.Value, error) {
    return time.Time(ct), nil
}

func (ct *CustomTime) Scan(src interface{}) error {
    if val, ok := src.(time.Time); ok {
        *ct = CustomTime(val)
    } else {
        return errors.New("time Scanner passed a non-time object")
    }

    return nil
}