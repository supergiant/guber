package guber

// ServiceCollection is a Collection interface for Services.
type ServiceCollection interface {
	Meta() *CollectionMeta
	New() *Service
	Create(e *Service) (*Service, error)
	Query(q *QueryParams) (*ServiceList, error)
	List() (*ServiceList, error)
	Get(name string) (*Service, error)
	Update(name string, r *Service) (*Service, error)
	Delete(name string) error
}

// Services implmenets ServiceCollection.
type Services struct {
	client    *RealClient
	Namespace string
}

// Meta implements the Collection interface.
func (c *Services) Meta() *CollectionMeta {
	return &CollectionMeta{
		DomainName: "",
		APIGroup:   "api",
		APIVersion: "v1",
		APIName:    "services",
		Kind:       "Service",
	}
}

func (c *Services) New() *Service {
	return &Service{
		collection: c,
	}
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
	if err := c.client.Get().Collection(c).Namespace(c.Namespace).Name(name).Do().Into(r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Services) Update(name string, r *Service) (*Service, error) {
	if err := c.client.Patch().Collection(c).Namespace(c.Namespace).Name(name).Entity(r).Do().Into(r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Services) Delete(name string) error {
	req := c.client.Delete().Collection(c).Namespace(c.Namespace).Name(name).Do()
	return req.err
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
	return r.collection.Delete(r.Metadata.Name)
}
