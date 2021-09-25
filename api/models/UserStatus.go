package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type UserStatus struct {
	ID          uint32    `gorm:"primary_key;auto_increment" json:"id"`
	LevelName   string    `gorm:"size:255;not null;unique" json:"level_name"`
	LevelNum    uint32	  `gorm:"default:0" json:"level_num"`
	SoftDelete  bool      `gorm:"default:false" json:"soft_delete"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeleteAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"delete_at"`
}

func (u *UserStatus) Prepare() {
	u.ID = 0
	u.LevelName = html.EscapeString(strings.TrimSpace(u.LevelName))
	u.SoftDelete = false
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *UserStatus) Validate(action string) error {
	if u.LevelName == "" {
		return errors.New("Required Level Name")
	}
	return nil
}

func (u *UserStatus) SaveUserStatus(db *gorm.DB) (*UserStatus, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &UserStatus{}, err
	}
	return u, nil
}

func (u *UserStatus) FindAllStatus(db *gorm.DB) (*[]UserStatus, error) {
	var err error
	userstatus := []UserStatus{}
	err = db.Debug().Model(&UserStatus{}).Where("SoftDelete = ?", false).Limit(100).Find(&userstatus).Error
	if err != nil {
		return &[]UserStatus{}, err
	}
	return &userstatus, err
}

func (u *UserStatus) FindUserStatusByID(db *gorm.DB, uid uint32) (*UserStatus, error) {
	var err error
	err = db.Debug().Model(UserStatus{}).Where("id = ? and SoftDelete ?", uid, false).Take(&u).Error
	if err != nil {
		return &UserStatus{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &UserStatus{}, errors.New("UserStatus NorFound")
	}
	return u, err
}

func (u *UserStatus) UpdateAUserStatus(db *gorm.DB, uid uint32) (*UserStatus, error) {
	db = db.Debug().Model(&UserStatus{}).Where("id = ?", uid).Take(&UserStatus{}).UpdateColumns(
		map[string]interface{}{
			"level_name":  u.LevelName,
			"level_num": u.LevelNum,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &UserStatus{}, db.Error
	}
	// This is the display the updated user
	err := db.Debug().Model(&UserStatus{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &UserStatus{}, err
	}
	return u, nil
}

func (u *UserStatus) DeleteAUserStatus(db *gorm.DB, uid uint32) (*UserStatus, error) {
	db = db.Debug().Model(&UserStatus{}).Where("id = ?", uid).Take(&UserStatus{}).UpdateColumns(
		map[string]interface{}{
			"soft_delete":  true,
			"delete_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &UserStatus{}, db.Error
	}
	// This is the display the updated user
	err := db.Debug().Model(&UserStatus{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &UserStatus{}, err
	}
	return u, nil
}