package guber

type Services struct {
	client    *Client
	Namespace string
}

func (r Services) DomainName() string {
	return ""
}

func (r Services) ApiGroup() string {
	return "api"
}

func (r Services) ApiVersion() string {
	return "v1"
}

func (r Services) ApiName() string {
	return "services"
}

func (r Services) Kind() string {
	return "Service"
}

func (r *Services) Create(e *Service) (*Service, error) {
	if err := r.client.Post().Resource(r).Namespace(r.Namespace).Entity(e).Do().Into(e); err != nil {
		return nil, err
	}
	return e, nil
}

func (r *Services) List(q *QueryParams) (*ServiceList, error) {
	list := new(ServiceList)
	err := r.client.Get().Resource(r).Namespace(r.Namespace).Query(q).Do().Into(list)
	return list, err
}

func (r *Services) Get(name string) (*Service, error) {
	e := new(Service)
	req := r.client.Get().Resource(r).Namespace(r.Namespace).Name(name).Do()
	if err := req.Into(e); err != nil {
		return nil, err
	}
	if req.found {
		return e, nil
	}
	return nil, nil
}

func (r *Services) Update(name string, e *Service) (*Service, error) {
	if err := r.client.Patch().Resource(r).Namespace(r.Namespace).Name(name).Entity(e).Do().Into(e); err != nil {
		return nil, err
	}
	return e, nil
}

func (r *Services) Delete(name string) (found bool, err error) {
	req := r.client.Delete().Resource(r).Namespace(r.Namespace).Name(name).Do()
	return req.found, req.err
}
