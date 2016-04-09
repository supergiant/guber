package guber

type Nodes struct {
	client *Client
}

func (c *Nodes) New() *Node {
	return &Node{
		collection: c,
	}
}

func (c Nodes) DomainName() string {
	return ""
}

func (c Nodes) APIGroup() string {
	return "api"
}

func (c Nodes) APIVersion() string {
	return "v1"
}

func (c Nodes) APIName() string {
	return "nodes"
}

func (c Nodes) Kind() string {
	return "Node"
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
	req := c.client.Get().Collection(c).Name(name).Do()
	if err := req.Into(r); err != nil {
		return nil, err
	}
	if req.found {
		return r, nil
	}
	return nil, nil
}

func (c *Nodes) Update(name string, r *Node) (*Node, error) {
	if err := c.client.Patch().Collection(c).Name(name).Entity(r).Do().Into(r); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Nodes) Delete(name string) (found bool, err error) {
	req := c.client.Delete().Collection(c).Name(name).Do()
	return req.found, req.err
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
	_, err := r.collection.Delete(r.Metadata.Name)
	return err
}
