package guber

// Namespaces is a kubernetes namespace holder for a Client object.
type Namespaces struct {
	client *Client
}

// DomainName returns domain name
func (r Namespaces) DomainName() string {
	return ""
}

// APIGroup returns the apigroup used with namespaces.
func (r Namespaces) APIGroup() string {
	return defaultAPIGroup
}

// APIVersion returns the apiversion used with namespaces.
func (r Namespaces) APIVersion() string {
	return defaultAPIVersion
}

// APIName returns the apiname "namespaces".
func (r Namespaces) APIName() string {
	return "namespaces"
}

// Kind returns the namespace kind "Namespace".
func (r Namespaces) Kind() string {
	return "Namespace"
}

// Create creates a new kubernetes namespace and returns a Namespace object.
func (r *Namespaces) Create(e *Namespace) (*Namespace, error) {
	if err := r.client.Post().Resource(r).Entity(e).Do().Into(e); err != nil {
		return nil, err
	}
	return e, nil
}

// Query searches kubernetes namespaces and returns a list NamespaceList object.
func (r *Namespaces) Query(q *QueryParams) (*NamespaceList, error) {
	list := new(NamespaceList)
	err := r.client.Get().Resource(r).Query(q).Do().Into(list)
	return list, err
}

// List returns a kubenrtes namespace list NamespaceList object.
func (r *Namespaces) List() (*NamespaceList, error) {
	list := new(NamespaceList)
	err := r.client.Get().Resource(r).Do().Into(list)
	return list, err
}

// Get fetches a kuebrnetes namespace and returns it as a Namespace object.
func (r *Namespaces) Get(name string) (*Namespace, error) {
	e := new(Namespace)
	req := r.client.Get().Resource(r).Name(name).Do()
	if err := req.Into(e); err != nil {
		return nil, err
	}
	if req.found {
		return e, nil
	}
	return nil, nil
}

// Update updates a kubernetes namespace and returns the updated namespace Namespace object.
func (r *Namespaces) Update(name string, e *Namespace) (*Namespace, error) {
	if err := r.client.Patch().Resource(r).Name(name).Entity(e).Do().Into(e); err != nil {
		return nil, err
	}
	return e, nil
}

// Delete deletes a kuebrnetes namespace.
func (r *Namespaces) Delete(name string) (found bool, err error) {
	req := r.client.Delete().Resource(r).Name(name).Do()
	return req.found, req.err
}
