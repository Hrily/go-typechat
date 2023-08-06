package models

import "github.com/hrily/go-typechat"

func (c *CalendarActions) Validate() error {
	for _, actions := range c.Actions {
		if err := actions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (a *Action) Validate() error {
	if a.AddEvent != nil {
		return a.AddEvent.Validate()
	}
	return nil
}

func (a *AddEventAction) Validate() error {
	if a.Event != nil {
		return a.Event.Validate()
	}
	return nil
}

func (e *Event) Validate() error {
	if e.TimeRange != nil {
		return e.TimeRange.Validate()
	}
	return nil
}

var (
	invalidTimeRange = typechat.ValidationError{
		Message: "invalid time range: required format is HH:MM in 24 hour format",
	}
)

func (t *EventTimeRange) Validate() error {
	// very basic validation
	if len(t.StartTime) != 5 {
		return invalidTimeRange
	}
	if len(t.EndTime) != 5 {
		return invalidTimeRange
	}
	return nil
}
