package cats

import "context"

type testDB struct {
	MockGetCats func(context.Context) ([]*Cat, error)
	MockAddCat  func(context.Context, *Cat) error
}

func (d *testDB) GetCats(ctx context.Context) ([]*Cat, error) {
	return d.MockGetCats(ctx)
}

func (d *testDB) AddCat(ctx context.Context, c *Cat) error {
	return d.MockAddCat(ctx, c)
}
