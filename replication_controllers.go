package guber

import "guber/model"

type ReplicationControllers struct {
	client    *Client
	Namespace string
}

func (r ReplicationControllers) DomainName() string {
	return ""
}

func (r ReplicationControllers) ApiGroup() string {
	return "api"
}

func (r ReplicationControllers) ApiVersion() string {
	return "v1"
}

func (r ReplicationControllers) ApiName() string {
	return "replicationcontrollers"
}

func (r ReplicationControllers) Kind() string {
	return "ReplicationController"
}

func (r ReplicationControllers) Create(e Entity) (Entity, error) {
	err := r.client.Post().Resource(r).Namespace(r.Namespace).Entity(e).Do().Into(e)
	return e, err
}

func (r ReplicationControllers) List() (EntityList, error) {
	list := new(model.EventList)
	err := r.client.Get().Resource(r).Namespace(r.Namespace).Do().Into(list)
	return list, err
}

func (r ReplicationControllers) Get(name string) (Entity, error) {
	e := new(model.Event)
	err := r.client.Get().Resource(r).Namespace(r.Namespace).Name(name).Do().Into(e)
	return e, err
}

func (r ReplicationControllers) Update(name string, e Entity) (Entity, error) {
	err := r.client.Patch().Resource(r).Namespace(r.Namespace).Name(name).Entity(e).Do().Into(e)
	return e, err
}

func (r ReplicationControllers) Delete(name string) error {
	return r.client.Delete().Resource(r).Namespace(r.Namespace).Name(name).Do().Into(nil)
}
