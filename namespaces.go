package guber

import "guber/model"

type Namespaces struct {
	client *Client
}

func (r Namespaces) DomainName() string {
	return ""
}

func (r Namespaces) ApiGroup() string {
	return "api"
}

func (r Namespaces) ApiVersion() string {
	return "v1"
}

func (r Namespaces) ApiName() string {
	return "namespaces"
}

func (r Namespaces) Kind() string {
	return "Namespace"
}

func (r Namespaces) Create(e Entity) (Entity, error) {
	err := r.client.Post().Resource(r).Entity(e).Do().Into(e)
	return e, err
}

func (r Namespaces) List() (EntityList, error) {
	list := new(model.EventList)
	err := r.client.Get().Resource(r).Do().Into(list)
	return list, err
}

func (r Namespaces) Get(name string) (Entity, error) {
	e := new(model.Event)
	err := r.client.Get().Resource(r).Name(name).Do().Into(e)
	return e, err
}

func (r Namespaces) Update(name string, e Entity) (Entity, error) {
	err := r.client.Patch().Resource(r).Name(name).Entity(e).Do().Into(e)
	return e, err
}

func (r Namespaces) Delete(name string) error {
	return r.client.Delete().Resource(r).Name(name).Do().Into(nil)
}
