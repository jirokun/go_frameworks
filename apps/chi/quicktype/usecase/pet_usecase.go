package usecase

import (
	"context"

	"github.com/jirokun/go_frameworks/apps/chi/quicktype/domain"
	"github.com/jirokun/go_frameworks/apps/chi/quicktype/usecase/input"
	"github.com/jirokun/go_frameworks/apps/chi/quicktype/usecase/output"
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

func (p *PetUsecase) AddPet(ctx context.Context, pet *input.NewPet) error {
	return p.petRepository.Save(ctx, mapToDomainPet(pet))
}

func (p *PetUsecase) FindPetById(ctx context.Context, id int64) (*output.Pet, error) {
	dPet, err := p.petRepository.Find(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapFromDomainPet(dPet), nil
}
