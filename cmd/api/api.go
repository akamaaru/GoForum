package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/akamaaru/go-forum/service/comment"
	"github.com/akamaaru/go-forum/service/post"
	"github.com/akamaaru/go-forum/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	postStore := post.NewStore(s.db)
	postHandler := post.NewHandler(postStore, userStore)
	postHandler.RegisterRoutes(subrouter)

	commentStore := comment.NewStore(s.db)
	commentHandler := comment.NewHandler(commentStore, userStore)
	commentHandler.RegisterRoutes(subrouter)

	log.Println("Listening server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}