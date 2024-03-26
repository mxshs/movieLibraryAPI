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

// CreateMovie
//
//	@Summary	Create a new movie entry
//	@Tags		movies
//	@Accept		json
//	@Produce	json
//	@Security	Bearer
//	@Param		movie	body		domain.Movie	true	"New movie"
//	@Success	200		{object}	domain.Movie
//	@Failure	400
//	@Failure	401
//	@Failure	404
//	@Router		/movies/ [post]
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
		m.ReleaseDate,
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

// GetMovie
//
//	@Summary	Get a movie entry by movie id
//	@Tags		movies
//	@Produce	json
//	@Security	Bearer
//	@Param		id	path		int	true	"movie id"	minimum(0)
//	@Success	200	{object}	domain.Movie
//	@Failure	400
//	@Failure	401
//	@Failure	404
//	@Router		/movies/{id} [get]
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

// GetMovies
//
//	@Summary		Get multiple movies
//	@Description	Retrieve all movies or pass query parameters "title" and "actor" to search by movie title and actor name (partial match)
//	@Tags			movies
//	@Produce		json
//	@Security		Bearer
//	@Param			title	query	string	false	"any part of the movie title"
//	@Param			actor	query	string	false	"any part of the actor's name"
//	@Success		200		{array}	domain.Movie
//	@Failure		400
//	@Failure		401
//	@Failure		404
//	@Router			/movies [get]
func (mh *MovieHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	sortKey := r.URL.Query().Get("sort_by")
	sortOrder := r.URL.Query().Get("order_by")

	filterTitle := r.URL.Query().Get("title")
	filterActor := r.URL.Query().Get("actor")

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

// UpdateMovie
//
//	@Summary	Update a movie entry by movie id
//	@Tags		movies
//	@Accept		json
//	@Produce	json
//	@Security	Bearer
//	@Param		id		path		int				true	"movie id"	minimum(0)
//	@Param		movie	body		domain.Movie	true	"Movie fields to update"
//	@Success	200		{object}	domain.Movie
//	@Failure	400
//	@Failure	401
//	@Failure	404
//	@Router		/movies/{id}/ [patch]
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
		a.ReleaseDate,
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

// DeleteMovie
//
//	@Summary	Delete a movie entry by movie id
//	@Tags		movies
//	@Security	Bearer
//	@Param		id	path	int	true	"movie id"	minimum(0)
//	@Success	200
//	@Failure	400
//	@Failure	401
//	@Failure	404
//	@Router		/movies/{id}/ [delete]
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
