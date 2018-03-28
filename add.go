package cats

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/NYTimes/marvin"
	"github.com/golang/protobuf/proto"

	"google.golang.org/appengine/log"
)

func (s *service) addCat(ctx context.Context, r interface{}) (interface{}, error) {
	// make type conversion to the expected Cat pointer
	req := r.(*PostAddFormatRequest)

	// hit the injected DB layer
	err := s.db.AddCat(ctx, req.Cat)
	if err != nil {
		log.Errorf(ctx, "unable to get count: %s", err)
		return nil, marvin.NewProtoStatusResponse(&ErrorResponse{
			Error: "server error"}, http.StatusInternalServerError)
	}

	// wrap respones so we can return with a 201 code
	return marvin.NewProtoStatusResponse(req.Cat, http.StatusCreated), nil
}

func decodeCat(ctx context.Context, r *http.Request) (interface{}, error) {
	var cat Cat
	err := json.NewDecoder(r.Body).Decode(&cat)
	if err != nil {
		return nil, marvin.NewJSONStatusResponse(&ErrorResponse{
			Error: "bad request"}, http.StatusBadRequest)
	}
	return &PostAddFormatRequest{Cat: &cat}, nil
}

func decodeCatProto(ctx context.Context, r *http.Request) (interface{}, error) {
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, marvin.NewProtoStatusResponse(&ErrorResponse{
			Error: "unable to read request"}, http.StatusBadRequest)
	}

	var cat Cat
	err = proto.Unmarshal(d, &cat)
	if err != nil {
		return nil, marvin.NewProtoStatusResponse(&ErrorResponse{
			Error: "bad request"}, http.StatusBadRequest)
	}
	return &PostAddFormatRequest{Cat: &cat}, nil
}
