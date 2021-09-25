package controllers

import (
	"net/http"
	"github.com/ach4ndi/onlineplatform/api/responses"
	//"github.com/ach4ndi/onlineplatform/api/models"
)

func (server *Server) Stat(w http.ResponseWriter, r *http.Request) {
	user, err := db.Debug().Model(&User{}).Count((&user_count))
	course, err := db.Debug().Model(&Course{}).Count((&course_count))
	coursefree, err := db.Debug().Model(&Course{}).Where("isFree = ? or Price = ?", true, 0).Count((&course_free_count))

	data := map[string]interface{}{
		"User": user_count,
		"Course": course_count,
		"FreeCourse": course_free_count,
	  }

	responses.JSON(w, http.StatusOK, data)

}