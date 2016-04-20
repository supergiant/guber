package guber

// ReplicationControllerCollection is a Collection interface for ReplicationControllers.
type ReplicationControllerCollection interface {
	Meta() *CollectionMeta
	New() *ReplicationController
	Create(e *ReplicationController) (*ReplicationController, error)
	Query(q *QueryParams) (*ReplicationControllerList, error)
	List() (*ReplicationControllerList, error)
	Get(name string) (*ReplicationController, error)
	Update(name string, r *ReplicationController) (*ReplicationController, error)
	Delete(name string) error
}

// ReplicationControllers implmenets ReplicationControllerCollection.
type ReplicationControllers struct {
	client    *RealClient
	Namespace string
}

// Meta implements the Collection interface.
func (c *ReplicationControllers) Meta() *CollectionMeta {
	return &CollectionMeta{
		DomainName: "",
		APIGroup:   "api",
		APIVersion: "v1",
		APIName:    "replicationcontrollers",
		Kind:       "ReplicationController",
	}
}

func (c *ReplicationControllers) New() *ReplicationController {
	return &ReplicationController{
		collection: c,
	}
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
	if err := c.client.Get().Collection(c).Namespace(c.Namespace).Name(name).Do().Into(r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *ReplicationControllers) Update(name string, r *ReplicationController) (*ReplicationController, error) {
	if err := c.client.Patch().Collection(c).Namespace(c.Namespace).Name(name).Entity(r).Do().Into(r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *ReplicationControllers) Delete(name string) error {
	req := c.client.Delete().Collection(c).Namespace(c.Namespace).Name(name).Do()
	return req.err
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
	return r.collection.Delete(r.Metadata.Name)
}
