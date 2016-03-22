package guber

type Events struct {
	client    *Client
	Namespace string
}

func (r Events) DomainName() string {
	return ""
}

func (r Events) ApiGroup() string {
	return "api"
}

func (r Events) ApiVersion() string {
	return "v1"
}

func (r Events) ApiName() string {
	return "events"
}

func (r Events) Kind() string {
	return "Event"
}

func (r *Events) Create(e *Event) (*Event, error) {
	err := r.client.Post().Resource(r).Namespace(r.Namespace).Entity(e).Do().Into(e)
	return e, err
}

func (r *Events) List() (*EventList, error) {
	list := new(EventList)
	err := r.client.Get().Resource(r).Namespace(r.Namespace).Do().Into(list)
	return list, err
}

func (r *Events) Get(name string) (*Event, error) {
	e := new(Event)
	err := r.client.Get().Resource(r).Namespace(r.Namespace).Name(name).Do().Into(e)
	return e, err
}

func (r *Events) Update(name string, e *Event) (*Event, error) {
	err := r.client.Patch().Resource(r).Namespace(r.Namespace).Name(name).Entity(e).Do().Into(e)
	return e, err
}

func (r *Events) Delete(name string) (found bool, err error) {
	req := r.client.Delete().Resource(r).Namespace(r.Namespace).Name(name).Do()
	return req.found, req.err
}
