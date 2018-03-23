package cats

import (
	"context"

	"github.com/pkg/errors"

	"google.golang.org/appengine/datastore"
)

type (
	DB interface {
		GetCats(context.Context) ([]*Cat, error)
		AddCat(context.Context, *Cat) error
	}

	db struct{}
)

func NewDB() *db {
	return &db{}
}

func (d *db) GetCats(ctx context.Context) ([]*Cat, error) {
	var cats []*Cat
	_, err := datastore.NewQuery("Cat").GetAll(ctx, &cats)
	return cats, errors.Wrap(err, "unable to get all cats")
}

func (d *db) AddCat(ctx context.Context, c *Cat) error {
	id, _, err := datastore.AllocateIDs(ctx, "Cat", nil, 1)
	if err != nil {
		return errors.Wrap(err, "unable to allocate ID")
	}
	c.Key = id
	_, err = datastore.Put(ctx, datastore.NewKey(ctx, "Cat", "", id, nil), c)
	return errors.Wrap(err, "unable to save the cat")
}
