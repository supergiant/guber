package guber

// NodeCollection is a Collection interface for Nodes.
type NodeCollection interface {
	Meta() *CollectionMeta
	New() *Node
	Create(e *Node) (*Node, error)
	Query(q *QueryParams) (*NodeList, error)
	List() (*NodeList, error)
	Get(name string) (*Node, error)
	Update(name string, r *Node) (*Node, error)
	Delete(name string) error
}

// Nodes implmenets NodeCollection.
type Nodes struct {
	client *RealClient
}

// Meta implements the Collection interface.
func (c *Nodes) Meta() *CollectionMeta {
	return &CollectionMeta{
		DomainName: "",
		APIGroup:   "api",
		APIVersion: "v1",
		APIName:    "nodes",
		Kind:       "Node",
	}
}

func (c *Nodes) New() *Node {
	return &Node{
		collection: c,
	}
}

func (c *Nodes) Create(e *Node) (*Node, error) {
	r := c.New()
	if err := c.client.Post().Collection(c).Entity(e).Do().Into(r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Nodes) Query(q *QueryParams) (*NodeList, error) {
	list := new(NodeList)
	if err := c.client.Get().Collection(c).Query(q).Do().Into(list); err != nil {
		return nil, err
	}
	for _, r := range list.Items {
		r.collection = c
	}
	return list, nil
}

func (c *Nodes) List() (*NodeList, error) {
	list := new(NodeList)
	if err := c.client.Get().Collection(c).Do().Into(list); err != nil {
		return nil, err
	}
	for _, r := range list.Items {
		r.collection = c
	}
	return list, nil
}

func (c *Nodes) Get(name string) (*Node, error) {
	r := c.New()
	if err := c.client.Get().Collection(c).Name(name).Do().Into(r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Nodes) Update(name string, r *Node) (*Node, error) {
	if err := c.client.Patch().Collection(c).Name(name).Entity(r).Do().Into(r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Nodes) Delete(name string) error {
	req := c.client.Delete().Collection(c).Name(name).Do()
	return req.err
}

// Resource-level

func (r *Node) Reload() (*Node, error) {
	return r.collection.Get(r.Metadata.Name)
}

func (r *Node) Save() error {
	_, err := r.collection.Update(r.Metadata.Name, r)
	return err
}

func (r *Node) Delete() error {
	return r.collection.Delete(r.Metadata.Name)
}

func (r *Node) IsOutOfDisk() bool {

	// TODO repeats code in pod.IsReady(), make helper for getting condition

	if len(r.Status.Conditions) == 0 {
		return false
	}

	var condition *NodeStatusCondition
	for _, cond := range r.Status.Conditions {
		if cond.Type == "OutOfDisk" {
			condition = cond
			break
		}
	}
	return condition.Status == "True"
}

func (r *Node) ExternalIP() (ip string) {
	if r.Status == nil {
		return
	}
	for _, addr := range r.Status.Addresses {
		if addr.Type == "ExternalIP" {
			ip = addr.Address
		}
	}
	return
}

func (r *Node) HeapsterStats() (*HeapsterStats, error) {
	path := "api/v1/proxy/namespaces/kube-system/services/heapster/api/v1/model/nodes/" + r.Metadata.Name + "/stats"
	out := new(HeapsterStats)
	err := r.collection.client.Get().Path(path).Do().Into(out)
	return out, err
}
