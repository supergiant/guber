package guber

type ResourceDefinition struct {
	Kind       string `json:"kind"`
	ApiVersion string `json:"apiVersion"`
}

type Metadata struct {
	Name              string
	Labels            map[string]string
	CreationTimestamp string
}
