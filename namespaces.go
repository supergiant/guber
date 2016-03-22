package guber

type Namespaces struct {
	client *Client
}

func (r Namespaces) DomainName() string {
	return ""
}

func (r Namespaces) ApiGroup() string {
	return "api"
}

func (r Namespaces) ApiVersion() string {
	return "v1"
}

func (r Namespaces) ApiName() string {
	return "namespaces"
}

func (r Namespaces) Kind() string {
	return "Namespace"
}

func (r *Namespaces) Create(e *Namespace) (*Namespace, error) {
	if err := r.client.Post().Resource(r).Entity(e).Do().Into(e); err != nil {
		return nil, err
	}
	return e, nil
}

func (r *Namespaces) List(q *QueryParams) (*NamespaceList, error) {
	list := new(NamespaceList)
	err := r.client.Get().Resource(r).Query(q).Do().Into(list)
	return list, err
}

func (r *Namespaces) Get(name string) (*Namespace, error) {
	e := new(Namespace)
	if err := r.client.Get().Resource(r).Name(name).Do().Into(e); err != nil {
		return nil, err
	}
	return e, nil
}

func (r *Namespaces) Update(name string, e *Namespace) (*Namespace, error) {
	if err := r.client.Patch().Resource(r).Name(name).Entity(e).Do().Into(e); err != nil {
		return nil, err
	}
	return e, nil
}

func (r *Namespaces) Delete(name string) (found bool, err error) {
	req := r.client.Delete().Resource(r).Name(name).Do()
	return req.found, req.err
}
