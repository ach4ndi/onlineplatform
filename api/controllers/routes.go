package controllers

import "github.com/ach4ndi/onlineplatform/api/middlewares"

func (s *Server) initializeRoutes() {
	//user api
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.UserLogin)).Methods("POST")
	s.Router.HandleFunc("/register", middlewares.SetMiddlewareJSON(s.UserRegister)).Methods("POST")

	s.Router.HandleFunc("/user", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/user/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/user/{id}", middlewares.SetMiddlewareAuthentication(s.UpdateUser)).Methods("PUT")
	s.Router.HandleFunc("/user/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	// user status
	s.Router.HandleFunc("/user_status", middlewares.SetMiddlewareJSON(s.GetUserAllStatus)).Methods("GET")
	s.Router.HandleFunc("/user_status", middlewares.SetMiddlewareJSON(s.CreateUserStatus)).Methods("POST")
	s.Router.HandleFunc("/user_status/{id}", middlewares.SetMiddlewareJSON(s.GetUserStatus)).Methods("GET")
	s.Router.HandleFunc("/user_status/{id}", middlewares.SetMiddlewareAuthentication(s.UpdateUserStatus)).Methods("PUT")
	s.Router.HandleFunc("/user_status/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUserStatus)).Methods("DELETE")

	s.Router.HandleFunc("/user_course", middlewares.SetMiddlewareJSON(s.GetUserCourses)).Methods("GET")
	s.Router.HandleFunc("/user_course", middlewares.SetMiddlewareJSON(s.CreateUserCourse)).Methods("POST")
	s.Router.HandleFunc("/user_course/{id}", middlewares.SetMiddlewareJSON(s.GetUserCourse)).Methods("GET")
	s.Router.HandleFunc("/user_course/{id}", middlewares.SetMiddlewareAuthentication(s.UpdateUserCourse)).Methods("PUT")
	s.Router.HandleFunc("/user_course/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUserCourse)).Methods("DELETE")

	s.Router.HandleFunc("/course", middlewares.SetMiddlewareJSON(s.GetCourses)).Methods("GET")
	s.Router.HandleFunc("/course_low", middlewares.SetMiddlewareJSON(s.GetCoursesLow)).Methods("GET")
	s.Router.HandleFunc("/course_high", middlewares.SetMiddlewareJSON(s.GetCoursesHigh)).Methods("GET")
	s.Router.HandleFunc("/course_free", middlewares.SetMiddlewareJSON(s.GetCoursesFree)).Methods("GET")

	s.Router.HandleFunc("/course", middlewares.SetMiddlewareJSON(s.CreateCourse)).Methods("POST")
	s.Router.HandleFunc("/course/{id}", middlewares.SetMiddlewareJSON(s.GetCourse)).Methods("GET")
	s.Router.HandleFunc("/course/{id}", middlewares.SetMiddlewareAuthentication(s.UpdateCourse)).Methods("PUT")
	s.Router.HandleFunc("/course/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteCourse)).Methods("DELETE")

	s.Router.HandleFunc("/course_category", middlewares.SetMiddlewareJSON(s.GetCourseCategories)).Methods("GET")
	s.Router.HandleFunc("/course_category", middlewares.SetMiddlewareJSON(s.CreateCourseCategory)).Methods("POST")
	s.Router.HandleFunc("/course_category/{id}", middlewares.SetMiddlewareJSON(s.GetCourseCategory)).Methods("GET")
	s.Router.HandleFunc("/course_category/{id}", middlewares.SetMiddlewareAuthentication(s.UpdateCourseCategory)).Methods("PUT")
	s.Router.HandleFunc("/course_category/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteCourseCategory)).Methods("DELETE")

	s.Router.HandleFunc("/course_search", middlewares.SetMiddlewareJSON(s.SearchCourse)).Methods("POST")
	s.Router.HandleFunc("/popular_course_category", middlewares.SetMiddlewareJSON(s.PopularUserCourse)).Methods("GET")
	s.Router.HandleFunc("/stat", middlewares.SetMiddlewareAuthentication(middlewares.SetMiddlewareJSON(s.Stat))).Methods("GET")
}
