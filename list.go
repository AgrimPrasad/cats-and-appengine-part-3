package cats

import (
	"context"
	"net/http"

	"github.com/NYTimes/marvin"

	"google.golang.org/appengine/log"
)

func (s *service) listCats(ctx context.Context, _ interface{}) (interface{}, error) {

	// get cats from injected DB
	cats, err := s.db.GetCats(ctx)
	if err != nil {
		log.Errorf(ctx, "unable to increment counter: %s", err)
		return nil, marvin.NewJSONStatusResponse(map[string]string{
			"error": "unable to increment counter"}, http.StatusInternalServerError)
	}

	return &CatsResponse{
		Total: len(cats),
		Cats:  cats,
	}, nil
}