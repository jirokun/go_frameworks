package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jirokun/go_frameworks/apps/chi/quicktype/api_server_error"
	"github.com/jirokun/go_frameworks/apps/chi/quicktype/usecase"
	"github.com/jirokun/go_frameworks/apps/chi/quicktype/usecase/input"
)

type PetHandler struct {
	petUsecase *usecase.PetUsecase
	handlers   http.Handler
}

func NewPetHandler(
	petUsecase *usecase.PetUsecase,
) *PetHandler {
	h := PetHandler{
		petUsecase: petUsecase,
	}
	router := chi.NewRouter()
	router.Get("/pets/{id}", h.FindPetById)
	router.Post("/pets", h.AddPet)

	return &PetHandler{
		petUsecase: petUsecase,
		handlers:   router,
	}
}

func handleError(w http.ResponseWriter, err error) {
	if errors.Is(err, api_server_error.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
		return
	}
	fmt.Println("Unknown Error")
}

func (s *PetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.handlers.ServeHTTP(w, r)
}

func (s *PetHandler) AddPet(w http.ResponseWriter, r *http.Request) {
	pet := input.NewPet{}
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

func (s *PetHandler) FindPetById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		handleError(w, err)
		return
	}

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
