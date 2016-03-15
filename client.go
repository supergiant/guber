package guber

import (
	"crypto/tls"
	"net/http"
)

type Client struct {
	Host     string
	Username string
	Password string
	http     *http.Client
}

func NewClient(host string, user string, pass string) *Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &Client{host, user, pass, &http.Client{Transport: tr}}
}

type Entity interface {
}

type EntityList interface {
}

type Resource interface {
	DomainName() string // empty unless something like ThirdPartyResource
	ApiGroup() string   // usually "api"
	ApiVersion() string // usually "v1"
	ApiName() string    // e.g. "replicationcontrollers"
	Kind() string       // e.g. "ReplicationController"

	Create(Entity) (Entity, error)
	List() (EntityList, error)
	Get(string) (Entity, error)
	Update(string, Entity) (Entity, error)
	Delete(string) error
}

func (c *Client) Get() *Request {
	return &Request{client: c, method: "GET"}
}

func (c *Client) Post() *Request {
	return &Request{client: c, method: "POST"}
}

func (c *Client) Patch() *Request {
	return &Request{client: c, method: "PATCH"}
}

func (c *Client) Delete() *Request {
	return &Request{client: c, method: "DELETE"}
}

func (c *Client) Namespaces() *Namespaces {
	return &Namespaces{c}
}

func (c *Client) Events(namespace string) *Events {
	return &Events{c, namespace}
}

func (c *Client) Secrets(namespace string) *Secrets {
	return &Secrets{c, namespace}
}

func (c *Client) Services(namespace string) *Services {
	return &Services{c, namespace}
}

func (c *Client) ReplicationControllers(namespace string) *ReplicationControllers {
	return &ReplicationControllers{c, namespace}
}
