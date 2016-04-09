package guber

type Services struct {
	client    *Client
	Namespace string
}

func (c *Services) New() *Service {
	return &Service{
		collection: c,
	}
}

func (c Services) DomainName() string {
	return ""
}

func (c Services) APIGroup() string {
	return "api"
}

func (c Services) APIVersion() string {
	return "v1"
}

func (c Services) APIName() string {
	return "services"
}

func (c Services) Kind() string {
	return "Service"
}

func (c *Services) Create(e *Service) (*Service, error) {
	r := c.New()
	if err := c.client.Post().Collection(c).Namespace(c.Namespace).Entity(e).Do().Into(r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Services) Query(q *QueryParams) (*ServiceList, error) {
	list := new(ServiceList)
	if err := c.client.Get().Collection(c).Namespace(c.Namespace).Query(q).Do().Into(list); err != nil {
		return nil, err
	}
	for _, r := range list.Items {
		r.collection = c
	}
	return list, nil
}

func (c *Services) List() (*ServiceList, error) {
	list := new(ServiceList)
	if err := c.client.Get().Collection(c).Namespace(c.Namespace).Do().Into(list); err != nil {
		return nil, err
	}
	for _, r := range list.Items {
		r.collection = c
	}
	return list, nil
}

func (c *Services) Get(name string) (*Service, error) {
	r := c.New()
	req := c.client.Get().Collection(c).Namespace(c.Namespace).Name(name).Do()
	if err := req.Into(r); err != nil {
		return nil, err
	}
	if req.found {
		return r, nil
	}
	return nil, nil
}

func (c *Services) Update(name string, r *Service) (*Service, error) {
	if err := c.client.Patch().Collection(c).Namespace(c.Namespace).Name(name).Entity(r).Do().Into(r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Services) Delete(name string) (found bool, err error) {
	req := c.client.Delete().Collection(c).Namespace(c.Namespace).Name(name).Do()
	return req.found, req.err
}

// Resource-level

func (r *Service) Reload() (*Service, error) {
	return r.collection.Get(r.Metadata.Name)
}

func (r *Service) Save() error {
	_, err := r.collection.Update(r.Metadata.Name, r)
	return err
}

func (r *Service) Delete() error {
	_, err := r.collection.Delete(r.Metadata.Name)
	return err
}
