package cats

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/NYTimes/marvin"
	"github.com/golang/protobuf/proto"
	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/user"
)

func TestAddCat(t *testing.T) {
	tests := []struct {
		name string

		givenCat    *Cat
		givenFormat string
		givenDB     DB

		wantCode     int
		wantResponse *Cat
		wantError    *ErrorResponse
	}{
		{
			name: "JSON Success",

			givenCat:    testCat1,
			givenFormat: "json",
			givenDB: &testDB{
				MockAddCat: func(_ context.Context, c *Cat) error {
					return nil
				},
			},

			wantCode:     http.StatusCreated,
			wantResponse: testCat1,
		},
		{
			name: "Protobuf Success",

			givenCat:    testCat1,
			givenFormat: "proto",
			givenDB: &testDB{
				MockAddCat: func(_ context.Context, c *Cat) error {
					return nil
				},
			},

			wantCode:     http.StatusCreated,
			wantResponse: testCat1,
		},
		{
			name: "JSON Error",

			givenCat:    testCat1,
			givenFormat: "json",
			givenDB: &testDB{
				MockAddCat: func(_ context.Context, c *Cat) error {
					return errors.New("whoops!")
				},
			},

			wantCode:  http.StatusInternalServerError,
			wantError: &ErrorResponse{Error: "server error"},
		},
		{
			name: "Protobuf Error",

			givenCat:    testCat1,
			givenFormat: "proto",
			givenDB: &testDB{
				MockAddCat: func(_ context.Context, c *Cat) error {
					return errors.New("whoops!")
				},
			},

			wantCode:  http.StatusInternalServerError,
			wantError: &ErrorResponse{Error: "server error"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			// init server with our injected DB impl
			svr := marvin.NewServer(&service{db: test.givenDB})

			var (
				err  error
				data []byte
			)
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

			// use our aetest instance to create our test request so it has proper
			// context attached
			r, err := testInst.NewRequest(http.MethodPost, "/add."+test.givenFormat,
				bytes.NewBuffer(data))
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
			var gotRes Cat
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
