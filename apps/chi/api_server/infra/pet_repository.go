package infra

import (
	"context"

	"github.com/jirokun/go_frameworks/chi/domain"
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
	return &domain.Pet{}, nil
}
