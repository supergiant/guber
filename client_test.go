package guber

import (
	"crypto/tls"
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	tClient = &Client{
		Host:     "test",
		Username: "test",
		Password: "test",
		http: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
	}
)

// Tests the NewClient function
func TestCreateNewClient(t *testing.T) {
	// Create a client
	Convey("When creating a new Kubernetes client.", t, func() {
		client := NewClient("test", "test", "test", true)

		Convey("We would expect the resulting client to look like our expected Client object.", func() {
			// Our expected output.
			expected := &Client{
				Host:     "test",
				Username: "test",
				Password: "test",
				http: &http.Client{
					Transport: &http.Transport{
						TLSClientConfig: &tls.Config{
							InsecureSkipVerify: true,
						},
					},
				},
			}
			So(client, ShouldResemble, expected)

		})
	})

}

func TestClientGet(t *testing.T) {
	Convey("When running a .Get() method on a Client object.", t, func() {
		resp := tClient.Get()
		Convey("We expect to get a Request object returned containing the Client object, and the \"GET\" method.", func() {
			expected := &Request{
				client: tClient,
				method: "GET",
			}
			So(resp, ShouldResemble, expected)
		})
	})
}

func TestClientPost(t *testing.T) {
	Convey("When running a .Post() method on a Client object.", t, func() {
		resp := tClient.Post()
		Convey("We expect to get a Request object returned containing the Client object, and the \"POST\" method.", func() {
			expected := &Request{
				client: tClient,
				method: "POST",
			}
			So(resp, ShouldResemble, expected)
		})
	})
}

func TestClientPatch(t *testing.T) {
	Convey("When running a .Patch() method on a Client object.", t, func() {
		resp := tClient.Patch()
		Convey("We expect to get a Request object returned containing the Client object, and the \"PATCH\" method.", func() {
			expected := &Request{
				client: tClient,
				method: "PATCH",
				headers: map[string]string{
					"Content-Type": "application/merge-patch+json",
				},
			}
			So(resp, ShouldResemble, expected)
		})
	})
}

func TestClientDelete(t *testing.T) {
	Convey("When running a .Delete() method on a Client object.", t, func() {
		resp := tClient.Delete()
		Convey("We expect to get a Request object returned containing the Client object, and the \"DELETE\" method.", func() {
			expected := &Request{
				client: tClient,
				method: "DELETE",
			}
			So(resp, ShouldResemble, expected)
		})
	})
}

func TestClientNamespaces(t *testing.T) {
	Convey("When running a .Namespaces() method on a Client object.", t, func() {
		resp := tClient.Namespaces()
		Convey("We expect to get a Namespaces object containing the Client object.", func() {
			expected := &Namespaces{
				client: tClient,
			}
			So(resp, ShouldResemble, expected)
		})
	})
}

func TestClientEvents(t *testing.T) {
	Convey("When running a .Events(\"test\") method on a Client object.", t, func() {
		resp := tClient.Events("test")
		Convey("We expect to get a Events object containing the Client object, and the Namespace string passed to the method.", func() {
			expected := &Events{
				client:    tClient,
				Namespace: "test",
			}
			So(resp, ShouldResemble, expected)
		})
	})
}

func TestClientSecrets(t *testing.T) {
	Convey("When running a .Secrets(\"test\") method on a Client object.", t, func() {
		resp := tClient.Secrets("test")
		Convey("We expect to get a Secrets object containing the Client object, and the Namespace string passed to the method.", func() {
			expected := &Secrets{
				client:    tClient,
				Namespace: "test",
			}
			So(resp, ShouldResemble, expected)
		})
	})
}

func TestClientServices(t *testing.T) {
	Convey("When running a .Services(\"test\") method on a Client object.", t, func() {
		resp := tClient.Services("test")
		Convey("We expect to get a Services object containing the Client object, and the Namespace string passed to the method.", func() {
			expected := &Services{
				client:    tClient,
				Namespace: "test",
			}
			So(resp, ShouldResemble, expected)
		})
	})
}

func TestClientReplicationControllers(t *testing.T) {
	Convey("When running a .ReplicationControllers(\"test\") method on a Client object.", t, func() {
		resp := tClient.ReplicationControllers("test")
		Convey("We expect to get a ReplicationControllers object containing the Client object, and the Namespace string passed to the method.", func() {
			expected := &ReplicationControllers{
				client:    tClient,
				Namespace: "test",
			}
			So(resp, ShouldResemble, expected)
		})
	})
}

func TestClientPods(t *testing.T) {
	Convey("When running a .Pods(\"test\") method on a Client object.", t, func() {
		resp := tClient.Pods("test")
		Convey("We expect to get a Pods object containing the Client object, and the Namespace string passed to the method.", func() {
			expected := &Pods{
				client:    tClient,
				Namespace: "test",
			}
			So(resp, ShouldResemble, expected)
		})
	})
}
