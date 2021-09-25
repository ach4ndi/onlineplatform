package controllers

import (
	"github.com/joho/godotenv"
	"github.com/ach4ndi/onlineplatform/api/auth"
	"github.com/ach4ndi/onlineplatform/api/models"
	"github.com/ach4ndi/onlineplatform/api/responses"
	"github.com/ach4ndi/onlineplatform/api/utils/formaterror"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"github.com/google/uuid"
	"strings"
	"os"
	"io"
)

func (server *Server) GetCourses(w http.ResponseWriter, r *http.Request) {
	course := models.Course{}

	courses, err := course.FindAllCourses(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}

func (server *Server) GetCoursesLow(w http.ResponseWriter, r *http.Request) {
	course := models.Course{}

	courses, err := course.FindAllCoursesCheap(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}

func (server *Server) GetCoursesHigh(w http.ResponseWriter, r *http.Request) {
	course := models.Course{}

	courses, err := course.FindAllCoursesExpensive(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
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
	courseGoten, err := course.FindCourseByID(server.DB, uint32(uid))
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
	tokenID = uid
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


	// <num> limit filesize <num> in MB
	limit_size, err := strconv.ParseInt(os.Getenv("IMG_LIMIT"),10,32)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	r.ParseMultipartForm(limit_size)
	course := models.Course{}

	course.CourseCategoryID = r.Form.Get("course_category_id")
	course.UserID = r.Form.Get("user_id")
	course.Name = r.Form.Get("course_name")
	course.Description = r.Form.Get("course_desc")

	res, err := strconv.ParseInt(r.Form.Get("price"),10,32)

	if err != nil {
		fmt.Printf(r.Form.Get("default_price"))
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	course.Price = uint32(res)
	course.Duration = r.Form.Get("duration")
	course.IsFree = r.Form.Get("is_free")
	course.IsOnline = r.Form.Get("is_online")

	file, handler, err := r.FormFile("OpeningImage")
	imageName := ""

	switch err {
		case nil:
			if os.Getenv("CLOUDINARY_APIKEY") == ""{
				// if empty will used local
				imageName = "product_"+strings.Replace(uuid.New().String(), "-", "", -1) + ".png"
			
				f, err := os.OpenFile(os.Getenv("IMG_DIR")+"/images/products/"+imageName, os.O_WRONLY|os.O_CREATE, 0666)
				if err != nil {
					fmt.Println(err)
				}
				defer f.Close()

				io.Copy(f, file)
			}
			if os.Getenv("CLOUDINARY_APIKEY") != ""{
				cld, _ := cloudinary.NewFromParams(os.Getenv("CLOUDINARY_CLOUDNAME"), os.Getenv("CLOUDINARY_APIKEY"), os.Getenv("CLOUDINARY_APISecret"))

				resp, err := cld.Upload.Upload(file, imageName, uploader.UploadParams{})

				imageName = resp.SecureURL
			}
		case http.ErrMissingFile:
			fmt.Printf("no file")
		default:
			fmt.Printf("errs")
	}
	
	if imageName == ""{
		imageName = nil
	}

	course.OpeningImage = imageName
	defer file.Close()

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

	_, err = course.DeleteAPost(server.DB, uint32(uid))
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
	limit_size, err := strconv.ParseInt(os.Getenv("IMG_LIMIT"),10,32)

	tokenID = uid
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

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	r.ParseMultipartForm(limit_size)
	course := models.Course{}

	course.CourseCategoryID = r.Form.Get("course_category_id")
	course.UserID = r.Form.Get("user_id")
	course.Name = r.Form.Get("course_name")
	course.Description = r.Form.Get("course_desc")

	res, err := strconv.ParseInt(r.Form.Get("price"),10,32)

	if err != nil {
		fmt.Printf(r.Form.Get("default_price"))
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	course.Price = uint32(res)
	course.Duration = r.Form.Get("duration")
	course.IsFree = r.Form.Get("is_free")
	course.IsOnline = r.Form.Get("is_online")

	file, handler, err := r.FormFile("OpeningImage")
	imageName := ""

	switch err {
		case nil:
			if os.Getenv("CLOUDINARY_APIKEY") == ""{
				// if empty will used local
				imageName = "product_"+strings.Replace(uuid.New().String(), "-", "", -1) + ".png"
			
				f, err := os.OpenFile(os.Getenv("IMG_DIR")+"/images/products/"+imageName, os.O_WRONLY|os.O_CREATE, 0666)
				if err != nil {
					fmt.Println(err)
				}
				defer f.Close()

				io.Copy(f, file)
			}
			if os.Getenv("CLOUDINARY_APIKEY") != ""{
				cld, _ := cloudinary.NewFromParams(os.Getenv("CLOUDINARY_CLOUDNAME"), os.Getenv("CLOUDINARY_APIKEY"), os.Getenv("CLOUDINARY_APISecret"))
				resp, err := cld.Upload.Upload(file, imageName, uploader.UploadParams{})

				imageName = resp.SecureURL
			}
		case http.ErrMissingFile:
			fmt.Printf("no file")
		default:
			fmt.Printf("errs")
	}
	
	if imageName == ""{
		imageName = nil
	}

	course.OpeningImage = imageName
	defer file.Close()

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

func (server *Server) SearchCourse(w http.ResponseWriter, r *http.Request) {
	course := models.Course{}

	courses, err := course.SearchCourseName(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, courses)
}