package controllers

import (
	"log"
	"net/http"

	"github.com/ach4ndi/onlineplatform/api/models"
	"github.com/ach4ndi/onlineplatform/api/responses"
)

func (server *Server) Stat(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	course := models.Course{}

	user_count, err := user.GetUserCount(server.DB)
	course_count, err := course.GetCourseCount(server.DB)
	coursefree_count, err := course.GetFreeCourseCount(server.DB)

	if err != nil {
		log.Print("sad .env file found")
	}

	data := map[string]interface{}{
		"User":       user_count,
		"Course":     course_count,
		"FreeCourse": coursefree_count,
	}

	responses.JSON(w, http.StatusOK, data)

}
