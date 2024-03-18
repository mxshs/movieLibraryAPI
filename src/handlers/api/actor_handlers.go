package api

import (
	"encoding/json"
	"fmt"
	adaptermodels "mxshs/movieLibrary/src/handlers/adapter_models"
	"mxshs/movieLibrary/src/services"
	"net/http"
	"strconv"
)

type ActorHandler struct {
	actorService *services.ActorService
}

func NewActorHandler(actorService *services.ActorService) *ActorHandler {
	return &ActorHandler{actorService}
}

func (ah *ActorHandler) CreateActor(w http.ResponseWriter, r *http.Request) {
	var a adaptermodels.BaseActor

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := ah.actorService.Create(
		a.Name,
		a.Gender,
		a.Birthdate.Time,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(response)
}

func (ah *ActorHandler) GetActor(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	aid, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse %s as int", id), http.StatusBadRequest)
	}

	result, err := ah.actorService.GetActor(aid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(response)
}

func (ah *ActorHandler) GetActors(w http.ResponseWriter, r *http.Request) {
	result, err := ah.actorService.GetActors()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(response)
}

func (ah *ActorHandler) UpdateActor(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	aid, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse %s as int", id), http.StatusBadRequest)
	}

	var a adaptermodels.BaseActor

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err = dec.Decode(&a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := ah.actorService.Update(
		aid,
		a.Name,
		a.Gender,
		a.Birthdate.Time,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(response)
}

func (ah *ActorHandler) DeleteActor(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	aid, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse %s as int", id), http.StatusBadRequest)
	}

	_, err = ah.actorService.Delete(aid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte(id))
}
