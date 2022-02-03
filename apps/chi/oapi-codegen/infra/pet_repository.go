package infra

import (
	"context"
	"fmt"

	"github.com/jirokun/go_frameworks/apps/chi/oapi-codegen/api_server_error"
	"github.com/jirokun/go_frameworks/apps/chi/oapi-codegen/domain"
)

type PetRepositoryImpl struct {
}

func NewPetRepositoryImpl() domain.PetRepository {
	return &PetRepositoryImpl{}
}

func (p *PetRepositoryImpl) Save(ctx context.Context, pet *domain.Pet) error {
	return nil
}

func (p *PetRepositoryImpl) Find(ctx context.Context, id int64) (*domain.Pet, error) {
	if id == 1 {
		tag := "タグ"
		return &domain.Pet{
			Id:   int(id),
			Name: "あいうえお",
			Tag:  &tag,
		}, nil
	}
	return nil, fmt.Errorf("id=%d: %w", id, api_server_error.ErrRecordNotFound)
}
