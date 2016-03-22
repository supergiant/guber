package guber

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

func (r *ReplicationControllers) Create(e *ReplicationController) (*ReplicationController, error) {
	err := r.client.Post().Resource(r).Namespace(r.Namespace).Entity(e).Do().Into(e)
	return e, err
}

func (r *ReplicationControllers) List() (*ReplicationControllerList, error) {
	list := new(ReplicationControllerList)
	err := r.client.Get().Resource(r).Namespace(r.Namespace).Do().Into(list)
	return list, err
}

// TODO ideally we should return nil instead of e on error
func (r *ReplicationControllers) Get(name string) (*ReplicationController, error) {
	e := new(ReplicationController)
	err := r.client.Get().Resource(r).Namespace(r.Namespace).Name(name).Do().Into(e)
	return e, err
}

func (r *ReplicationControllers) Update(name string, e *ReplicationController) (*ReplicationController, error) {
	err := r.client.Patch().Resource(r).Namespace(r.Namespace).Name(name).Entity(e).Do().Into(e)
	return e, err
}

func (r *ReplicationControllers) Delete(name string) (found bool, err error) {
	req := r.client.Delete().Resource(r).Namespace(r.Namespace).Name(name).Do()
	return req.found, req.err
}
