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
		log.Errorf(ctx, "unable to get cat list: %s", err)
		return nil, marvin.NewProtoStatusResponse(&ErrorResponse{
			Error: "unable to get cat list"}, http.StatusInternalServerError)
	}

	return &CatsResponse{
		Total: int32(len(cats)),
		Cats:  cats,
	}, nil
}
