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

// CreateActor godoc
//
//	@Summary	Create a new actor entry
//	@Tags		actors
//	@Accept		json
//	@Produce	json
//	@Security	Bearer
//	@Param		actor	body		domain.Actor	true	"New actor"
//	@Success	200		{object}	domain.Actor
//	@Failure	400
//	@Failure	401
//	@Failure	404
//	@Router		/actors/ [post]
func (ah *ActorHandler) CreateActor(w http.ResponseWriter, r *http.Request) {
	var a adaptermodels.BaseActor

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := ah.actorService.CreateActor(
		a.Name,
		a.Gender,
		a.Birthdate,
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

// GetActor
//
//	@Summary	Get an actor entry by actor id
//	@Tags		actors
//	@Produce	json
//	@Security	Bearer
//	@Param		id	path		int	true	"actor id"	minimum(0)
//	@Success	200	{object}	domain.Actor
//	@Failure	400
//	@Failure	401
//	@Failure	404
//	@Router		/actors/{id} [get]
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

// GetActors
//
//	@Summary	Get all actors
//	@Tags		actors
//	@Produce	json
//	@Security	Bearer
//	@Success	200	{array}	domain.Actor
//	@Failure	400
//	@Failure	401
//	@Failure	404
//	@Router		/actors [get]
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

// UpdateActor
//
//	@Summary	Update an actor entry by actor id
//	@Tags		actors
//	@Accepts	json
//	@Produce	json
//	@Security	Bearer
//	@Param		id		path		int				true	"int valid"	minimum(0)
//	@Param		actor	body		domain.Actor	true	"Actor fields to update"
//	@Success	200		{object}	domain.Actor
//	@Failure	400
//	@Failure	401
//	@Failure	404
//	@Router		/actors/{id}/ [put]
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

	result, err := ah.actorService.UpdateActor(
		aid,
		a.Name,
		a.Gender,
		a.Birthdate,
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

// DeleteActor
//
//	@Summary	Remove an actor entry by actor id
//	@Tags		actors
//	@Security	Bearer
//	@Param		id	path	int	true	"actor id"	minimum(0)
//	@Success	200
//	@Failure	400
//	@Failure	401
//	@Failure	404
//	@Router		/actors/{id}/ [delete]
func (ah *ActorHandler) DeleteActor(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	aid, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse %s as int", id), http.StatusBadRequest)
	}

	_, err = ah.actorService.DeleteActor(aid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte(id))
}
