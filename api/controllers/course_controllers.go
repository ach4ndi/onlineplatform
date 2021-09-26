package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/ach4ndi/onlineplatform/api/auth"
	"github.com/ach4ndi/onlineplatform/api/models"
	"github.com/ach4ndi/onlineplatform/api/responses"
	"github.com/ach4ndi/onlineplatform/api/utils/formaterror"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func (server *Server) GetCourses(w http.ResponseWriter, r *http.Request) {
	course := models.Course{}

	courses, err := course.FindAllCourses(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, courses)
}

func (server *Server) GetCoursesLow(w http.ResponseWriter, r *http.Request) {
	course := models.Course{}

	courses, err := course.FindAllCoursesCheap(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, courses)
}

func (server *Server) GetCoursesHigh(w http.ResponseWriter, r *http.Request) {
	course := models.Course{}

	courses, err := course.FindAllCoursesExpensive(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, courses)
}

func (server *Server) GetCoursesFree(w http.ResponseWriter, r *http.Request) {
	course := models.Course{}

	courses, err := course.FindAllCoursesFree(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, courses)
}

func (server *Server) GetCourse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	course := models.Course{}
	courseGoten, err := course.FindCourseByID(server.DB, uint64(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, courseGoten)
}

func (server *Server) UpdateCourse(w http.ResponseWriter, r *http.Request) {
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	err = godotenv.Load()
	tokenID := uid
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

	// <num> limit filesize <num> in MB
	limit_size, err := strconv.ParseInt(os.Getenv("IMG_LIMIT"), 10, 32)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	r.ParseMultipartForm(limit_size)
	course := models.Course{}

	c_coursecategoryid, err := strconv.ParseInt(r.Form.Get("course_category_id"), 10, 32)
	//c_userid, err := strconv.ParseInt(r.Form.Get("user_id"), 10, 32)

	course.CourseCategoryID = uint32(c_coursecategoryid)
	course.UserID = uint32(tokenID)
	course.Name = r.Form.Get("course_name")
	course.Description = r.Form.Get("course_desc")

	res, err := strconv.ParseInt(r.Form.Get("price"), 10, 32)

	if err != nil {
		fmt.Printf(r.Form.Get("default_price"))
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	c_duration, err := strconv.ParseInt(r.Form.Get("duration"), 10, 32)
	c_isfree, err := strconv.ParseInt(r.Form.Get("is_free"), 10, 32)
	c_isonline, err := strconv.ParseInt(r.Form.Get("is_online"), 10, 32)
	b_isfree := false
	b_isonline := false

	if c_isfree == 1 {
		b_isfree = true
	}

	if c_isonline == 1 {
		b_isonline = true
	}

	course.Price = uint32(res)
	course.Duration = uint32(c_duration)
	course.IsFree = b_isfree
	course.IsOnline = b_isonline

	if r.Form.Get("opening_image_update") == "1" {
		file, handler, err := r.FormFile("opening_image")
		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		imageName := ""

		switch err {
		case nil:
			if os.Getenv("CLOUDINARY_APIKEY") == "" {
				// if empty will used local
				imageName = "product_" + strings.Replace(uuid.New().String(), "-", "", -1) + ".png"

				f, err := os.OpenFile(os.Getenv("IMG_DIR")+"/images/products/"+imageName, os.O_WRONLY|os.O_CREATE, 0666)
				if err != nil {
					fmt.Println(err)
				}
				defer f.Close()

				io.Copy(f, file)
			}
			if os.Getenv("CLOUDINARY_APIKEY") != "" {
				var ctx = context.Background()
				cld, _ := cloudinary.NewFromParams(os.Getenv("CLOUDINARY_CLOUDNAME"), os.Getenv("CLOUDINARY_APIKEY"), os.Getenv("CLOUDINARY_APISecret"))
				resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{})
				if err != nil {
					fmt.Println(err)
				}
				imageName = resp.SecureURL
				defer file.Close()
			}
		case http.ErrMissingFile:
			defer file.Close()
			fmt.Printf("no file")
		default:
			defer file.Close()
			fmt.Printf("errs")
		}

		if imageName == "" {
			imageName = ""
		}
		course.OpeningImage = imageName
	} else {
		fmt.Println("user decide not update or anything")
	}

	course.Prepare()
	err = course.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if uid != course.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	courseCreated, err := course.UpdateACourse(server.DB, uint64(pid))
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, courseCreated.ID))
	responses.JSON(w, http.StatusCreated, courseCreated)

}

