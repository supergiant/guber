package guber

type Secrets struct {
	client    *Client
	Namespace string
}

func (r Secrets) DomainName() string {
	return ""
}

func (r Secrets) APIGroup() string {
	return "api"
}

func (r Secrets) APIVersion() string {
	return "v1"
}

func (r Secrets) APIName() string {
	return "secrets"
}

func (r Secrets) Kind() string {
	return "Secret"
}

func (r *Secrets) Create(e *Secret) (*Secret, error) {
	if err := r.client.Post().Resource(r).Namespace(r.Namespace).Entity(e).Do().Into(e); err != nil {
		return nil, err
	}
	return e, nil
}

func (r *Secrets) Query(q *QueryParams) (*SecretList, error) {
	list := new(SecretList)
	err := r.client.Get().Resource(r).Namespace(r.Namespace).Query(q).Do().Into(list)
	return list, err
}

func (r *Secrets) List() (*SecretList, error) {
	list := new(SecretList)
	err := r.client.Get().Resource(r).Namespace(r.Namespace).Do().Into(list)
	return list, err
}

func (r *Secrets) Get(name string) (*Secret, error) {
	e := new(Secret)
	req := r.client.Get().Resource(r).Namespace(r.Namespace).Name(name).Do()
	if err := req.Into(e); err != nil {
		return nil, err
	}
	if req.found {
		return e, nil
	}
	return nil, nil
}

func (r *Secrets) Update(name string, e *Secret) (*Secret, error) {
	if err := r.client.Patch().Resource(r).Namespace(r.Namespace).Name(name).Entity(e).Do().Into(e); err != nil {
		return nil, err
	}
	return e, nil
}

func (r *Secrets) Delete(name string) (found bool, err error) {
	req := r.client.Delete().Resource(r).Namespace(r.Namespace).Name(name).Do()
	return req.found, req.err
}
