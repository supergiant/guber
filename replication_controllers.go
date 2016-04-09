package guber

type ReplicationControllers struct {
	client    *Client
	Namespace string
}

func (c *ReplicationControllers) New() *ReplicationController {
	return &ReplicationController{
		collection: c,
	}
}

func (c ReplicationControllers) DomainName() string {
	return ""
}

func (c ReplicationControllers) APIGroup() string {
	return "api"
}

func (c ReplicationControllers) APIVersion() string {
	return "v1"
}

func (c ReplicationControllers) APIName() string {
	return "replicationcontrollers"
}

func (c ReplicationControllers) Kind() string {
	return "ReplicationController"
}

func (c *ReplicationControllers) Create(e *ReplicationController) (*ReplicationController, error) {
	r := c.New()
	if err := c.client.Post().Collection(c).Namespace(c.Namespace).Entity(e).Do().Into(r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *ReplicationControllers) Query(q *QueryParams) (*ReplicationControllerList, error) {
	list := new(ReplicationControllerList)
	if err := c.client.Get().Collection(c).Namespace(c.Namespace).Query(q).Do().Into(list); err != nil {
		return nil, err
	}
	for _, r := range list.Items {
		r.collection = c
	}
	return list, nil
}

func (c *ReplicationControllers) List() (*ReplicationControllerList, error) {
	list := new(ReplicationControllerList)
	if err := c.client.Get().Collection(c).Namespace(c.Namespace).Do().Into(list); err != nil {
		return nil, err
	}
	for _, r := range list.Items {
		r.collection = c
	}
	return list, nil
}

func (c *ReplicationControllers) Get(name string) (*ReplicationController, error) {
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

func (c *ReplicationControllers) Update(name string, r *ReplicationController) (*ReplicationController, error) {
	if err := c.client.Patch().Collection(c).Namespace(c.Namespace).Name(name).Entity(r).Do().Into(r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *ReplicationControllers) Delete(name string) (found bool, err error) {
	req := c.client.Delete().Collection(c).Namespace(c.Namespace).Name(name).Do()
	return req.found, req.err
}

// Resource-level

func (r *ReplicationController) Reload() (*ReplicationController, error) {
	return r.collection.Get(r.Metadata.Name)
}

func (r *ReplicationController) Save() error {
	_, err := r.collection.Update(r.Metadata.Name, r)
	return err
}

func (r *ReplicationController) Delete() error {
	_, err := r.collection.Delete(r.Metadata.Name)
	return err
}
