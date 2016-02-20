package guber

type Version struct {
	Name string
}

type ThirdPartyResource struct {
	Metadata
	ApiVersion  string
	Kind        string
	Description string
	Versions    []Version
}
