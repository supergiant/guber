package guber

type ThirdPartyResource struct {
	Metadata struct {
		Name string
	}
	ApiVersion  string
	Kind        string
	Description string
	Versions    []struct {
		Name string
	}
}
