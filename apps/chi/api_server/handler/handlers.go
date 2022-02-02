package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/jirokun/go_frameworks/chi/api_server_error"
	"github.com/jirokun/go_frameworks/chi/domain"
	"github.com/jirokun/go_frameworks/chi/usecase"
)

type ServerInterfaceImpl struct {
	petUsecase *usecase.PetUsecase
}

func NewServerInterfaceImpl(
	petUsecase *usecase.PetUsecase,
) *ServerInterfaceImpl {
	return &ServerInterfaceImpl{
		petUsecase: petUsecase,
	}
}

func handleError(w http.ResponseWriter, err error) {
	if errors.Is(err, api_server_error.RecordNotFoundError{}) {
		fmt.Printf(err.Error())
		return
	}
	fmt.Println("Unknown Error")
}

func (s *ServerInterfaceImpl) AddPet(w http.ResponseWriter, r *http.Request) {
	pet := domain.Pet{}
	err := json.NewDecoder(r.Body).Decode(&pet)
	if err != nil {
		handleError(w, err)
		return
	}

	s.petUsecase.AddPet(r.Context(), &pet)
	b, err := json.Marshal(pet)
	if err != nil {
		handleError(w, err)
		return
	}
	w.Write(b)
}

func (s *ServerInterfaceImpl) FindPetById(w http.ResponseWriter, r *http.Request, id int64) {
	m, err := s.petUsecase.FindPetById(r.Context(), id)
	if err != nil {
		handleError(w, err)
		return
	}
	b, err := json.Marshal(m)
	if err != nil {
		handleError(w, err)
		return
	}

	w.Write(b)
}
