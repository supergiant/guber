package guber

import (
	"crypto/tls"
	"net/http"
	"reflect"
	"testing"
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
func TestNewClient(t *testing.T) {
	// Create a client
	client := NewClient("test", "test", "test")

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

	if !reflect.DeepEqual(expected, client) {
		t.Error("ERROR NewClient(): Expected,", expected, "-- But got, ", client)
	}

}

func TestGet(t *testing.T) {
	expected := &Request{
		client: tClient,
		method: "GET",
	}

	resp := tClient.Get()

	if !reflect.DeepEqual(expected, resp) {
		t.Error("ERROR: .Get(): expected,", expected, "-- But got,", resp)
	}
}

func TestPost(t *testing.T) {
	expected := &Request{
		client: tClient,
		method: "POST",
	}

	resp := tClient.Post()

	if !reflect.DeepEqual(expected, resp) {
		t.Error("ERROR: .Post(): expected,", expected, "-- But got,", resp)
	}
}

func TestPatch(t *testing.T) {
	expected := &Request{
		client: tClient,
		method: "PATCH",
	}

	resp := tClient.Patch()

	if !reflect.DeepEqual(expected, resp) {
		t.Error("ERROR: .Patch(): expected,", expected, "-- But got,", resp)
	}
}

func TestDelete(t *testing.T) {
	expected := &Request{
		client: tClient,
		method: "DELETE",
	}

	resp := tClient.Delete()

	if !reflect.DeepEqual(expected, resp) {
		t.Error("ERROR: .Delete(): expected,", expected, "-- But got,", resp)
	}
}

func TestNamespaces(t *testing.T) {
	expected := &Namespaces{
		client: tClient,
	}

	resp := tClient.Namespaces()

	if !reflect.DeepEqual(expected, resp) {
		t.Error("ERROR: .Namespaces(): expected,", expected, "-- But got,", resp)
	}
}

func TestEvents(t *testing.T) {
	expected := &Events{
		client:    tClient,
		Namespace: "test",
	}

	resp := tClient.Events("test")

	if !reflect.DeepEqual(expected, resp) {
		t.Error("ERROR: .Events(): expected,", expected, "-- But got,", resp)
	}
}

func TestSecrets(t *testing.T) {
	expected := &Secrets{
		client:    tClient,
		Namespace: "test",
	}

	resp := tClient.Secrets("test")

	if !reflect.DeepEqual(expected, resp) {
		t.Error("ERROR: .Secrets(): expected,", expected, "-- But got,", resp)
	}
}

func TestServices(t *testing.T) {
	expected := &Services{
		client:    tClient,
		Namespace: "test",
	}

	resp := tClient.Services("test")

	if !reflect.DeepEqual(expected, resp) {
		t.Error("ERROR: .Services(): expected,", expected, "-- But got,", resp)
	}
}

func TestReplicationControllers(t *testing.T) {
	expected := &ReplicationControllers{
		client:    tClient,
		Namespace: "test",
	}

	resp := tClient.ReplicationControllers("test")

	if !reflect.DeepEqual(expected, resp) {
		t.Error("ERROR: .ReplicationControllers(): expected,", expected, "-- But got,", resp)
	}
}

func TestPods(t *testing.T) {
	expected := &Pods{
		client:    tClient,
		Namespace: "test",
	}

	resp := tClient.Pods("test")

	if !reflect.DeepEqual(expected, resp) {
		t.Error("ERROR: .Pods(): expected,", expected, "-- But got,", resp)
	}
}
