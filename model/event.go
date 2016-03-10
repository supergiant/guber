package model

type Source struct {
	Host string `json:"host"`
}

type Event struct {
	Message string  `json:"message"`
	Count   int     `json:"count"`
	Source  *Source `json:"source"`
}

type EventList struct {
	Items []*Event `json:"items"`
}
