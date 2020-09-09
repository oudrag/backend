package events

import (
	"time"

	"github.com/oudrag/server/internal/platform"
)

type Event struct {
	Title      string         `json:"title"`
	Date       platform.Date  `json:"date"`
	Time       platform.Clock `json:"time"`
	Importance *Importance    `json:"importance"`
	Repeat     Repeat         `json:"repeat"`
	Category   *Category      `json:"category"`
	Tags       []*Tag         `json:"tags"`
	History    []*EventLog    `json:"history"`
}

type Tag struct {
	Title string  `json:"title"`
	Color *string `json:"color"`
}

type Category struct {
	Title string  `json:"title"`
	Icon  *string `json:"icon"`
	Color *string `json:"color"`
}

type EventLog struct {
	Action      EventAction `json:"action"`
	Mood        *Mood       `json:"mood"`
	Description *string     `json:"description"`
	SpentTime   *int        `json:"spentTime"`
	Datetime    time.Time   `json:"datetime"`
}
