package pdevent

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Client is the main object for communicating with the event API
type Client struct {
	serviceKey string
	clientName string
}

func NewClient(key string, name string) Client {
	return Client{serviceKey: key, clientName: name}
}

func (c Client) Trigger(desc string, key string) {
	e := Event{
		ServiceKey:  c.serviceKey,
		EventType:   "trigger",
		IncidentKey: key,
		Description: desc,
		Client:      c.clientName,
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(e)
	http.Post(
		"https://events.pagerduty.com/generic/2010-04-15/create_event.json",
		"application/json",
		b)
}
