package usecase

import (
	"github.com/jirokun/go_frameworks/apps/chi/quicktype/domain"
	"github.com/jirokun/go_frameworks/apps/chi/quicktype/usecase/input"
	"github.com/jirokun/go_frameworks/apps/chi/quicktype/usecase/output"
)

func mapFromDomainPet(pet *domain.Pet) *output.Pet {
	return &output.Pet{
		ID:   int64(pet.ID),
		Name: pet.Name,
		Tag:  pet.Tag,
	}
}

func mapToDomainPet(pet *input.NewPet) *domain.Pet {
	return &domain.Pet{
		ID:   0,
		Name: pet.Name,
		Tag:  pet.Tag,
	}
}
