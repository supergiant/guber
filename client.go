package guber

import (
	"crypto/tls"
	"net/http"
)

// Interfaces

// Entity is a data holder, usually used for holding json data objects
type Entity interface {
}

// Resource holds interfaces to kubernetes api artifacts.
type Resource interface {
	DomainName() string // empty unless something like ThirdPartyResource
	ApiGroup() string   // usually "api"
	ApiVersion() string // usually "v1"
	ApiName() string    // e.g. "replicationcontrollers"
	Kind() string       // e.g. "ReplicationController"
}

// structs

// Client is a kubernetes client object.
type Client struct {
	Host     string
	Username string
	Password string
	http     *http.Client
}

//Functions

// NewClient creates a new kubernetes client object.
func NewClient(host string, user string, pass string) *Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &Client{host, user, pass, &http.Client{Transport: tr}}
}

//Methods

// Get performs a GET request against a Client object.
func (c *Client) Get() *Request {
	return &Request{client: c, method: "GET"}
}

// Post performs a POST request against a Client object.
func (c *Client) Post() *Request {
	return &Request{client: c, method: "POST"}
}

// Patch performs a PATCH request against a Client object.
func (c *Client) Patch() *Request {
	return &Request{client: c, method: "PATCH"}
}

// Delete performs a DELETE request against a Client object.
func (c *Client) Delete() *Request {
	return &Request{client: c, method: "DELETE"}
}

// Namespaces returns a Namespaces object from a Client object.
func (c *Client) Namespaces() *Namespaces {
	return &Namespaces{c}
}

// Events returns a Events object from a Client object.
func (c *Client) Events(namespace string) *Events {
	return &Events{c, namespace}
}

// Secrets returns a Secrets object from a Client object.
func (c *Client) Secrets(namespace string) *Secrets {
	return &Secrets{c, namespace}
}

// Services returns a Services object from a Client object.
func (c *Client) Services(namespace string) *Services {
	return &Services{c, namespace}
}

// ReplicationControllers returns a ReplicationControllers object from a Client object.
func (c *Client) ReplicationControllers(namespace string) *ReplicationControllers {
	return &ReplicationControllers{c, namespace}
}

// Pods returns a Pods object from a Client object.
func (c *Client) Pods(namespace string) *Pods {
	return &Pods{c, namespace}
}
