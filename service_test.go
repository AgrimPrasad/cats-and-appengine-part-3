package cats

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/user"

	"github.com/NYTimes/marvin"
	"github.com/golang/protobuf/proto"
	"github.com/kr/pretty"
)

var testInst aetest.Instance

// TestMain will let us start the appengine instance in the background, one time.
func TestMain(t *testing.M) {
	var err error
	testInst, err = aetest.NewInstance(nil)
	if err != nil {
		log.Print("unable to start GAE instance: ", err)
		return
	}
	defer testInst.Close()

	t.Run()
}

// TestService is an integration test that hits the datastore emulator
func TestService(t *testing.T) {
	tests := []struct {
		name string

		givenCat    *Cat
		givenFormat string
		givenPath   string
		givenMethod string

		wantCode     int
		wantResponse *CatsResponse
		wantCat      *Cat
		wantError    *ErrorResponse
	}{
		{
			name: "Add JSON cat Success",

			givenFormat: "json",
			givenPath:   "add",
			givenMethod: http.MethodPost,
			givenCat:    testCat1,

			wantCode: http.StatusCreated,
			wantCat:  testCat1ID,
		},
		{
			name: "Add proto cat Success",

			givenFormat: "proto",
			givenPath:   "add",
			givenMethod: http.MethodPost,
			givenCat:    testCat2,

			wantCode: http.StatusCreated,
			wantCat:  testCat2ID,
		},
		{
			name: "List cats",

			givenPath:   "list",
			givenMethod: http.MethodGet,
			givenFormat: "json",

			wantCode: http.StatusOK,
			wantResponse: &CatsResponse{
				Total: 2,
				Cats:  testCatSliceID,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			// init server with our injected DB impl
			svr := marvin.NewServer(&service{db: NewDB()})

			var (
				err  error
				data []byte
			)
			if test.givenCat != nil {
				// marshal the payload into the appropriate format
				switch test.givenFormat {
				case "proto":
					data, err = proto.Marshal(test.givenCat)
				case "json":
					data, err = json.Marshal(test.givenCat)
				}
				if err != nil {
					t.Fatal("unable to marshal cat payload: ", err)
				}
			}

			// use our aetest instance to create our test request so it has proper
			// context attached
			r, err := testInst.NewRequest(test.givenMethod, "/"+test.givenPath+"."+
				test.givenFormat, bytes.NewBuffer(data))
			if err != nil {
				t.Fatal("unable to create GAE requeest: ", err)
			}
			// make it look like someone is logged in
			aetest.Login(&user.User{
				Email:             "jp@nytimes.com",
				FederatedIdentity: "jp"}, r)
			w := httptest.NewRecorder()

			// run the test
			svr.ServeHTTP(w, r)
			got := w.Result()

			// check the response code
			if got.StatusCode != test.wantCode {
				t.Errorf("expected response code of %d, got %d",
					test.wantCode, got.StatusCode)
			}

			// check the response body
			var (
				gotRes CatsResponse
				gotCat Cat
				gotErr ErrorResponse
			)
			switch test.givenFormat {
			case "proto":
				switch {
				case test.wantResponse != nil:
					readProto(t, got, &gotRes)
					compareResults(t, &gotRes, test.wantResponse)
				case test.wantCat != nil:
					readProto(t, got, &gotCat)
					compareResults(t, &gotCat, test.wantCat)
				default:
					readProto(t, got, &gotErr)
					compareResults(t, &gotErr, test.wantError)
				}
			case "json":
				switch {
				case test.wantResponse != nil:
					readJSON(t, got, &gotRes)
					compareResults(t, &gotRes, test.wantResponse)
				case test.wantCat != nil:
					readJSON(t, got, &gotCat)
					compareResults(t, &gotCat, test.wantCat)
				default:
					readJSON(t, got, &gotErr)
					compareResults(t, &gotErr, test.wantError)
				}
			}

			// the datastore emulator is...slow. We need to give it a moment after each
			// write before we attempt to query it.
			time.Sleep(500 * time.Millisecond)
		})
	}
}

var (
	testCat1       = &Cat{Name: "Gus", Breed: "Tabby", Weight: 8.5}
	testCat1ID     = &Cat{Key: 1, Name: "Gus", Breed: "Tabby", Weight: 8.5}
	testCat2       = &Cat{Name: "Ziggy", Breed: "Fat", Weight: 14.5}
	testCat2ID     = &Cat{Key: 2, Name: "Ziggy", Breed: "Fat", Weight: 14.5}
	testCatSlice   = []*Cat{testCat1, testCat2}
	testCatSliceID = []*Cat{testCat1ID, testCat2ID}
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
