package pdevent

// Event is the type that we send to the event API
type Event struct {
	ServiceKey  string `json:"service_key"`
	EventType   string `json:"event_type"`
	IncidentKey string `json:"incident_key"`
	Description string `json:"description"`
	Client      string `json:"client"`
}
