package pdevent

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Event is the type that we send to the event API
type Event struct {
	ServiceKey  string `json:"service_key"`
	EventType   string `json:"event_type"`
	IncidentKey string `json:"incident_key"`
	Description string `json:"description,omitempty"`
	Client      string `json:"client,omitempty"`
}

// Client is the main object for communicating with the event API
type Client struct {
	serviceKey string
	clientName string
}

// NewClient initiates a client with the given service key and name
func NewClient(key string, name string) Client {
	return Client{serviceKey: key, clientName: name}
}

func (c Client) sendEvent(e Event) error {
	const (
		ep = "https://events.pagerduty.com/generic/2010-04-15/create_event.json"
		bt = "application/json"
	)
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(e)
	resp, err := http.Post(ep, bt, b)
	if err != nil || resp.StatusCode != http.StatusOK {
		return err
	}
	return nil
}

// Trigger sends a trigger event, with the given description and incident key
func (c Client) Trigger(desc string, key string) error {
	e := Event{
		ServiceKey:  c.serviceKey,
		EventType:   "trigger",
		IncidentKey: key,
		Description: desc,
		Client:      c.clientName,
	}
	return c.sendEvent(e)
}

// Acknowledge sends an acknowledge event with the given incident key
func (c Client) Acknowledge(key string) error {
	e := Event{
		ServiceKey:  c.serviceKey,
		EventType:   "acknowledge",
		IncidentKey: key,
	}
	return c.sendEvent(e)
}

// Resolve sends a resolve event with the given incident key
func (c Client) Resolve(key string) error {
	e := Event{
		ServiceKey:  c.serviceKey,
		EventType:   "resolve",
		IncidentKey: key,
	}
}
