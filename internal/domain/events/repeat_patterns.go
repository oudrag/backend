package events

import "github.com/oudrag/server/internal/platform"

type Repeat interface {
	IsRepeat()
}

type DailyPattern struct {
	EndsAt *platform.Date `json:"endsAt"`
}

func (DailyPattern) IsRepeat() {}

type WeeklyPattern struct {
	EndsAt     *platform.Date `json:"endsAt"`
	DaysOfWeek []int          `json:"daysOfWeek"`
}

func (WeeklyPattern) IsRepeat() {}

type MonthlyPattern struct {
	EndsAt      *platform.Date `json:"endsAt"`
	DaysOfMonth []int          `json:"daysOfMonth"`
}

func (MonthlyPattern) IsRepeat() {}

type YearlyPattern struct {
	EndsAt     *platform.Date `json:"endsAt"`
	DaysOfYear []*int         `json:"daysOfYear"`
}

func (YearlyPattern) IsRepeat() {}

type CustomPattern struct {
	EndsAt       *platform.Date `json:"endsAt"`
	IntervalDays int            `json:"intervalDays"`
}

func (CustomPattern) IsRepeat() {}
