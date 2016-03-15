package guber

type Namespace struct {
	*ResourceDefinition
	Metadata *Metadata `json:"metadata"`
}

type NamespaceList struct {
	Items []*Namespace `json:"items"`
}
