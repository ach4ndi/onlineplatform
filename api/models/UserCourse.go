package models

import (
	"errors"
	"fmt"
	"time"

	//"github.com/ach4ndi/onlineplatform/api/models"
	"github.com/jinzhu/gorm"
)

type UserCourse struct {
	ID         uint32    `gorm:"primary_key;auto_increment" json:"id"`
	UserID     uint32    `sql:"type:int REFERENCES users(id)" json:"user_id"`
	User       User      `json:"user"`
	CourseID   uint32    `sql:"type:int REFERENCES courses(id)" json:"course_id"`
	Course     Course    `json:"course"`
	Buy        bool      `gorm:"default:false" json:"buy"`
	SoftDelete bool      `gorm:"default:false" json:"soft_delete"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeleteAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"delete_at"`
}

func (u *UserCourse) Prepare() {
	u.ID = 0
	u.SoftDelete = false
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *UserCourse) Validate(action string) error {
	return nil
}

func (u *UserCourse) SaveUserStatus(db *gorm.DB) (*UserCourse, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &UserCourse{}, err
	}
	return u, nil
}

func (u *UserCourse) FindAllUserCourse(db *gorm.DB) (*[]UserCourse, error) {
	var err error
	usercourse := []UserCourse{}
	err = db.Debug().Model(&UserCourse{}).Where("soft_delete = ?", false).Limit(100).Find(&usercourse).Error
	if err != nil {
		return &[]UserCourse{}, err
	}
	return &usercourse, err
}

func (u *UserCourse) FindUserCourseByID(db *gorm.DB, uid uint32) (*UserCourse, error) {
	var err error
	err = db.Debug().Model(UserCourse{}).Where("id = ? and soft_delete = ?", uid, false).Take(&u).Error
	if err != nil {
		return &UserCourse{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &UserCourse{}, errors.New("UserCourse NorFound")
	}
	return u, err
}

func (u *UserCourse) UpdateAUserCourse(db *gorm.DB, uid uint32) (*UserCourse, error) {
	db = db.Debug().Model(&UserCourse{}).Where("id = ?", uid).Take(&UserCourse{}).UpdateColumns(
		map[string]interface{}{
			"Buy":        u.Buy,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &UserCourse{}, db.Error
	}
	// This is the display the updated user
	err := db.Debug().Model(&UserCourse{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &UserCourse{}, err
	}
	return u, nil
}

func (u *UserCourse) DeleteAUserCourse(db *gorm.DB, uid uint32) (int64, error) {
	db = db.Debug().Model(&UserCourse{}).Where("id = ?", uid).Take(&UserCourse{}).UpdateColumns(
		map[string]interface{}{
			"soft_delete": true,
			"delete_at":   time.Now(),
		},
	)
	if db.Error != nil {
		return 0, db.Error
	}
	// This is the display the updated user
	err := db.Debug().Model(&UserCourse{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return 0, err
	}
	return 0, nil
}

func (u *UserCourse) GetCourseID(db *gorm.DB) (int64, error) {
	var courseid int64

	out := UserCourse{}
	result := db.Debug().Model(&UserCourse{}).Where("soft_delete = ?", false).Group("course_id").Order("count(course_id) desc").First(&out)
	courseid = int64(out.CourseID)
	fmt.Print(result)
	//coursedata := models.Course{}
	//db = db.Debug().Model(&Course{}).Where("id = ?", courseid).Take(&coursedata)

	//course_category_id := coursedata.CourseCategoryID

	//db = db.Debug().Model(&CourseCategory{}).Where("id = ?", course_category_id).Take(&u)

	return courseid, nil
}
