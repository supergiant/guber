package guber

import (
	"crypto/tls"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	tevents = &Events{
		client:    tClient,
		Namespace: "test",
	}
)

func TestEventsDomainName(t *testing.T) {
	Convey("When calling the .DomainName() on an Events object.", t, func() {
		resp := tevents.DomainName()
		Convey("We should expect to get a blank string return.", func() {
			expected := ""
			So(resp, ShouldEqual, expected)
		})
	})
}

func TestEventsAPIGroup(t *testing.T) {
	Convey("When calling the .APIGroup() on an Events object.", t, func() {
		resp := tevents.APIGroup()
		Convey("We should expect to get a string \"api\" return.", func() {
			expected := "api"
			So(resp, ShouldEqual, expected)
		})
	})
}

func TestEventsAPIVersion(t *testing.T) {
	Convey("When calling the .APIVersion() on an Events object.", t, func() {
		resp := tevents.APIVersion()
		Convey("We should expect to get a string \"v1\" return.", func() {
			expected := "v1"
			So(resp, ShouldEqual, expected)
		})
	})
}

func TestEventsAPIName(t *testing.T) {
	Convey("When calling the .APIName() on an Events object.", t, func() {
		resp := tevents.APIName()
		Convey("We should expect to get a string \"events\" return.", func() {
			expected := "events"
			So(resp, ShouldEqual, expected)
		})
	})
}

func TestEventsKind(t *testing.T) {
	Convey("When calling the .Kind() on an Events object.", t, func() {
		resp := tevents.Kind()
		Convey("We should expect to get a string \"Event\" return.", func() {
			expected := "Event"
			So(resp, ShouldEqual, expected)
		})
	})
}

func TestEventsCreate(t *testing.T) {
	Convey("We start a mock api server.", t, func() {
		ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, "Much Success")
		}))
		defer ts.Close()
		So(ts, ShouldNotBeNil)
		Convey("When calling the .Create(&Event) method on a Client object.", func() {
			tClient.Host = strings.Replace(ts.URL, "https://", "", -1)
			resp, _ := tevents.Create(&Event{
				Message: "test",
				Count:   1,
			})
			Convey("We expect to get a Event object that matches our expected Event object.", func() {
				expected := &Event{
					Message: "test",
					Count:   1,
				}
				So(resp, ShouldResemble, expected)
			})
		})
	})
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
