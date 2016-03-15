package guber

type Source struct {
	Host string `json:"host"`
}

type Event struct {
	*ResourceDefinition
	Message string  `json:"message"`
	Count   int     `json:"count"`
	Source  *Source `json:"source"`
}

type EventList struct {
	Items []*Event `json:"items"`
}
