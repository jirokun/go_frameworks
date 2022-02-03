package domain

import "context"

type PetRepository interface {
	Save(ctx context.Context, pet *Pet) error
	Find(ctx context.Context, id int64) (*Pet, error)
}

type Pet struct {
	ID   int
	Name string
	Tag  *string
}
