package api

import (
	"encoding/json"
	"fmt"
	"mxshs/movieLibrary/src/services"
	"net/http"
	"strconv"
)

type MovieActorHandler struct {
	movieActorService *services.MovieActorService
}

func NewMovieActorHandler(movieActorService *services.MovieActorService) *MovieActorHandler {
	return &MovieActorHandler{movieActorService}
}

func (mah *MovieActorHandler) GetMovieActors(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("mid")

	mid, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse %s as int", id), http.StatusBadRequest)
	}

	actors, err := mah.movieActorService.GetMovieActors(mid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(actors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(response)
}

func (mah *MovieActorHandler) GetActorMovies(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("aid")

	aid, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse %s as int", id), http.StatusBadRequest)
	}

	movies, err := mah.movieActorService.GetActorMovies(aid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(movies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(response)
}

func (mah *MovieActorHandler) CreateMovieActor(w http.ResponseWriter, r *http.Request) {
	rawMid, rawAid := r.PathValue("mid"), r.PathValue("aid")

	mid, err := strconv.Atoi(rawMid)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse %s as int", rawMid), http.StatusBadRequest)
	}

	aid, err := strconv.Atoi(rawAid)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse %s as int", rawAid), http.StatusBadRequest)
	}

	err = mah.movieActorService.CreateMovieActor(mid, aid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (mah *MovieActorHandler) DeleteMovieActor(w http.ResponseWriter, r *http.Request) {
	rawMid, rawAid := r.PathValue("mid"), r.PathValue("aid")

	mid, err := strconv.Atoi(rawMid)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse %s as int", rawMid), http.StatusBadRequest)
	}

	aid, err := strconv.Atoi(rawAid)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse %s as int", rawAid), http.StatusBadRequest)
	}

	err = mah.movieActorService.DeleteMovieActor(mid, aid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
