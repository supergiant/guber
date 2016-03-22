package guber

type Pods struct {
	client    *Client
	Namespace string
}

func (r Pods) DomainName() string {
	return ""
}

func (r Pods) ApiGroup() string {
	return "api"
}

func (r Pods) ApiVersion() string {
	return "v1"
}

func (r Pods) ApiName() string {
	return "pods"
}

func (r Pods) Kind() string {
	return "Pod"
}

func (r *Pods) Create(e *Pod) (*Pod, error) {
	if err := r.client.Post().Resource(r).Namespace(r.Namespace).Entity(e).Do().Into(e); err != nil {
		return nil, err
	}
	return e, nil
}

func (r *Pods) List(q *QueryParams) (*PodList, error) {
	list := new(PodList)
	err := r.client.Get().Resource(r).Namespace(r.Namespace).Query(q).Do().Into(list)
	return list, err
}

func (r *Pods) Get(name string) (*Pod, error) {
	e := new(Pod)
	req := r.client.Get().Resource(r).Namespace(r.Namespace).Name(name).Do()
	if err := req.Into(e); err != nil {
		return nil, err
	}
	if req.found {
		return e, nil
	}
	return nil, nil
}

func (r *Pods) Update(name string, e *Pod) (*Pod, error) {
	if err := r.client.Patch().Resource(r).Namespace(r.Namespace).Name(name).Entity(e).Do().Into(e); err != nil {
		return nil, err
	}
	return e, nil
}

func (r *Pods) Delete(name string) (found bool, err error) {
	req := r.client.Delete().Resource(r).Namespace(r.Namespace).Name(name).Do()
	return req.found, req.err
}
