package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type CourseCategory struct {
	ID         uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name       string    `gorm:"size:255;not null;unique" json:"name"`
	SoftDelete bool      `gorm:"default:false" json:"soft_delete"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeleteAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"delete_at"`
}

func (u *CourseCategory) Prepare() {
	u.ID = 0
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.SoftDelete = false
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *CourseCategory) Validate(action string) error {
	if u.Name == "" {
		return errors.New("Required Course Category Name")
	}
	return nil
}

func (u *CourseCategory) SaveCourseCategory(db *gorm.DB) (*CourseCategory, error) {
	var err error
	err = db.Debug().Model(&CourseCategory{}).Create(&u).Error
	if err != nil {
		return &CourseCategory{}, err
	}
	return u, nil
}

func (u *CourseCategory) FindAllCourseCategory(db *gorm.DB) (*[]CourseCategory, error) {
	var err error
	coursecategory := []CourseCategory{}
	err = db.Debug().Model(&CourseCategory{}).Where("soft_delete = ?", false).Limit(100).Find(&coursecategory).Error
	if err != nil {
		return &[]CourseCategory{}, err
	}
	return &coursecategory, err
}

func (u *CourseCategory) FindCourseCategoryByID(db *gorm.DB, uid uint32) (*CourseCategory, error) {
	var err error
	err = db.Debug().Model(CourseCategory{}).Where("id = ? and soft_delete = ?", uid, false).Take(&u).Error
	if err != nil {
		return &CourseCategory{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &CourseCategory{}, errors.New("Course Category NorFound")
	}
	return u, err
}

func (u *CourseCategory) UpdateACourseCategory(db *gorm.DB, uid uint32) (*CourseCategory, error) {
	db = db.Debug().Model(&CourseCategory{}).Where("id = ?", uid).Take(&CourseCategory{}).UpdateColumns(
		map[string]interface{}{
			"name":       u.Name,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &CourseCategory{}, db.Error
	}
	// This is the display the updated user
	err := db.Debug().Model(&CourseCategory{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &CourseCategory{}, err
	}
	return u, nil
}

func (u *CourseCategory) DeleteACourseCategory(db *gorm.DB, uid uint32) (int64, error) {
	// soft delete
	db = db.Debug().Model(&CourseCategory{}).Where("id = ?", uid).Take(&CourseCategory{}).UpdateColumns(
		map[string]interface{}{
			"soft_delete": true,
			"delete_at":   time.Now(),
		},
	)
	if db.Error != nil {
		return 0, db.Error
	}
	// This is the display the updated user
	err := db.Debug().Model(&CourseCategory{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return 0, err
	}
	return 0, nil
}
