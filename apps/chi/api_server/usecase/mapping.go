package usecase

import "github.com/jirokun/go_frameworks/apps/chi/api_server/domain"

func mapFromDomainPet(pet *domain.Pet) *Pet {
	return &Pet{
		Id: int64(pet.Id),
		NewPet: NewPet{
			Name: pet.Name,
			Tag:  pet.Tag,
		},
	}
}
