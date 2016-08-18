package guber

import "net/http"

// Client describes behavior of the root Kubernetes client object.
type Client interface {
	// Namespaces returns a NamespaceCollection.
	Namespaces() NamespaceCollection

	// Events returns a EventCollection.
	Events(namespace string) EventCollection

	// Secrets returns a SecretCollection.
	Secrets(namespace string) SecretCollection

	// Services returns a ServiceCollection.
	Services(namespace string) ServiceCollection

	// ReplicationControllers returns a ReplicationControllerCollection.
	ReplicationControllers(namespace string) ReplicationControllerCollection

	// Pods returns a PodCollection.
	Pods(namespace string) PodCollection

	// Nodes returns a NodeCollection.
	Nodes() NodeCollection
}

var (
	defaultAPIGroup   = "api"
	defaultAPIVersion = "v1"
)

type Entity interface {
}

// CollectionMeta holds info required by all Kubernetes Resources defined.
type CollectionMeta struct {
	DomainName string // empty unless something like ThirdPartyResource
	APIGroup   string // usually "api"
	APIVersion string // usually "v1"
	APIName    string // e.g. "replicationcontrollers"
	Kind       string // e.g. "ReplicationController"
}

// Collection defines an interface for collections of Kubernetes resources.
type Collection interface {
	Meta() *CollectionMeta
}

// RealClient implements Client.
type RealClient struct {
	Host     string
	Username string
	Password string
	http     *http.Client
}

// NewClient creates a new Client.
func NewClient(host string, user string, pass string, httpClient *http.Client) Client {
	return &RealClient{host, user, pass, httpClient}
}

// Get performs a GET request against a Client object.
func (c *RealClient) Get() *Request {
	return &Request{client: c, method: "GET"}
}

// Post performs a POST request against a Client object.
func (c *RealClient) Post() *Request {
	return &Request{client: c, method: "POST"}
}

// Patch performs a PATCH request against a Client object.
func (c *RealClient) Patch() *Request {
	return &Request{
		client: c,
		method: "PATCH",
		headers: map[string]string{
			"Content-Type": "application/merge-patch+json",
		},
	}
}

// Delete performs a DELETE request against a Client object.
func (c *RealClient) Delete() *Request {
	return &Request{client: c, method: "DELETE"}
}

// Namespaces returns a Namespaces object from a Client object.
func (c *RealClient) Namespaces() NamespaceCollection {
	return &Namespaces{c}
}

// Events returns a Events object from a Client object.
func (c *RealClient) Events(namespace string) EventCollection {
	return &Events{c, namespace}
}

// Secrets returns a Secrets object from a Client object.
func (c *RealClient) Secrets(namespace string) SecretCollection {
	return &Secrets{c, namespace}
}

// Services returns a Services object from a Client object.
func (c *RealClient) Services(namespace string) ServiceCollection {
	return &Services{c, namespace}
}

// ReplicationControllers returns a ReplicationControllers object from a Client object.
func (c *RealClient) ReplicationControllers(namespace string) ReplicationControllerCollection {
	return &ReplicationControllers{c, namespace}
}

// Pods returns a Pods object from a Client object.
func (c *RealClient) Pods(namespace string) PodCollection {
	return &Pods{c, namespace}
}

// Namespaces returns a Nodes object from a Client object.
func (c *RealClient) Nodes() NodeCollection {
	return &Nodes{c}
}
