package cats

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/NYTimes/marvin"

	"google.golang.org/appengine/log"
)

func (s *service) addCat(ctx context.Context, r interface{}) (interface{}, error) {
	// make type conversion to the expected Cat pointer
	cat := r.(*Cat)

	// hit the injected DB layer
	err := s.db.AddCat(ctx, cat)
	if err != nil {
		log.Errorf(ctx, "unable to get count: %s", err)
		return nil, marvin.NewJSONStatusResponse(map[string]string{
			"error": "server error"}, http.StatusInternalServerError)
	}

	// wrap respones so we can return with a 201 code
	return marvin.NewJSONStatusResponse(cat, http.StatusCreated), nil
}

func decodeCat(ctx context.Context, r *http.Request) (interface{}, error) {
	var cat Cat
	err := json.NewDecoder(r.Body).Decode(&cat)
	if err != nil {
		return nil, marvin.NewJSONStatusResponse(map[string]string{
			"error": "bad request"}, http.StatusBadRequest)
	}
	return &cat, nil
}
