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

func TestEventsCreate(t *testing.T) {
	Convey("We start a mock api server.", t, func() {
		ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO: Should we mock a full response here, and should a kubernetes lib be aware of a malformed response?
			// Right now we seem to accept any successful status code.
			io.WriteString(w, "Much Success")
		}))
		defer ts.Close()
		So(ts, ShouldNotBeNil)
		Convey("When calling the .Create(&Event) method on an Events object.", func() {
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
	Convey("When calling the .Create(&Event) method on a Client object. And there is a problem. We should get an error.", t, func() {
		_, err := tevents.Create(&Event{
			Message: "test",
			Count:   1,
		})
		So(err, ShouldNotBeNil)
	})
	tClient.Host = "test"
}

func TestEventQuery(t *testing.T) {
	Convey("We start a mock api server.", t, func() {
		ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, `{
   "kind":"EventList",
   "apiVersion":"v1",
   "metadata":{
      "selfLink":"/api/v1/events",
      "resourceVersion":"test"
   },
   "items":[
      {
         "message":"test"
      }
   ]
}`)
		}))
		defer ts.Close()
		So(ts, ShouldNotBeNil)
		Convey("When calling the .Query(&QueryParams) method on an Events object that exists.", func() {
			tClient.Host = strings.Replace(ts.URL, "https://", "", -1)
			resp, _ := tevents.Query(&QueryParams{
				LabelSelector: "test",
			})
			Convey("The response object should be a EventsList containing the messge value \"test\" from our server response.", func() {
				So(resp.Items[0].Message, ShouldEqual, "test")
			})
		})
	})
	Convey("When calling the .Query(&QueryParams) method on an Events object. And there is an error", t, func() {
		_, err := tevents.Query(&QueryParams{
			LabelSelector: "Cheese",
		})
		Convey("The response should be nil, and there should be an error.", func() {
			So(err, ShouldNotBeNil)
		})
	})
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
