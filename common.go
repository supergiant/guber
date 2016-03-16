package guber

type ResourceDefinition struct {
	Kind       string `json:"kind"`
	ApiVersion string `json:"apiVersion"`
}

type Metadata struct {
	Name              string            `json:"name"`
	Labels            map[string]string `json:"labels,omitempty"`
	CreationTimestamp string            `json:"creationTimestamp,omitempty"`
}
