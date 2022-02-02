package usecase

import (
	"context"

	"github.com/jirokun/go_frameworks/chi/domain"
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

func (p *PetUsecase) AddPet(ctx context.Context, pet *domain.Pet) {
	p.petRepository.Save(ctx, pet)
}

func (p *PetUsecase) FindPetById(ctx context.Context, id int64) {
	p.petRepository.Find(ctx, id)
}
