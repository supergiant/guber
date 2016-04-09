package guber

type Pods struct {
	client    *Client
	Namespace string
}

func (c *Pods) New() *Pod {
	return &Pod{
		collection: c,
	}
}

func (c Pods) DomainName() string {
	return ""
}

func (c Pods) APIGroup() string {
	return "api"
}

func (c Pods) APIVersion() string {
	return "v1"
}

func (c Pods) APIName() string {
	return "pods"
}

func (c Pods) Kind() string {
	return "Pod"
}

func (c *Pods) Create(e *Pod) (*Pod, error) {
	r := c.New()
	if err := c.client.Post().Collection(c).Namespace(c.Namespace).Entity(e).Do().Into(r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Pods) Query(q *QueryParams) (*PodList, error) {
	list := new(PodList)
	if err := c.client.Get().Collection(c).Namespace(c.Namespace).Query(q).Do().Into(list); err != nil {
		return nil, err
	}
	for _, r := range list.Items {
		r.collection = c
	}
	return list, nil
}

func (c *Pods) List() (*PodList, error) {
	list := new(PodList)
	if err := c.client.Get().Collection(c).Namespace(c.Namespace).Do().Into(list); err != nil {
		return nil, err
	}
	for _, r := range list.Items {
		r.collection = c
	}
	return list, nil
}

func (c *Pods) Get(name string) (*Pod, error) {
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

func (c *Pods) Update(name string, r *Pod) (*Pod, error) {
	if err := c.client.Patch().Collection(c).Namespace(c.Namespace).Name(name).Entity(r).Do().Into(r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Pods) Delete(name string) (found bool, err error) {
	req := c.client.Delete().Collection(c).Namespace(c.Namespace).Name(name).Do()
	return req.found, req.err
}

// Resource-level

func (r *Pod) Reload() (*Pod, error) {
	return r.collection.Get(r.Metadata.Name)
}

func (r *Pod) Save() error {
	_, err := r.collection.Update(r.Metadata.Name, r)
	return err
}

func (r *Pod) Delete() error {
	_, err := r.collection.Delete(r.Metadata.Name)
	return err
}

func (r *Pod) Log(container string) (string, error) {
	// TODO we could consolidate all these collection-based methods with one Resource() mtehod
	return r.collection.client.Get().Collection(r.collection).Namespace(r.collection.Namespace).Name(r.Metadata.Name).Path("log?container=" + container).Do().Body()
}

func (r *Pod) IsReady() bool {
	if len(r.Status.Conditions) == 0 {
		return false
	}

	var readyCondition *PodStatusCondition
	for _, cond := range r.Status.Conditions {
		if cond.Type == "Ready" {
			readyCondition = cond
			break
		}
	}
	return readyCondition.Status == "True"
}
