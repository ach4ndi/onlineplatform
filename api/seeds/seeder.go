package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/ach4ndi/onlineplatform/api/models"
)

var user_status = []models.UserStatus{
	models.UserStatus{
		LevelName: "Admin",
		LevelNum: 1,
		SoftDelete : false
	},
	models.UserStatus{
		LevelName: "User",
		LevelNum: 5,
		SoftDelete : false
	},
	models.UserStatus{
		LevelName: "Teller",
		LevelNum: 6,
		SoftDelete : True
	},
}

var users = []models.User{
	models.User{
		Name: "Admin Kita",
		UserStatusID: 1,
		Email:    "admin@gmail.com",
		Password: "password",
		SoftDelete : false
		SoftDelete : false,
	},
	models.User{
		Nickname: "Rayhan Kuncoro",
		UserStatusID: 2,
		Email:    "kuncoro@gmail.com",
		Password: "password",
		SoftDelete : false,
	},
	models.User{
		Nickname: "Monika Lukiawan",
		UserStatusID: 2,
		Email:    "monika@gmail.com",
		Password: "password",
		SoftDelete : false,
	},
	models.User{
		Nickname: "Yayha Rahmati",
		UserStatusID: 2,
		Email:    "yahya@gmail.com",
		Password: "password",
		SoftDelete : false,
	},
	models.User{
		Nickname: "Dede Kirana",
		UserStatusID: 2,
		Email:    "ddee@gmail.com",
		Password: "password",
		SoftDelete : false,
	},
	models.User{
		Nickname: "Ade Pratamana",
		UserStatusID: 2,
		Email:    "pratama@gmail.com",
		Password: "password",
		SoftDelete : true,
	},
}

var course_category=[]models.CourseCategory{
	models.CourseCategory{
		Name: "Olah Raga",
		SoftDelete : false,
	},
	models.CourseCategory{
		LevelName: "Teknologi",
		SoftDelete : false,
	},
	models.CourseCategory{
		LevelName: "Biologi",
		SoftDelete : false,
	},
	models.CourseCategory{
		LevelName: "Sains",
		SoftDelete : false,
	},
	models.CourseCategory{
		LevelName: "Mutimedia",
		SoftDelete : false,
	},
	models.CourseCategory{
		LevelName: "Kimia",
		SoftDelete : true,
	},
}

var course = []models.Course{
	models.Course{
		CourseCategoryID : 1,
		User: 2,
		Name: "Pencak Silat",
		Description: "Belajar Pencak Silat dengan baik dan benar selama 3 bulan, course ditempat",
		Durarion : 3
		Price : 300000,
		IsFree : false,
		IsOnline : false,
		OpeningImage : nil,
		SoftDelete : false,
	},
	models.Course{
		CourseCategoryID : 1,
		User: 2,
		Name: "Renang",
		Description: "Belajar dasar dasar berenang sampiai bisa, selama 3 bulan, course ditempat",
		Price : 200000,
		Durarion : 3
		IsFree : false,
		IsOnline : false,
		OpeningImage : nil,
		SoftDelete : false,
	},
	models.Course{
		CourseCategoryID : 2,
		User: 3,
		Name: "Pembelajaran Full stack Javascript",
		Description: "Pembelajaran Full stack selama 6 bulan",
		Durarion : 6
		Price : 600000,
		IsFree : false,
		IsOnline : true,
		OpeningImage : nil,
		SoftDelete : false,
	},
	models.Course{
		CourseCategoryID : 2,
		User: 3,
		Name: "Pembelajaran Full stack Python",
		Description: "Pembelajaran Full stack selama 6 bulan",
		Durarion : 6
		Price : 600000,
		IsFree : false,
		IsOnline : true,
		OpeningImage : nil,
		SoftDelete : true,
	},
}


var user_course = []models.UserCourse{
	models.UserCourse{
		UserID: 2,
		CourseID: 1,
		Buy: 1,
		SoftDelete : false,
	},
	models.UserCourse{
		UserID: 2,
		CourseID: 2,
		Buy: 1,
		SoftDelete : false,
	},
	models.UserCourse{
		UserID: 3,
		CourseID: 2,
		Buy: 1,
		SoftDelete : false,
	},
	models.UserCourse{
		UserID: 4,
		CourseID: 2,
		Buy: 1,
		SoftDelete : false,
	},
	models.UserCourse{
		UserID: 5,
		CourseID: 4,
		Buy: 1,
		SoftDelete : true,
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