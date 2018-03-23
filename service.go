package cats

import (
	"context"
	"net/http"
	"strings"

	"github.com/NYTimes/marvin"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"

	"google.golang.org/appengine/user"
)

// service is the heart of our process. It is a one stop shop for dependency injection,
// server middleware and endpoints.
type service struct {
	db DB
}

// NewService returns a new marvin.JSONService to register with marvin
func NewService() *service {
	return &service{db: NewDB()}
}

func (s *service) HTTPMiddleware(h http.Handler) http.Handler {
	return h
}

// Middleware will contain our check for logged in users. If users are not logged in,
// this will short circuit their request and respond with a 401
func (s *service) Middleware(ep endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		usr := user.Current(ctx)
		if usr == nil {
			return nil, marvin.NewJSONStatusResponse(map[string]string{
				"error": "please log in"}, http.StatusUnauthorized)
		}

		// all clear, propagate to the next layer down
		return ep(ctx, r)
	}
}

func (s *service) Options() []httptransport.ServerOption {
	return []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(func(ctx context.Context, err error, w http.ResponseWriter) {
			// check proto/json by inspecting url
			path := ctx.Value(httptransport.ContextKeyRequestPath).(string)
			if strings.HasSuffix(path, ".json") {
				httptransport.EncodeJSONResponse(ctx, w, err)
				return
			}
			marvin.EncodeProtoResponse(ctx, w, err)
		}),
	}
}

func (s *service) RouterOptions() []marvin.RouterOption {
	// we dont need any fancy routing so override the default (gorilla) with stdlib
	return []marvin.RouterOption{marvin.RouterSelect("stdlib")}
}

func (s *service) JSONEndpoints() map[string]map[string]marvin.HTTPEndpoint {
	return map[string]map[string]marvin.HTTPEndpoint{
		"/list.json": {
			"GET": {
				Endpoint: s.listCats,
			},
		},
		"/add.json": {
			"POST": {
				Endpoint: s.addCat,
				Decoder:  decodeCat,
			},
		},
	}
}

func (s *service) ProtoEndpoints() map[string]map[string]marvin.HTTPEndpoint {
	return map[string]map[string]marvin.HTTPEndpoint{
		"/list.proto": {
			"GET": {
				Endpoint: s.listCats,
			},
		},
		"/add.proto": {
			"POST": {
				Endpoint: s.addCat,
				Decoder:  decodeCatProto,
			},
		},
	}
}
