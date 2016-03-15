package guber

type Secret struct {
	*ResourceDefinition
	Metadata *Metadata         `json:"metadata"`
	Type     string            `json:"type"`
	Data     map[string]string `json:"data"`
}

type SecretList struct {
	Items []*Secret `json:"items"`
}
