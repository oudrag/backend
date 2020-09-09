package platform

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

const (
	dateLayout  = "2006-01-02"
	clockLayout = "15:04"
)

type Date struct {
	time.Time
}

func CastToDate(i interface{}) (*Date, error) {
	s, ok := i.(string)
	if !ok {
		return nil, fmt.Errorf("input must be strings")
	}

	date, err := time.Parse(dateLayout, s)
	if err != nil {
		return nil, fmt.Errorf("invalid date format")
	}

	return &Date{date}, nil
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (d *Date) UnmarshalGQL(v interface{}) error {
	d, err := CastToDate(v)
	if err != nil {
		return err
	}

	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (d Date) MarshalGQL(w io.Writer) {
	_, _ = w.Write([]byte(strconv.Quote(d.Format(dateLayout))))
}

func (d Date) String() string {
	return d.Format(dateLayout)
}

type Clock struct {
	time.Time
}

func CastToClock(i interface{}) (*Clock, error) {
	s, ok := i.(string)
	if !ok {
		return nil, fmt.Errorf("input must be strings")
	}

	clock, err := time.Parse(clockLayout, s)
	if err != nil {
		return nil, fmt.Errorf("invalid clock format")
	}

	return &Clock{clock}, nil
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (c *Clock) UnmarshalGQL(v interface{}) error {
	c, err := CastToClock(v)
	if err != nil {
		return err
	}

	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (c Clock) MarshalGQL(w io.Writer) {
	_, _ = w.Write([]byte(strconv.Quote(c.Format(clockLayout))))
}

func (c Clock) String() string {
	return c.Format(clockLayout)
}
