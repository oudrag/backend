package events

import (
	"fmt"
	"io"
	"strconv"
)

// ****************************** Importance Enum ******************************

type Importance string

const (
	ImportanceVeryHigh Importance = "VeryHigh"
	ImportanceHigh     Importance = "High"
	ImportanceMedium   Importance = "Medium"
	ImportanceLow      Importance = "Low"
	ImportanceVeryLow  Importance = "VeryLow"
)

var AllImportance = []Importance{
	ImportanceVeryHigh,
	ImportanceHigh,
	ImportanceMedium,
	ImportanceLow,
	ImportanceVeryLow,
}

func (e Importance) IsValid() bool {
	switch e {
	case ImportanceVeryHigh, ImportanceHigh, ImportanceMedium, ImportanceLow, ImportanceVeryLow:
		return true
	}
	return false
}

func (e Importance) String() string {
	return string(e)
}

func (e *Importance) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Importance(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Importance", str)
	}
	return nil
}

func (e Importance) MarshalGQL(w io.Writer) {
	_, _ = fmt.Fprint(w, strconv.Quote(e.String()))
}

// ********************************* Mood Enum *********************************

type Mood string

const (
	MoodGreat    Mood = "Great"
	MoodGood     Mood = "Good"
	MoodMeh      Mood = "Meh"
	MoodBad      Mood = "Bad"
	MoodTerrible Mood = "Terrible"
)

var AllMood = []Mood{
	MoodGreat,
	MoodGood,
	MoodMeh,
	MoodBad,
	MoodTerrible,
}

func (e Mood) IsValid() bool {
	switch e {
	case MoodGreat, MoodGood, MoodMeh, MoodBad, MoodTerrible:
		return true
	}
	return false
}

func (e Mood) String() string {
	return string(e)
}

func (e *Mood) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Mood(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Mood", str)
	}
	return nil
}

func (e Mood) MarshalGQL(w io.Writer) {
	_, _ = fmt.Fprint(w, strconv.Quote(e.String()))
}

// ***************************** Event Action Enum *****************************

type EventAction string

const (
	EventActionDismissed EventAction = "Dismissed"
	EventActionDone      EventAction = "Done"
)

var AllEventAction = []EventAction{
	EventActionDismissed,
	EventActionDone,
}

func (e EventAction) IsValid() bool {
	switch e {
	case EventActionDismissed, EventActionDone:
		return true
	}
	return false
}

func (e EventAction) String() string {
	return string(e)
}

func (e *EventAction) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EventAction(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EventAction", str)
	}
	return nil
}

func (e EventAction) MarshalGQL(w io.Writer) {
	_, _ = fmt.Fprint(w, strconv.Quote(e.String()))
}
