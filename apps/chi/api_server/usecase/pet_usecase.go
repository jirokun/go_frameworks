package usecase

import (
	"context"

	"github.com/jirokun/go_frameworks/apps/chi/api_server/domain"
)

type PetUsecase struct {
	petRepository domain.PetRepository
}

func NewPetUsecase(
	petRepository domain.PetRepository,
) *PetUsecase {
	return &PetUsecase{
		petRepository: petRepository,
	}
}

func (p *PetUsecase) AddPet(ctx context.Context, pet *domain.Pet) error {
	return p.petRepository.Save(ctx, pet)
}

func (p *PetUsecase) FindPetById(ctx context.Context, id int64) (*Pet, error) {
	dPet, err := p.petRepository.Find(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapFromDomainPet(dPet), nil
}
