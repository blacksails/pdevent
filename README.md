# pdevent

pdevent is a golibrary for communicating with the [Pagerduty Event API](https://v2.developer.pagerduty.com/docs/events-api).

## Usage example

```go
package main

func main() {
  // Initiate client with pagerduty service key and give it a name.
  c := pdevent.NewClient("PAGERDUTY_SERVICEKEY", "Go pagerduty client")
  
  // Trigger an event and optionally give it an incident key to be recognized by
  c.Trigger("Something happened", "Incident Key")
  
  // Acknowledge or resolve an event using the incident key
  c.Acknowledge("Incident Key")
  c.Resolve("Incident Key")
}
```

## Work in progress

The library is work in progress. New backwards incompatible releases will be tagged using semver.
