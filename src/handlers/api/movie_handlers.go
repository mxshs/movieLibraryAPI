package api

import (
	"encoding/json"
	"fmt"
	adaptermodels "mxshs/movieLibrary/src/handlers/adapter_models"
	"mxshs/movieLibrary/src/services"
	"net/http"
	"strconv"
)

type MovieHandler struct {
	movieService *services.MovieService
}

func NewMovieHandler(movieService *services.MovieService) *MovieHandler {
	return &MovieHandler{movieService}
}

func (mh *MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	var m adaptermodels.BaseMovie

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := mh.movieService.CreateMovie(
		m.Title,
		m.Description,
		m.ReleaseDate.Time,
		m.Rating,
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

func (mh *MovieHandler) GetMovie(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	mid, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse %s as int", id), http.StatusBadRequest)
	}

	result, err := mh.movieService.GetMovie(mid)
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

func (mh *MovieHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	sortKey := r.URL.Query().Get("sort_by")
	sortOrder := r.URL.Query().Get("order_by")

	filterTitle := r.URL.Query().Get("title")
	filterActor := r.URL.Query().Get("actor")
	fmt.Println(filterActor)

	result, err := mh.movieService.GetMovies(sortKey, sortOrder, filterTitle, filterActor)
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

func (mh *MovieHandler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	mid, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse %s as int", id), http.StatusBadRequest)
	}

	var a adaptermodels.BaseMovie

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err = dec.Decode(&a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := mh.movieService.UpdateMovie(
		mid,
		a.Title,
		a.Description,
		a.ReleaseDate.Time,
		a.Rating,
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

func (mh *MovieHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	mid, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse %s as int", id), http.StatusBadRequest)
	}

	_, err = mh.movieService.DeleteMovie(mid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte(id))
}
