package controllers

import "github.com/cigetbudi/fullstack/api/middlewares"

func (s *Server) initializeRoutes() {
	// Home
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJson(s.Home)).Methods("GET")

	// Login
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJson(s.Login)).Methods("POST")

	// Users
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJson(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJson(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJson(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJson(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	// Posts
	s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJson(s.CreatePost)).Methods("POST")
	s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJson(s.GetPosts)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJson(s.GetPost)).Methods("GET")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJson(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareAuthentication(s.DeletePost)).Methods("DELETE")
}
