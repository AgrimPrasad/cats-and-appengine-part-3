package cats

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/user"

	"github.com/NYTimes/marvin"
	"github.com/NYTimes/marvin/marvintest"
	"github.com/golang/protobuf/proto"
	"github.com/kr/pretty"
)

func TestListCats(t *testing.T) {
	tests := []struct {
		name string

		givenFormat string
		givenDB     DB

		wantCode     int
		wantResponse *CatsResponse
		wantError    *ErrorResponse
	}{
		{
			name: "JSON Success",

			givenFormat: "json",
			givenDB: &testDB{
				MockGetCats: func(_ context.Context) ([]*Cat, error) {
					return testCatSlice, nil
				},
			},

			wantCode: http.StatusOK,
			wantResponse: &CatsResponse{
				Total: 2,
				Cats:  testCatSlice,
			},
			wantError: nil,
		},
	}

	// we need to init the App Engine server to deal with any logs/GAE interaction
	done := marvintest.SetupTestContext(t)
	defer done()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			// init server with our injected DB impl
			svr := marvin.NewServer(&service{db: test.givenDB})

			// setup the request and response capture
			//			r := httptest.NewRequest(http.MethodGet, "/list."+test.givenFormat, nil)
			r, _ := http.NewRequest(http.MethodGet, "/list."+test.givenFormat, &bytes.Buffer{})
			marvintest.SetServerContext(ctx)
			// make it look like someone is logged in
			aetest.Login(&user.User{
				Email:             "jp@nytimes.com",
				FederatedIdentity: "jp"}, r)
			w := httptest.NewRecorder()

			// run the test
			svr.ServeHTTP(w, r)
			got := w.Result()

			// check resp code
			if got.StatusCode != test.wantCode {
				t.Errorf("expected response code of %d, got %d",
					test.wantCode, got.StatusCode)
			}

			var gotRes CatsResponse
			var gotErr ErrorResponse
			switch test.givenFormat {
			case "proto":
				if test.wantResponse != nil {
					readProto(t, got, &gotRes)
					compareResults(t, &gotRes, test.wantResponse)
				} else {
					readProto(t, got, &gotErr)
					compareResults(t, &gotErr, test.wantError)
				}
			case "json":
				if test.wantResponse != nil {
					readJSON(t, got, &gotRes)
					compareResults(t, &gotRes, test.wantResponse)
				} else {
					readJSON(t, got, &gotErr)
					compareResults(t, &gotErr, test.wantError)
				}
			}
		})
	}
}

var (
	testCat1     = &Cat{Name: "Gus", Breed: "Tabby", Weight: 8.5}
	testCat2     = &Cat{Name: "Ziggy", Breed: "Fat", Weight: 14.5}
	testCatSlice = []*Cat{testCat1, testCat2}
)

func readJSON(t *testing.T, res *http.Response, val interface{}) {
	err := json.NewDecoder(res.Body).Decode(val)
	if err != nil {
		t.Errorf("unable to parse response: %s", err)
	}
	defer res.Body.Close()
}

func readProto(t *testing.T, res *http.Response, val proto.Message) {
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("unable to read response: %s", err)
		return
	}
	defer res.Body.Close()

	err = proto.Unmarshal(b, val)
	if err != nil {
		t.Errorf("unable to proto unmarshal response: %s", err)
		return
	}
}

func compareResults(t *testing.T, got, want interface{}) {
	if reflect.DeepEqual(got, want) {
		return
	}
	diffs := pretty.Diff(got, want)
	t.Errorf("found %d difference(s) in actual vs. expected:", len(diffs))
	for _, diff := range diffs {
		t.Log(diff)
	}
}
