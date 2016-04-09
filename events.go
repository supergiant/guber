package guber

type Events struct {
	client    *Client
	Namespace string
}

func (c *Events) New() *Event {
	return &Event{
		collection: c,
	}
}

func (c Events) DomainName() string {
	return ""
}

func (c Events) APIGroup() string {
	return "api"
}

func (c Events) APIVersion() string {
	return "v1"
}

func (c Events) APIName() string {
	return "events"
}

func (c Events) Kind() string {
	return "Event"
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
	req := c.client.Get().Collection(c).Namespace(c.Namespace).Name(name).Do()
	if err := req.Into(r); err != nil {
		return nil, err
	}
	if req.found {
		return r, nil
	}
	return nil, nil
}

func (c *Events) Update(name string, r *Event) (*Event, error) {
	if err := c.client.Patch().Collection(c).Namespace(c.Namespace).Name(name).Entity(r).Do().Into(r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Events) Delete(name string) (found bool, err error) {
	req := c.client.Delete().Collection(c).Namespace(c.Namespace).Name(name).Do()
	return req.found, req.err
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
	_, err := r.collection.Delete(r.Metadata.Name)
	return err
}
