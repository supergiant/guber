package guber

import "guber/model"

type Events struct {
	client *Client
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

func (r Events) Create(e Entity) (Entity, error) {
	err := r.client.Post().Resource(r).Entity(e).Do().Into(e)
	return e, err
}

func (r Events) List() (EntityList, error) {
	list := new(model.EventList)
	err := r.client.Get().Resource(r).Do().Into(list)
	return list, err
}

func (r Events) Get(name string) (Entity, error) {
	e := new(model.Event)
	err := r.client.Get().Resource(r).Name(name).Do().Into(e)
	return e, err
}

func (r Events) Update(name string, e Entity) (Entity, error) {
	err := r.client.Patch().Resource(r).Name(name).Entity(e).Do().Into(e)
	return e, err
}

func (r Events) Delete(name string) error {
	return r.client.Delete().Resource(r).Name(name).Do().Into(nil)
}
