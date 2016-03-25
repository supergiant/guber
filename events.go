package guber

// Events is an object that holds a route of namespace for a client connection.
type Events struct {
	client    *Client
	Namespace string
}

// DomainName returns domain name
func (r Events) DomainName() string {
	return ""
}

// APIGroup returns the kubernetes api group.
func (r Events) APIGroup() string {
	return defaultAPIGroup
}

// APIVersion returns the kubernetes api version.
func (r Events) APIVersion() string {
	return defaultAPIVersion
}

// APIName returns the "events" api name.
func (r Events) APIName() string {
	return "events"
}

// Kind returns the events object kind "Event".
func (r Events) Kind() string {
	return "Event"
}

// Create creates a new kubernetes event.
func (r *Events) Create(e *Event) (*Event, error) {
	if err := r.client.Post().Resource(r).Namespace(r.Namespace).Entity(e).Do().Into(e); err != nil {
		return nil, err
	}
	return e, nil
}

// Query searches kubernetes events and returns in an EventList object.
func (r *Events) Query(q *QueryParams) (*EventList, error) {
	list := new(EventList)
	err := r.client.Get().Resource(r).Namespace(r.Namespace).Query(q).Do().Into(list)
	return list, err
}

// List returns a list of kuebrnetes events in an EventList object.
func (r *Events) List() (*EventList, error) {
	list := new(EventList)
	err := r.client.Get().Resource(r).Namespace(r.Namespace).Do().Into(list)
	return list, err
}

// Get gets a kubernetes object and returns and Event object.
func (r *Events) Get(name string) (*Event, error) {
	e := new(Event)
	req := r.client.Get().Resource(r).Namespace(r.Namespace).Name(name).Do()
	if err := req.Into(e); err != nil {
		return nil, err
	}
	if req.found {
		return e, nil
	}
	return nil, nil
}

// Update updates a event object, and returns the new Event object.
func (r *Events) Update(name string, e *Event) (*Event, error) {
	if err := r.client.Patch().Resource(r).Namespace(r.Namespace).Name(name).Entity(e).Do().Into(e); err != nil {
		return nil, err
	}
	return e, nil
}

// Delete deletes a kuberntes event object.
func (r *Events) Delete(name string) (found bool, err error) {
	req := r.client.Delete().Resource(r).Namespace(r.Namespace).Name(name).Do()
	return req.found, req.err
}
