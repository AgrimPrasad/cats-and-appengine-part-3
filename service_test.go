package cats

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"testing"

	"google.golang.org/appengine/aetest"

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
