package controllers

import "github.com/ach4ndi/onlineplatform/api/middlewares"

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.UserLogin)).Methods("POST")
	s.Router.HandleFunc("/register", middlewares.SetMiddlewareJSON(s.UserRegister)).Methods("POST")

	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.UpdateUser)).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	s.Router.HandleFunc("/user/category", middlewares.SetMiddlewareJSON(s.GetUserAllStatus)).Methods("GET")
	s.Router.HandleFunc("/user/category/{id}", middlewares.SetMiddlewareJSON(s.GetUserStatus)).Methods("GET")
	s.Router.HandleFunc("/user/category/{id}", middlewares.SetMiddlewareAuthentication(s.UpdateUserStatus)).Methods("PUT")
	s.Router.HandleFunc("/user/category/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUserStatus)).Methods("DELETE")