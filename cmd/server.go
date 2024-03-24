package main

import (
	repositories "mxshs/movieLibrary/src/adapters/repositories"
	"mxshs/movieLibrary/src/domain"
	"mxshs/movieLibrary/src/handlers/api"
	"mxshs/movieLibrary/src/services"
	"net/http"

	"github.com/swaggo/http-swagger"

	_ "mxshs/movieLibrary/docs"
)

func BootstrapAPI(prefix string) {
	db := repositories.NewPgDB()
	actorSvc := services.NewActorService(db, db)
	actorHandler := api.NewActorHandler(actorSvc)
	movieSvc := services.NewMovieService(db, db)
	movieHandler := api.NewMovieHandler(movieSvc)
	movieActorSvc := services.NewMovieActorService(db)
	movieActorHandler := api.NewMovieActorHandler(movieActorSvc)
	authSvc := services.NewAuthService()
	authMiddleware := api.NewAuthHandler(authSvc)
	userSvc := services.NewUserService(db)
	userHandler := api.NewUserHandler(userSvc, authSvc)

	srv := http.NewServeMux()

	srv.HandleFunc("GET /api/v1/movies/{mid}/actors", authMiddleware.Authenticate(movieActorHandler.GetMovieActors, domain.USR))
	srv.HandleFunc("GET /api/v1/actors/{aid}/movies", authMiddleware.Authenticate(movieActorHandler.GetActorMovies, domain.USR))
	srv.HandleFunc("POST /api/v1/movies/{mid}/actors/{aid}/", authMiddleware.Authenticate(movieActorHandler.CreateMovieActor, domain.ADM))
	srv.HandleFunc("DELETE /api/v1/movies/{mid}/actors/{aid}/", authMiddleware.Authenticate(movieActorHandler.DeleteMovieActor, domain.ADM))

	srv.HandleFunc("GET /api/v1/actors/{id}", authMiddleware.Authenticate(actorHandler.GetActor, domain.USR))
	srv.HandleFunc("PATCH /api/v1/actors/{id}/", authMiddleware.Authenticate(actorHandler.UpdateActor, domain.ADM))
	srv.HandleFunc("DELETE /api/v1/actors/{id}/", authMiddleware.Authenticate(actorHandler.DeleteActor, domain.ADM))
	srv.HandleFunc("GET /api/v1/actors", authMiddleware.Authenticate(actorHandler.GetActors, domain.USR))
	srv.HandleFunc("POST /api/v1/actors/", authMiddleware.Authenticate(actorHandler.CreateActor, domain.ADM))

	srv.HandleFunc("GET /api/v1/movies/{id}", authMiddleware.Authenticate(movieHandler.GetMovie, domain.USR))
	srv.HandleFunc("PATCH /api/v1/movies/{id}/", authMiddleware.Authenticate(movieHandler.UpdateMovie, domain.ADM))
	srv.HandleFunc("DELETE /api/v1/movies/{id}/", authMiddleware.Authenticate(movieHandler.DeleteMovie, domain.ADM))
	srv.HandleFunc("GET /api/v1/movies", authMiddleware.Authenticate(movieHandler.GetMovies, domain.USR))
	srv.HandleFunc("POST /api/v1/movies/", authMiddleware.Authenticate(movieHandler.CreateMovie, domain.ADM))

	srv.HandleFunc("PATCH /api/v1/users/{username}/", authMiddleware.Authenticate(userHandler.UpdateUser, domain.ADM))
	srv.HandleFunc("DELETE /api/v1/users/{username}/", authMiddleware.Authenticate(userHandler.DeleteUser, domain.ADM))
	srv.HandleFunc("POST /api/v1/users/login/", userHandler.LoginUser)
	srv.HandleFunc("POST /api/v1/users/reset_token/", authMiddleware.Authenticate(userHandler.DeleteUser, domain.USR))
	srv.HandleFunc("POST /api/v1/users/", authMiddleware.Authenticate(userHandler.CreateUser, domain.ADM))
	srv.HandleFunc("GET /api/v1/users", authMiddleware.Authenticate(userHandler.GetUsers, domain.ADM))

	srv.HandleFunc("/swagger/*", httpSwagger.WrapHandler)

	err := http.ListenAndServe(":3000", srv)
	if err != nil {
		panic(err)
	}
}

// @title			Movies & Actors API
// @version		1.0
// @description	...

// @host		localhost:3000
// @BasePath	/api/v1

// @securityDefinitions.apikey	Bearer
// @in							header
// @name						Authorization
func main() {
	BootstrapAPI("/api/v1")
}
