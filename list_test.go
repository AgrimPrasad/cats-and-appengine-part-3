package cats

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/user"

	"github.com/NYTimes/marvin"
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
		},
		{
			name: "Protobuf Success",

			givenFormat: "proto",
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
		},
		{
			name: "JSON Error",

			givenFormat: "json",
			givenDB: &testDB{
				MockGetCats: func(_ context.Context) ([]*Cat, error) {
					return nil, errors.New("aw shucks")
				},
			},

			wantCode:  http.StatusInternalServerError,
			wantError: &ErrorResponse{Error: "unable to get cat list"},
		},
		{
			name: "Protobuf Error",

			givenFormat: "proto",
			givenDB: &testDB{
				MockGetCats: func(_ context.Context) ([]*Cat, error) {
					return nil, errors.New("aw shucks")
				},
			},

			wantCode:  http.StatusInternalServerError,
			wantError: &ErrorResponse{Error: "unable to get cat list"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			// init server with our injected DB impl
			svr := marvin.NewServer(&service{db: test.givenDB})

			// use our aetest instance to create our test request so it has proper
			// context attached
			r, err := testInst.NewRequest(http.MethodGet, "/list."+test.givenFormat, &bytes.Buffer{})
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
