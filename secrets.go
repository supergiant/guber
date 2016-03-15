package guber

import "guber/model"

type Secrets struct {
	client    *Client
	Namespace string
}

func (r Secrets) DomainName() string {
	return ""
}

func (r Secrets) ApiGroup() string {
	return "api"
}

func (r Secrets) ApiVersion() string {
	return "v1"
}

func (r Secrets) ApiName() string {
	return "secrets"
}

func (r Secrets) Kind() string {
	return "Secret"
}

func (r Secrets) Create(e Entity) (Entity, error) {
	err := r.client.Post().Resource(r).Namespace(r.Namespace).Entity(e).Do().Into(e)
	return e, err
}

func (r Secrets) List() (EntityList, error) {
	list := new(model.EventList)
	err := r.client.Get().Resource(r).Namespace(r.Namespace).Do().Into(list)
	return list, err
}

func (r Secrets) Get(name string) (Entity, error) {
	e := new(model.Event)
	err := r.client.Get().Resource(r).Namespace(r.Namespace).Name(name).Do().Into(e)
	return e, err
}

func (r Secrets) Update(name string, e Entity) (Entity, error) {
	err := r.client.Patch().Resource(r).Namespace(r.Namespace).Name(name).Entity(e).Do().Into(e)
	return e, err
}

func (r Secrets) Delete(name string) error {
	return r.client.Delete().Resource(r).Namespace(r.Namespace).Name(name).Do().Into(nil)
}
