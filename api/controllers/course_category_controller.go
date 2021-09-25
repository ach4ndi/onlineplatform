package controllers

import (
	"github.com/joho/godotenv"
	"github.com/ach4ndi/onlineplatform/api/auth"
	"github.com/ach4ndi/onlineplatform/api/models"
	"github.com/ach4ndi/onlineplatform/api/responses"
	"github.com/ach4ndi/onlineplatform/api/utils/formaterror"
	"github.com/gorilla/mux"
	"encoding/json"
	"errors"
	"io/ioutil"
	"fmt"
	"net/http"
	"strconv"
	"log"
	"os"
)

func (server *Server) GetCourseCategories(w http.ResponseWriter, r *http.Request) {
	coursecategory := models.CourseCategory{}

	coursecategories, err := coursecategory.FindAllCourseCategory(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, coursecategories)
}

func (server *Server) GetCourseCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	coursecategory := models.CourseCategory{}
	coursecategoryeGotten, err := coursecategory.FindCourseCategoryByID(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, coursecategoryeGotten)
}

func (server *Server) UpdateCourseCategory(w http.ResponseWriter, r *http.Request) {
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
	coursecategory := models.CourseCategory{}
	user := models.User{}
	err = json.Unmarshal(body, &user)
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

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	limit_level, err := strconv.Atoi(os.Getenv("LIMITLV")) 

	user := User{}
	user = user.FindUserByID(tokenID).Take(&u).Error
	if user != nil {
		if user.UserStatus.LevelNum != limit_level{
			responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
			return
		}
	}
	
	user.Prepare()
	err = user.Validate("update")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedCourseCategory, err := coursecategory.UpdateACourse(server.DB, uint32(uid))
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedCourseCategory)
}

func (server *Server) DeleteCourseCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	coursecategory := models.CourseCategory{}

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

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	limit_level = strconv.Atoi(os.Getenv("LIMITLV")) 

	err = db.Debug().Model(User{}).Where("id = ?", tokenID).Take(&u).Error
	if err != nil {
		if err.UserStatus.LevelNum != limit_level{
			responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
			return
		}
	}

	_, err = coursecategory.DeleteACourseCategory(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	responses.JSON(w, http.StatusNoContent, "")
}

func (server *Server) CreateCourseCategory(w http.ResponseWriter, r *http.Request) {
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

	limit_level = strconv.Atoi(os.Getenv("LIMITLV")) 

	err = db.Debug().Model(User{}).Where("id = ?", tokenID).Take(&u).Error
	if err != nil {
		if err.UserStatus.LevelNum != limit_level{
			responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
			return
		}
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	coursecategory := models.CourseCategory{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}



	coursecategory.Prepare()
	err = coursecategory.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	usercourseCreated, err := coursecategory.SaveCourseCategory(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, usercourseCreated.ID))
	responses.JSON(w, http.StatusCreated, usercourseCreated)
}