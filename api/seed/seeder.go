package seed

import (
	"log"
	"time"

	"github.com/ach4ndi/onlineplatform/api/models"
	"github.com/jinzhu/gorm"
)

var user_status = []models.UserStatus{
	models.UserStatus{
		ID:         1,
		LevelName:  "Admin",
		LevelNum:   1,
		SoftDelete: false,
	},
	models.UserStatus{
		ID:         2,
		LevelName:  "User",
		LevelNum:   5,
		SoftDelete: false,
	},
	models.UserStatus{
		ID:         3,
		LevelName:  "Teller",
		LevelNum:   6,
		SoftDelete: true,
	},
}

var users = []models.User{
	models.User{
		ID:           1,
		UserStatusID: 1,
		UserStatus:   models.UserStatus{},
		Name:         "Admin Kita",
		Email:        "admin@gmail.com",
		Password:     "password",
		SoftDelete:   false,
	},
	models.User{
		ID:           2,
		UserStatusID: 2,
		UserStatus:   models.UserStatus{},
		Name:         "Rayhan Kuncoro",
		Email:        "kuncoro@gmail.com",
		Password:     "password",
		SoftDelete:   false,
	},
	models.User{
		ID:           3,
		UserStatusID: 2,
		UserStatus:   models.UserStatus{},
		Name:         "Monika Lukiawan",
		Email:        "monika@gmail.com",
		Password:     "password",
		SoftDelete:   false,
	},
	models.User{
		ID:           4,
		UserStatusID: 2,
		UserStatus:   models.UserStatus{},
		Name:         "Yayha Rahmati",
		Email:        "yahya@gmail.com",
		Password:     "password",
		SoftDelete:   false,
	},
	models.User{
		ID:           5,
		UserStatusID: 2,
		UserStatus:   models.UserStatus{},
		Name:         "Dede Kirana",
		Email:        "ddee@gmail.com",
		Password:     "password",
		SoftDelete:   false,
	},
	models.User{
		ID:           6,
		UserStatusID: 2,
		UserStatus:   models.UserStatus{},
		Name:         "Ade Pratamana",
		Email:        "pratama@gmail.com",
		Password:     "password",
		SoftDelete:   true,
	},
	models.User{
		ID:           7,
		UserStatusID: 2,
		UserStatus:   models.UserStatus{},
		Name:         "Nunung Kasandra",
		Email:        "nunug@gmail.com",
		Password:     "password",
		SoftDelete:   true,
	},
	models.User{
		ID:           8,
		UserStatusID: 2,
		UserStatus:   models.UserStatus{},
		Name:         "Reza Prateso",
		Email:        "praza@gmail.com",
		Password:     "password",
		SoftDelete:   true,
	},
	models.User{
		ID:           9,
		UserStatusID: 2,
		UserStatus:   models.UserStatus{},
		Name:         "Renaldi Phalevi",
		Email:        "renaldi@gmail.com",
		Password:     "password",
		SoftDelete:   true,
	},
	models.User{
		ID:           10,
		UserStatusID: 2,
		UserStatus:   models.UserStatus{},
		Name:         "Sofyan Hidayathulloh",
		Email:        "sofyan@gmail.com",
		Password:     "password",
		SoftDelete:   true,
	},
}

var course_category = []models.CourseCategory{
	models.CourseCategory{
		ID:         1,
		Name:       "Olah Raga",
		SoftDelete: false,
	},
	models.CourseCategory{
		ID:         2,
		Name:       "Teknologi",
		SoftDelete: false,
	},
	models.CourseCategory{
		ID:         3,
		Name:       "Biologi",
		SoftDelete: false,
	},
	models.CourseCategory{
		ID:         4,
		Name:       "Sains",
		SoftDelete: false,
	},
	models.CourseCategory{
		ID:         5,
		Name:       "Lain Lain",
		SoftDelete: false,
	},
}

