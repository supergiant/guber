package guber

// PodCollection is a Collection interface for Pods.
type PodCollection interface {
	Meta() *CollectionMeta
	New() *Pod
	Create(e *Pod) (*Pod, error)
	Query(q *QueryParams) (*PodList, error)
	List() (*PodList, error)
	Get(name string) (*Pod, error)
	Update(name string, r *Pod) (*Pod, error)
	Delete(name string) error
}

// Pods implmenets PodCollection.
type Pods struct {
	client    *RealClient
	Namespace string
}

// Meta implements the Collection interface.
func (c *Pods) Meta() *CollectionMeta {
	return &CollectionMeta{
		DomainName: "",
		APIGroup:   "api",
		APIVersion: "v1",
		APIName:    "pods",
		Kind:       "Pod",
	}
}

func (c *Pods) New() *Pod {
	return &Pod{
		collection: c,
	}
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
	if err := c.client.Get().Collection(c).Namespace(c.Namespace).Name(name).Do().Into(r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Pods) Update(name string, r *Pod) (*Pod, error) {
	if err := c.client.Patch().Collection(c).Namespace(c.Namespace).Name(name).Entity(r).Do().Into(r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Pods) Delete(name string) error {
	req := c.client.Delete().Collection(c).Namespace(c.Namespace).Name(name).Do()
	return req.err
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
	return r.collection.Delete(r.Metadata.Name)
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

func (r *Pod) HeapsterStats() (*HeapsterStats, error) {
	path := "api/v1/proxy/namespaces/kube-system/services/heapster/api/v1/model/namespaces/" + r.Metadata.Namespace + "/pods/" + r.Metadata.Name + "/stats"
	out := new(HeapsterStats)
	err := r.collection.client.Get().Path(path).Do().Into(out)
	return out, err
}
