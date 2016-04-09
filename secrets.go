package guber

type Secrets struct {
	client    *Client
	Namespace string
}

func (c *Secrets) New() *Secret {
	return &Secret{
		collection: c,
	}
}

func (c Secrets) DomainName() string {
	return ""
}

func (c Secrets) APIGroup() string {
	return "api"
}

func (c Secrets) APIVersion() string {
	return "v1"
}

func (c Secrets) APIName() string {
	return "secrets"
}

func (c Secrets) Kind() string {
	return "Secret"
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
	req := c.client.Get().Collection(c).Namespace(c.Namespace).Name(name).Do()
	if err := req.Into(r); err != nil {
		return nil, err
	}
	if req.found {
		return r, nil
	}
	return nil, nil
}

func (c *Secrets) Update(name string, r *Secret) (*Secret, error) {
	if err := c.client.Patch().Collection(c).Namespace(c.Namespace).Name(name).Entity(r).Do().Into(r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Secrets) Delete(name string) (found bool, err error) {
	req := c.client.Delete().Collection(c).Namespace(c.Namespace).Name(name).Do()
	return req.found, req.err
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
	_, err := r.collection.Delete(r.Metadata.Name)
	return err
}
