package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type CourseCategory struct {
	ID          uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name   		string    `gorm:"size:255;not null;unique" json:"cource_category_name"`
	SoftDelete  bool      `gorm:"default:false" json:"soft_delete"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeleteAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"delete_at"`
}

func (u *CourseCategory) Prepare() {
	u.ID = 0
	u.Name = html.EscapeString(strings.TrimSpace(u.LevelName))
	u.SoftDelete = false
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *CourseCategory) Validate(action string) error {
	if u.Name == "" {
		return errors.New("Required Caourse Category Name")
	}
}

func (u *User) SaveCourseCategory(db *gorm.DB) (*CourseCategory, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &CourseCategory{}, err
	}
	return u, nil
}

func (u *CourseCategory) FindAllCourseCategory(db *gorm.DB) (*[]User, error) {
	var err error
	coursecategory := []CourseCategory{}
	err = db.Debug().Model(&CourseCategory{}).Where("SoftDelete = ?", false).Limit(100).Find(&coursecategory).Error
	if err != nil {
		return &[]CourseCategory{}, err
	}
	return &CourseCategory, err
}

func (u *CourseCategory) FindUserStatusByID(db *gorm.DB, uid uint32) (*CourseCategory, error) {
	var err error
	err = db.Debug().Model(CourseCategory{}).Where("id = ? and SoftDelete ?", uid, false).Take(&u).Error
	if err != nil {
		return &UserStatus{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &CourseCategory{}, errors.New("Course Category NorFound")
	}
	return u, err
}

func (u *CourseCategory) UpdateACourseCategory(db *gorm.DB, uid uint32) (*CourseCategory, error) {

	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&CourseCategory{}).Where("id = ?", uid).Take(&CourseCategory{}).UpdateColumns(
		map[string]interface{}{
			"cource_category_name":  u.Name
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &CourseCategory{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&CourseCategory{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &CourseCategory{}, err
	}
	return u, nil
}

func (u *UserStatus) DeleteACourseCategory(db *gorm.DB, uid uint32) (int64, error) {

	// because soft delete is only update to True

	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&CourseCategory{}).Where("id = ?", uid).Take(&CourseCategory{}).UpdateColumns(
		map[string]interface{}{
			"soft_delete":  true,
			"delete_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &UserStatus{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&UserStatus{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &UserStatus{}, err
	}
	return u, nil

func (u *UserStatus) DeleteBCourseCategory(db *gorm.DB, uid uint32) (int64, error) {
	db = db.Debug().Model(&CourseCategory{}).Where("id = ?", uid).Take(&CourseCategory{}).Delete(&CourseCategory{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}