package guber

import "guber/model"

type Services struct {
	client    *Client
	Namespace string
}

func (r Services) DomainName() string {
	return ""
}

func (r Services) ApiGroup() string {
	return "api"
}

func (r Services) ApiVersion() string {
	return "v1"
}

func (r Services) ApiName() string {
	return "services"
}

func (r Services) Kind() string {
	return "Service"
}

func (r Services) Create(e Entity) (Entity, error) {
	err := r.client.Post().Resource(r).Namespace(r.Namespace).Entity(e).Do().Into(e)
	return e, err
}

func (r Services) List() (EntityList, error) {
	list := new(model.EventList)
	err := r.client.Get().Resource(r).Namespace(r.Namespace).Do().Into(list)
	return list, err
}

func (r Services) Get(name string) (Entity, error) {
	e := new(model.Event)
	err := r.client.Get().Resource(r).Namespace(r.Namespace).Name(name).Do().Into(e)
	return e, err
}

func (r Services) Update(name string, e Entity) (Entity, error) {
	err := r.client.Patch().Resource(r).Namespace(r.Namespace).Name(name).Entity(e).Do().Into(e)
	return e, err
}

func (r Services) Delete(name string) error {
	return r.client.Delete().Resource(r).Namespace(r.Namespace).Name(name).Do().Into(nil)
}
