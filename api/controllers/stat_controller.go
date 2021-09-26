package controllers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/ach4ndi/onlineplatform/api/auth"
	"github.com/ach4ndi/onlineplatform/api/models"
	"github.com/ach4ndi/onlineplatform/api/responses"
	"github.com/joho/godotenv"
)

func (server *Server) Stat(w http.ResponseWriter, r *http.Request) {
	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	limit_level, err := strconv.Atoi(os.Getenv("LIMITLV"))

	user := models.User{}
	userGotten, err := user.FindUserByID(server.DB, tokenID)

	userst := models.UserStatus{}
	userstatusGotten, err := userst.FindUserStatusByID(server.DB, userGotten.UserStatusID)

	if userstatusGotten.LevelNum != uint32(limit_level) {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	user = models.User{}
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
