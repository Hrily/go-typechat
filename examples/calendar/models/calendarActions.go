package models

type CalendarActions struct {
	Actions []*Action `json:"actions,omitempty"`
}

type Action struct {
	AddEvent        *AddEventAction        `json:"addEvent,omitempty"`
	RemoveEvent     *RemoveEventAction     `json:"removeEvent,omitempty"`
	FindEvents      *FindEventsAction      `json:"findEvents,omitempty"`
	AddParticipants *AddParticipantsAction `json:"addParticipants,omitempty"`
	Unknown         *UnknownAction         `json:"unknown,omitempty"`
}

type AddEventAction struct {
	Event *Event `json:"event,omitempty"`
}

type RemoveEventAction struct {
	EventReference *EventReference `json:"eventReference,omitempty"`
}

type FindEventsAction struct {
	// one or more event properties to use to search for matching events
	EventReference *EventReference `json:"eventReference,omitempty"`
}

type AddParticipantsAction struct {
	EventReference *EventReference `json:"eventReference,omitempty"`
	Participants   []string        `json:"participants,omitempty"`
}

// UnknownAction is used when the user types text that can not easily be
// understood as a calendar action
type UnknownAction struct {
	// Text typed by the user that the system did not understand
	Text string `json:"text,omitempty"`
}

type Event struct {
	// Day (example: March 22, 2024) or relative date (example: after EventReference)
	Day         string          `json:"day,omitempty"`
	TimeRange   *EventTimeRange `json:"timeRange,omitempty"`
	Description string          `json:"description,omitempty"`
	Location    string          `json:"location,omitempty"`
	// Participants is list of people or named groups like 'team'
	Participants []string `json:"participants,omitempty"`
}

// EventReference is properties used by the requester in referring to an event
// these properties are only specified if given directly by the requester
type EventReference struct {
	// Day (example: March 22, 2024) or relative date (example: after EventReference)
	Day string `json:"day,omitempty"`
	// DayRange (examples: this month, this week, in the next two days)
	DayRange     string          `json:"dayRange,omitempty"`
	TimeRange    *EventTimeRange `json:"timeRange,omitempty"`
	Description  string          `json:"description,omitempty"`
	Location     string          `json:"location,omitempty"`
	Participants []string        `json:"participants,omitempty"`
}

type EventTimeRange struct {
	StartTime string `json:"startTime,omitempty"`
	EndTime   string `json:"endTime,omitempty"`
	Duration  string `json:"duration,omitempty"`
}