var course = []models.Course{
	models.Course{
		ID:               1,
		CourseCategoryID: 1,
		CourseCategory:   models.CourseCategory{},
		User:             models.User{},
		UserID:           2,
		Name:             "Pencak Silat",
		Description:      "Belajar Pencak Silat dengan baik dan benar selama 3 bulan, course ditempat",
		Price:            300000,
		Duration:         90,
		IsFree:           false,
		IsOnline:         false,
		OpeningImage:     "",
		SoftDelete:       false,
		CreatedAt:        time.Time{},
	},
	models.Course{
		ID:               2,
		CourseCategoryID: 1,
		CourseCategory:   models.CourseCategory{},
		User:             models.User{},
		UserID:           1,
		Name:             "Renang",
		Description:      "Belajar dasar dasar berenang sampiai bisa, selama 3 bulan, course ditempat",
		Price:            200000,
		Duration:         90,
		IsFree:           false,
		IsOnline:         false,
		OpeningImage:     "",
		SoftDelete:       false,
		CreatedAt:        time.Time{},
	},
	models.Course{
		ID:               3,
		CourseCategoryID: 2,
		CourseCategory:   models.CourseCategory{},
		User:             models.User{},
		UserID:           1,
		Name:             "Pembelajaran Full stack Javascript",
		Description:      "Pembelajaran Full stack selama 6 bulan",
		Price:            600000,
		Duration:         180,
		IsFree:           false,
		IsOnline:         true,
		OpeningImage:     "",
		SoftDelete:       false,
		CreatedAt:        time.Time{},
	},
	models.Course{
		ID:               4,
		CourseCategoryID: 2,
		CourseCategory:   models.CourseCategory{},
		User:             models.User{},
		UserID:           1,
		Name:             "Pembelajaran Full stack Python",
		Description:      "Pembelajaran Full stack selama 6 bulan",
		Price:            600000,
		Duration:         180,
		IsFree:           false,
		IsOnline:         true,
		OpeningImage:     "",
		SoftDelete:       true,
		CreatedAt:        time.Time{},
	},
	models.Course{
		ID:               5,
		CourseCategoryID: 2,
		CourseCategory:   models.CourseCategory{},
		User:             models.User{},
		UserID:           1,
		Name:             "Pembelajaran Full stack Golang",
		Description:      "Pembelajaran Full stack selama 6 bulan",
		Price:            600000,
		Duration:         180,
		IsFree:           false,
		IsOnline:         true,
		OpeningImage:     "",
		SoftDelete:       false,
		CreatedAt:        time.Time{},
	},
	models.Course{
		ID:               6,
		CourseCategoryID: 2,
		CourseCategory:   models.CourseCategory{},
		User:             models.User{},
		UserID:           1,
		Name:             "Pembelajaran Full stack Ruby",
		Description:      "Pembelajaran Full stack selama 3 bulan",
		Price:            0,
		Duration:         190,
		IsFree:           true,
		IsOnline:         true,
		OpeningImage:     "",
		SoftDelete:       false,
		CreatedAt:        time.Time{},
	},
}

var user_course = []models.UserCourse{
	models.UserCourse{
		ID:         1,
		UserID:     2,
		User:       models.User{},
		CourseID:   1,
		Course:     models.Course{},
		Buy:        true,
		SoftDelete: false,
	},
	models.UserCourse{
		ID:         2,
		UserID:     2,
		User:       models.User{},
		CourseID:   2,
		Course:     models.Course{},
		Buy:        true,
		SoftDelete: false,
	},
	models.UserCourse{
		ID:         3,
		UserID:     3,
		User:       models.User{},
		CourseID:   2,
		Course:     models.Course{},
		Buy:        true,
		SoftDelete: false,
	},
	models.UserCourse{
		ID:         4,
		UserID:     4,
		User:       models.User{},
		CourseID:   2,
		Course:     models.Course{},
		Buy:        true,
		SoftDelete: false,
	},
	models.UserCourse{
		ID:         5,
		UserID:     5,
		User:       models.User{},
		CourseID:   4,
		Course:     models.Course{},
		Buy:        true,
		SoftDelete: true,
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Course{}, &models.CourseCategory{}, &models.User{}, &models.UserStatus{}, &models.UserCourse{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.Course{}, &models.CourseCategory{}, &models.User{}, &models.UserStatus{}, &models.UserCourse{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range user_status {
		err = db.Debug().Model(&models.UserStatus{}).Create(&user_status[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users status table: %v", err)
		}
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users status table: %v", err)
		}
	}

	for i, _ := range course_category {
		err = db.Debug().Model(&models.CourseCategory{}).Create(&course_category[i]).Error
		if err != nil {
			log.Fatalf("cannot seed course category table: %v", err)
		}
	}

	for i, _ := range course {
		err = db.Debug().Model(&models.Course{}).Create(&course[i]).Error
		if err != nil {
			log.Fatalf("cannot seed course table: %v", err)
		}
	}

	for i, _ := range user_course {
		err = db.Debug().Model(&models.UserCourse{}).Create(&user_course[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users status table: %v", err)
		}
	}
}
