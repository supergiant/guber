package guber

// SecretCollection is a Collection interface for Secrets.
type SecretCollection interface {
	Meta() *CollectionMeta
	New() *Secret
	Create(e *Secret) (*Secret, error)
	Query(q *QueryParams) (*SecretList, error)
	List() (*SecretList, error)
	Get(name string) (*Secret, error)
	Update(name string, r *Secret) (*Secret, error)
	Delete(name string) error
}

// Secrets implmenets SecretCollection.
type Secrets struct {
	client    *RealClient
	Namespace string
}

// Meta implements the Collection interface.
func (c *Secrets) Meta() *CollectionMeta {
	return &CollectionMeta{
		DomainName: "",
		APIGroup:   "api",
		APIVersion: "v1",
		APIName:    "secrets",
		Kind:       "Secret",
	}
}

func (c *Secrets) New() *Secret {
	return &Secret{
		collection: c,
	}
}

func (c *Secrets) Create(e *Secret) (*Secret, error) {
	r := c.New()
	if err := c.client.Post().Collection(c).Namespace(c.Namespace).Entity(e).Do().Into(r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Secrets) Query(q *QueryParams) (*SecretList, error) {
	list := new(SecretList)
	if err := c.client.Get().Collection(c).Namespace(c.Namespace).Query(q).Do().Into(list); err != nil {
		return nil, err
	}
	for _, r := range list.Items {
		r.collection = c
	}
	return list, nil
}

func (c *Secrets) List() (*SecretList, error) {
	list := new(SecretList)
	if err := c.client.Get().Collection(c).Namespace(c.Namespace).Do().Into(list); err != nil {
		return nil, err
	}
	for _, r := range list.Items {
		r.collection = c
	}
	return list, nil
}

func (c *Secrets) Get(name string) (*Secret, error) {
	r := c.New()
	if err := c.client.Get().Collection(c).Namespace(c.Namespace).Name(name).Do().Into(r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Secrets) Update(name string, r *Secret) (*Secret, error) {
	if err := c.client.Patch().Collection(c).Namespace(c.Namespace).Name(name).Entity(r).Do().Into(r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Secrets) Delete(name string) error {
	req := c.client.Delete().Collection(c).Namespace(c.Namespace).Name(name).Do()
	return req.err
}

// Resource-level

func (r *Secret) Reload() (*Secret, error) {
	return r.collection.Get(r.Metadata.Name)
}

func (r *Secret) Save() error {
	_, err := r.collection.Update(r.Metadata.Name, r)
	return err
}

func (r *Secret) Delete() error {
	return r.collection.Delete(r.Metadata.Name)
}
