package server

import (
    repositories "mxshs/movieLibrary/src/adapters/repositories"
	"mxshs/movieLibrary/src/domain"
	"mxshs/movieLibrary/src/handlers/api"
	"mxshs/movieLibrary/src/services"
	"net/http"
)

func BootstrapAPI(prefix string) {
	db := repositories.NewPgDB()
	actorSvc := services.NewActorService(db)
	actorHandler := api.NewActorHandler(actorSvc)
	movieSvc := services.NewMovieService(db)
	movieHandler := api.NewMovieHandler(movieSvc)
	movieActorSvc := services.NewMovieActorService(db)
	movieActorHandler := api.NewMovieActorHandler(movieActorSvc)
	authSvc := services.NewAuthService()
	authMiddleware := api.NewAuthHandler(authSvc)
	userSvc := services.NewUserService(db)
	userHandler := api.NewUserHandler(userSvc, authSvc)

	srv := http.NewServeMux()

	srv.HandleFunc("GET /movies/{mid}/actors", authMiddleware.Authenticate(movieActorHandler.GetMovieActors, domain.USR))
	srv.HandleFunc("GET /actors/{aid}/movies", authMiddleware.Authenticate(movieActorHandler.GetActorMovies, domain.USR))
	srv.HandleFunc("POST /movies/{mid}/actors/{aid}/", authMiddleware.Authenticate(movieActorHandler.CreateMovieActor, domain.ADM))
	srv.HandleFunc("DELETE /movies/{mid}/actors/{aid}/", authMiddleware.Authenticate(movieActorHandler.DeleteMovieActor, domain.ADM))

	srv.HandleFunc("GET /actors/{id}", authMiddleware.Authenticate(actorHandler.GetActor, domain.USR))
	srv.HandleFunc("PATCH /actors/{id}/", authMiddleware.Authenticate(actorHandler.UpdateActor, domain.ADM))
	srv.HandleFunc("DELETE /actors/{id}/", authMiddleware.Authenticate(actorHandler.DeleteActor, domain.ADM))
	srv.HandleFunc("GET /actors", authMiddleware.Authenticate(actorHandler.GetActors, domain.USR))
	srv.HandleFunc("POST /actors/", authMiddleware.Authenticate(actorHandler.CreateActor, domain.ADM))

	srv.HandleFunc("GET /movies/{id}", authMiddleware.Authenticate(movieHandler.GetMovie, domain.USR))
	srv.HandleFunc("PATCH /movies/{id}/", authMiddleware.Authenticate(movieHandler.UpdateMovie, domain.ADM))
	srv.HandleFunc("DELETE /movies/{id}/", authMiddleware.Authenticate(movieHandler.DeleteMovie, domain.ADM))
	srv.HandleFunc("GET /movies", authMiddleware.Authenticate(movieHandler.GetMovies, domain.USR))
	srv.HandleFunc("POST /movies/", authMiddleware.Authenticate(movieHandler.CreateMovie, domain.ADM))

	srv.HandleFunc("PATCH /users/{username}/", authMiddleware.Authenticate(userHandler.UpdateUser, domain.ADM))
	srv.HandleFunc("DELETE /users/{username}/", authMiddleware.Authenticate(userHandler.DeleteUser, domain.ADM))
	srv.HandleFunc("POST /users/login/", userHandler.LoginUser)
	srv.HandleFunc("POST /users/reset_token/", authMiddleware.Authenticate(userHandler.DeleteUser, domain.USR))
	srv.HandleFunc("POST /users/", authMiddleware.Authenticate(userHandler.CreateUser, domain.ADM))

	err := http.ListenAndServe(":3000", srv)
	if err != nil {
		panic(err)
	}
}