func (server *Server) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	course := models.Course{}

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

	_, err = course.DeleteACourse(server.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	responses.JSON(w, http.StatusNoContent, "")
}

func (server *Server) CreateCourse(w http.ResponseWriter, r *http.Request) {
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	err = godotenv.Load()
	// <num> limit filesize <num> in MB
	limit_size, err := strconv.ParseInt(os.Getenv("IMG_LIMIT"), 10, 32)

	tokenID := uid
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

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	r.ParseMultipartForm(limit_size)
	course := models.Course{}

	c_coursecategoryid, err := strconv.ParseInt(r.Form.Get("course_category_id"), 10, 32)
	//c_userid, err := strconv.ParseInt(r.Form.Get("user_id"), 10, 32)

	course.CourseCategoryID = uint32(c_coursecategoryid)
	course.UserID = uint32(tokenID)
	course.Name = r.Form.Get("course_name")
	course.Description = r.Form.Get("course_desc")

	res, err := strconv.ParseInt(r.Form.Get("price"), 10, 32)

	if err != nil {
		fmt.Printf(r.Form.Get("default_price"))
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	c_duration, err := strconv.ParseInt(r.Form.Get("duration"), 10, 32)
	c_isfree, err := strconv.ParseInt(r.Form.Get("is_free"), 10, 32)
	c_isonline, err := strconv.ParseInt(r.Form.Get("is_online"), 10, 32)
	b_isfree := false
	b_isonline := false

	if c_isfree == 1 {
		b_isfree = true
	}

	if c_isonline == 1 {
		b_isonline = true
	}

	course.Price = uint32(res)
	course.Duration = uint32(c_duration)
	course.IsFree = b_isfree
	course.IsOnline = b_isonline

	file, handler, err := r.FormFile("opening_image")
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	imageName := ""

	switch err {
	case nil:
		if os.Getenv("CLOUDINARY_APIKEY") == "" {
			// if empty will used local
			imageName = "product_" + strings.Replace(uuid.New().String(), "-", "", -1) + ".png"

			f, err := os.OpenFile(os.Getenv("IMG_DIR")+"/images/products/"+imageName, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err)
			}
			defer f.Close()

			io.Copy(f, file)
		}
		if os.Getenv("CLOUDINARY_APIKEY") != "" {
			var ctx = context.Background()
			cld, _ := cloudinary.NewFromParams(os.Getenv("CLOUDINARY_CLOUDNAME"), os.Getenv("CLOUDINARY_APIKEY"), os.Getenv("CLOUDINARY_APISecret"))
			resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{})

			if err != nil {
				fmt.Println(err)
			}

			imageName = resp.SecureURL
			defer file.Close()
		}
	case http.ErrMissingFile:
		defer file.Close()
		fmt.Printf("no file")
	default:
		defer file.Close()
		fmt.Printf("errs")
	}

	if imageName == "" {
		imageName = ""
	}

	course.OpeningImage = imageName

	course.Prepare()
	err = course.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if uid != course.UserID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	courseCreated, err := course.SaveCourse(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, courseCreated.ID))
	responses.JSON(w, http.StatusCreated, courseCreated)
}

type SearchField struct {
	Search string `json:"search"`
}

func (server *Server) SearchCourse(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	data := SearchField{}
	var result = json.Unmarshal(body, &data)

	if result != nil {
		fmt.Print("oK")
	}

	defer r.Body.Close()
	fmt.Print(data.Search)
	course := models.Course{}
	courses, err := course.SearchCourseName(server.DB, data.Search)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, courses)
}
