package guber

// NamespaceCollection is a Collection interface for Namespaces.
type NamespaceCollection interface {
	Meta() *CollectionMeta
	New() *Namespace
	Create(e *Namespace) (*Namespace, error)
	Query(q *QueryParams) (*NamespaceList, error)
	List() (*NamespaceList, error)
	Get(name string) (*Namespace, error)
	Update(name string, r *Namespace) (*Namespace, error)
	Delete(name string) error
}

// Namespaces implements NamespaceCollection.
type Namespaces struct {
	client *RealClient
}

// Meta implements the Collection interface.
func (c *Namespaces) Meta() *CollectionMeta {
	return &CollectionMeta{
		DomainName: "",
		APIGroup:   "api",
		APIVersion: "v1",
		APIName:    "namespaces",
		Kind:       "Namespace",
	}
}

func (c *Namespaces) New() *Namespace {
	return &Namespace{
		collection: c,
	}
}

func (c *Namespaces) Create(e *Namespace) (*Namespace, error) {
	r := c.New()
	if err := c.client.Post().Collection(c).Entity(e).Do().Into(r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Namespaces) Query(q *QueryParams) (*NamespaceList, error) {
	list := new(NamespaceList)
	if err := c.client.Get().Collection(c).Query(q).Do().Into(list); err != nil {
		return nil, err
	}
	for _, r := range list.Items {
		r.collection = c
	}
	return list, nil
}

func (c *Namespaces) List() (*NamespaceList, error) {
	list := new(NamespaceList)
	if err := c.client.Get().Collection(c).Do().Into(list); err != nil {
		return nil, err
	}
	for _, r := range list.Items {
		r.collection = c
	}
	return list, nil
}

func (c *Namespaces) Get(name string) (*Namespace, error) {
	r := c.New()
	if err := c.client.Get().Collection(c).Name(name).Do().Into(r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Namespaces) Update(name string, r *Namespace) (*Namespace, error) {
	if err := c.client.Patch().Collection(c).Name(name).Entity(r).Do().Into(r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Namespaces) Delete(name string) error {
	req := c.client.Delete().Collection(c).Name(name).Do()
	return req.err
}

// Resource-level

func (r *Namespace) Reload() (*Namespace, error) {
	return r.collection.Get(r.Metadata.Name)
}

func (r *Namespace) Save() error {
	_, err := r.collection.Update(r.Metadata.Name, r)
	return err
}

func (r *Namespace) Delete() error {
	return r.collection.Delete(r.Metadata.Name)
}
