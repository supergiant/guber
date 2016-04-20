package guber

// EventCollection is a Collection interface for Events.
type EventCollection interface {
	Meta() *CollectionMeta
	New() *Event
	Create(e *Event) (*Event, error)
	Query(q *QueryParams) (*EventList, error)
	List() (*EventList, error)
	Get(name string) (*Event, error)
	Update(name string, r *Event) (*Event, error)
	Delete(name string) error
}

// Events implmenets EventCollection.
type Events struct {
	client    *RealClient
	Namespace string
}

// Meta implements the Collection interface.
func (c *Events) Meta() *CollectionMeta {
	return &CollectionMeta{
		DomainName: "",
		APIGroup:   "api",
		APIVersion: "v1",
		APIName:    "events",
		Kind:       "Event",
	}
}

func (c *Events) New() *Event {
	return &Event{
		collection: c,
	}
}

func (c *Events) Create(e *Event) (*Event, error) {
	r := c.New()
	if err := c.client.Post().Collection(c).Namespace(c.Namespace).Entity(e).Do().Into(r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Events) Query(q *QueryParams) (*EventList, error) {
	list := new(EventList)
	if err := c.client.Get().Collection(c).Namespace(c.Namespace).Query(q).Do().Into(list); err != nil {
		return nil, err
	}
	for _, r := range list.Items {
		r.collection = c
	}
	return list, nil
}

func (c *Events) List() (*EventList, error) {
	list := new(EventList)
	if err := c.client.Get().Collection(c).Namespace(c.Namespace).Do().Into(list); err != nil {
		return nil, err
	}
	for _, r := range list.Items {
		r.collection = c
	}
	return list, nil
}

func (c *Events) Get(name string) (*Event, error) {
	r := c.New()
	if err := c.client.Get().Collection(c).Namespace(c.Namespace).Name(name).Do().Into(r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Events) Update(name string, r *Event) (*Event, error) {
	if err := c.client.Patch().Collection(c).Namespace(c.Namespace).Name(name).Entity(r).Do().Into(r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Events) Delete(name string) error {
	req := c.client.Delete().Collection(c).Namespace(c.Namespace).Name(name).Do()
	return req.err
}

// Resource-level

func (r *Event) Reload() (*Event, error) {
	return r.collection.Get(r.Metadata.Name)
}

func (r *Event) Save() error {
	_, err := r.collection.Update(r.Metadata.Name, r)
	return err
}

func (r *Event) Delete() error {
	return r.collection.Delete(r.Metadata.Name)
}
