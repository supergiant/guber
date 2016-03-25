package guber

import (
	"crypto/tls"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

var (
	tevents = &Events{
		client:    tClient,
		Namespace: "test",
	}
)

func TestDomainName(t *testing.T) {
	resp := tevents.DomainName()
	expected := ""

	if expected != resp {
		t.Error("ERROR .DomainName(): expected,", expected, "-- But got,", resp)
	}
}

func TestApiGroup(t *testing.T) {
	resp := tevents.APIGroup()
	expected := "api"

	if expected != resp {
		t.Error("ERROR .ApiGroup(): expected,", expected, "-- But got,", resp)
	}
}

func TestApiVersion(t *testing.T) {
	resp := tevents.APIVersion()
	expected := "v1"

	if expected != resp {
		t.Error("ERROR .ApiVersion(): expected,", expected, "-- But got,", resp)
	}
}

func TestApiName(t *testing.T) {
	resp := tevents.APIName()
	expected := "events"

	if expected != resp {
		t.Error("ERROR .ApiName(): expected,", expected, "-- But got,", resp)
	}
}

func TestKind(t *testing.T) {
	resp := tevents.Kind()
	expected := "Event"

	if expected != resp {
		t.Error("ERROR .Kind(): expected,", expected, "-- But got,", resp)
	}
}

func TestCreate(t *testing.T) {
	// Setup our test server
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Much Success")
	}))
	defer ts.Close()

	// test client
	url := strings.Replace(ts.URL, "https://", "", -1)
	tsClient := &Client{
		Host:     url,
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

	tsEvents := &Events{
		client:    tsClient,
		Namespace: "test",
	}
	resp, _ := tsEvents.Create(&Event{
		Message: "test",
		Count:   1,
	})

	expected := &Event{
		Message: "test",
		Count:   1,
	}

	if !reflect.DeepEqual(expected, resp) {
		t.Error("ERROR .Create(): expected,", expected, "-- But got,", resp)
	}
}

func TestCreateError(t *testing.T) {
	_, err := tevents.Create(&Event{
		Message: "test",
		Count:   1,
	})

	if err == nil {
		t.Error("ERROR .Create(): event create to fail.. but it did not. ")
	}
}

func TestQuery(t *testing.T) {
	// Setup our test server
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `{
  "kind": "EventList",
  "apiVersion": "v1",
  "metadata": {
    "selfLink": "/api/v1/events",
    "resourceVersion": "test"
  },
  "items": [
	{
            "message": "test"
        }
	]
}`)
	}))
	defer ts.Close()

	// test client
	url := strings.Replace(ts.URL, "https://", "", -1)
	tsClient := &Client{
		Host:     url,
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

	tsEvents := &Events{
		client:    tsClient,
		Namespace: "test",
	}
	resp, _ := tsEvents.Query(&QueryParams{
		LabelSelector: "test",
	})

	if resp.Items[0].Message != "test" {
		t.Error("ERROR .Query(): expected, \"test\"  -- But got,", resp.Items[0].Message)
	}
}

func TestQueryError(t *testing.T) {
	_, err := tevents.Query(&QueryParams{
		LabelSelector: "test",
	})
	if err == nil {
		t.Error("ERROR .Create(): event create to fail.. but it did not. ")
	}
}

func TestList(t *testing.T) {
	// Setup our test server
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `{
  "kind": "EventList",
  "apiVersion": "v1",
  "metadata": {
    "selfLink": "/api/v1/events",
    "resourceVersion": "test"
  },
  "items": [
	{
            "message": "test"
        }
	]
}`)
	}))
	defer ts.Close()

	// test client
	url := strings.Replace(ts.URL, "https://", "", -1)
	tsClient := &Client{
		Host:     url,
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

	tsEvents := &Events{
		client:    tsClient,
		Namespace: "test",
	}
	resp, _ := tsEvents.List()

	if resp.Items[0].Message != "test" {
		t.Error("ERROR .List(): expected, \"test\"  -- But got,", resp.Items[0].Message)
	}
}

func TestListError(t *testing.T) {
	_, err := tevents.List()
	if err == nil {
		t.Error("ERROR .List(): event create to fail.. but it did not. ")
	}
}

func TestEventGet(t *testing.T) {
	// Setup our test server
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `{
            "kind": "Event",
            "apiVersion": "v1",
            "metadata": {
                "name": "test"
            },
            "message": "test"
        }`)
	}))
	defer ts.Close()

	// test client
	url := strings.Replace(ts.URL, "https://", "", -1)
	tsClient := &Client{
		Host:     url,
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

	tsEvents := &Events{
		client:    tsClient,
		Namespace: "test",
	}
	resp, _ := tsEvents.Get("test")

	if resp.Message != "test" {
		t.Error("ERROR .Get(): expected, \"test\"  -- But got,", resp.Message)
	}
}

func TestEventGetNilResult(t *testing.T) {
	// Setup our test server
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "")
	}))
	defer ts.Close()

	// test client
	url := strings.Replace(ts.URL, "https://", "", -1)
	tsClient := &Client{
		Host:     url,
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

	tsEvents := &Events{
		client:    tsClient,
		Namespace: "test",
	}
	resp, _ := tsEvents.Get("test")

	if resp != nil {
		t.Error("ERROR .Get(): expected, Nil  -- But got,", resp)
	}
}
func TestGetError(t *testing.T) {
	_, err := tevents.Get("test")
	if err == nil {
		t.Error("ERROR .Get(): event create to fail.. but it did not. ")
	}
}

func TestEventUpdate(t *testing.T) {
	// Setup our test server
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `{
						"kind": "Event",
						"apiVersion": "v1",
						"metadata": {
								"name": "test"
						},
						"message": "test"
				}`)
	}))
	defer ts.Close()

	// test client
	url := strings.Replace(ts.URL, "https://", "", -1)
	tsClient := &Client{
		Host:     url,
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

	tsEvents := &Events{
		client:    tsClient,
		Namespace: "test",
	}
	resp, _ := tsEvents.Update("test", &Event{
		Message: "test",
		Count:   1,
	})
	if resp.Message != "test" {
		t.Error("ERROR .Update(): expected, \"test\"  -- But got,", resp.Message)
	}
}
func TestUpdateError(t *testing.T) {
	_, err := tevents.Update("test", &Event{
		Message: "test",
		Count:   1,
	})
	if err == nil {
		t.Error("ERROR .Update(): event create to fail.. but it did not. ")
	}
}

func TestEventDelete(t *testing.T) {
	// Setup our test server
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `{
						"kind": "Event",
						"apiVersion": "v1",
						"metadata": {
								"name": "test"
						},
						"message": "test"
				}`)
	}))
	defer ts.Close()

	// test client
	url := strings.Replace(ts.URL, "https://", "", -1)
	tsClient := &Client{
		Host:     url,
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

	tsEvents := &Events{
		client:    tsClient,
		Namespace: "test",
	}
	resp, _ := tsEvents.Delete("test")
	if !resp {
		t.Error("ERROR .Delete(): expected, \"true\"  -- But got,", resp)
	}
}
func TestDeleteError(t *testing.T) {
	_, err := tevents.Delete("test")
	if err == nil {
		t.Error("ERROR .Delete(): event create to fail.. but it did not. ")
	}
}
