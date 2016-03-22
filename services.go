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
	err := r.client.Post().Resource(r).Namespace(r.Namespace).Entity(e).Do().Into(e)
	return e, err
}

func (r *Services) List() (*ServiceList, error) {
	list := new(ServiceList)
	err := r.client.Get().Resource(r).Namespace(r.Namespace).Do().Into(list)
	return list, err
}

func (r *Services) Get(name string) (*Service, error) {
	e := new(Service)
	err := r.client.Get().Resource(r).Namespace(r.Namespace).Name(name).Do().Into(e)
	return e, err
}

func (r *Services) Update(name string, e *Service) (*Service, error) {
	err := r.client.Patch().Resource(r).Namespace(r.Namespace).Name(name).Entity(e).Do().Into(e)
	return e, err
}

func (r *Services) Delete(name string) (found bool, err error) {
	req := r.client.Delete().Resource(r).Namespace(r.Namespace).Name(name).Do()
	return req.found, req.err
}
