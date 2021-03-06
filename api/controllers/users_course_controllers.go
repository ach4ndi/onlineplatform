package controllers

import (
	//"github.com/joho/godotenv"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/ach4ndi/onlineplatform/api/auth"
	"github.com/ach4ndi/onlineplatform/api/models"
	"github.com/ach4ndi/onlineplatform/api/responses"
	"github.com/ach4ndi/onlineplatform/api/utils/formaterror"
	"github.com/gorilla/mux"
	//"github.com/google/uuid"
	//"strings"
	//"os"
	//"io"
)

func (server *Server) GetUserCourses(w http.ResponseWriter, r *http.Request) {
	usercourse := models.UserCourse{}

	usercourses, err := usercourse.FindAllUserCourse(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, usercourses)

}

func (server *Server) GetUserCourse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	usercourse := models.UserCourse{}
	usercourseGotten, err := usercourse.FindUserCourseByID(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, usercourseGotten)
}

func (server *Server) UpdateUserCourse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	usercourse := models.UserCourse{}
	err = json.Unmarshal(body, &usercourse)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if tokenID != uint32(uid) {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	usercourse.Prepare()
	err = usercourse.Validate("update")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedUsercourse, err := usercourse.UpdateAUserCourse(server.DB, uint32(uid))
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedUsercourse)
}

func (server *Server) DeleteUserCourse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	usercourse := models.UserCourse{}

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if tokenID != 0 && tokenID != uint32(uid) {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	_, err = usercourse.DeleteAUserCourse(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	responses.JSON(w, http.StatusNoContent, "")
}

func (server *Server) CreateUserCourse(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	usercourse := models.UserCourse{}
	err = json.Unmarshal(body, &usercourse)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	usercourse.Prepare()
	err = usercourse.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	usercourseCreated, err := usercourse.SaveUserStatus(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, usercourseCreated.ID))
	responses.JSON(w, http.StatusCreated, usercourseCreated)
}

func (server *Server) PopularUserCourse(w http.ResponseWriter, r *http.Request) {
	usercourse := models.UserCourse{}

	course_id, err := usercourse.GetCourseID(server.DB)
	fmt.Print(course_id)

	course := models.Course{}
	courseData, err := course.FindCourseByID(server.DB, uint64(course_id))

	course_category_id := courseData.CourseCategoryID

	course_category := models.CourseCategory{}
	coursecatData, err := course_category.FindCourseCategoryByID(server.DB, course_category_id)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, coursecatData.ID))
	responses.JSON(w, http.StatusCreated, coursecatData)
}
