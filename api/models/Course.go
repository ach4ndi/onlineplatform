package models

import (
	"errors"
	"html"
	"strings"
	"time"

	//"github.com/ach4ndi/onlineplatform/api/models"
	"github.com/jinzhu/gorm"
)

type Course struct {
	ID                 uint32       	`gorm:"primary_key;auto_increment" json:"id"`
	CourseCategoryID   uint32       	`sql:"type:int REFERENCES coursecategories(id)"json:"course_category_id"`
	CourseCategory     CourseCategory   `json:"course_category"`
	User          	   User      		`json:"user"`
	UserID        	   uint32    		`sql:"type:int REFERENCES users(id)" json:"user_id"`
	Name               string       	`gorm:"size:255;not null" json:"course_name"`
	Description		   string       	`gorm:"size:255;not null" json:"course_desc"`
	Price              uint32       	`gorm:"default:0" json:"price"`
	Duration           uint32    		`gorm:"default:0" json:"duration"`
	IsFree			   bool				`gorm:"default:false" json:"is_free"`
	IsOnline		   bool				`gorm:"default:false" json:"is_online"`
	OpeningImage	   string			`gorm:"size:255;null" json:"course_op_img_url"`
	SoftDelete         bool      		`gorm:"default:false" json:"soft_delete"`
	CreatedAt          time.Time 		`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt          time.Time 		`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeleteAt           time.Time 		`gorm:"default:CURRENT_TIMESTAMP" json:"delete_at"`
}

func (p *Course) Prepare() {
	p.ID = 0
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	p.User = User{}
	p.Description = html.EscapeString(strings.TrimSpace(p.Description))
	p.Price = p.Price
	p.IsFree = false
	p.IsOnline = false
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Course) Validate() error {

	if p.Name == "" {
		return errors.New("Required Course Name")
	}
	if p.Description == "" {
		return errors.New("Required Description of Course detail")
	}
	if p.CourseCategoryID <1 {
		return errors.New("Required Course Category to Asseign Name")
	}
	if p.UserID < 1 {
		return errors.New("Required User to Assign")
	}
	if p.Price == 0 {
		p.IsFree = true
	}
	return nil
}

func (p *Course) SaveCourse(db *gorm.DB) (*Course, error) {
	var err error
	err = db.Debug().Model(&Course{}).Create(&p).Error
	if err != nil {
		return &Course{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &Course{}, err
		}
	}
	return p, nil
}

func (p *Course) FindAllCourses(db *gorm.DB) (*[]Course, error) {
	var err error
	courses := []Course{}
	err = db.Debug().Model(&Course{}).Limit(100).Find(&courses).Error
	if err != nil {
		return &[]Course{}, err
	}
	if len(courses) > 0 {
		for i, _ := range courses {
			err := db.Debug().Model(&User{}).Where("id = ?", courses[i].UserID).Take(&courses[i].User).Error
			if err != nil {
				return &[]Course{}, err
			}
		}
	}
	return &courses, nil
}

func (p *Course) FindAllCoursesCheap(db *gorm.DB) (*[]Course, error) {
	var err error
	courses := []Course{}
	err = db.Debug().Model(&Course{}).Order("price desc").Limit(100).Find(&courses).Error
	if err != nil {
		return &[]Course{}, err
	}
	if len(courses) > 0 {
		for i, _ := range courses {
			err := db.Debug().Model(&User{}).Where("id = ?", courses[i].UserID).Take(&courses[i].User).Error
			if err != nil {
				return &[]Course{}, err
			}
		}
	}
	return &courses, nil
}

func (p *Course) FindAllCoursesExpensive(db *gorm.DB) (*[]Course, error) {
	var err error
	courses := []Course{}
	err = db.Debug().Model(&Course{}).Order("price asc").Limit(100).Find(&courses).Error
	if err != nil {
		return &[]Course{}, err
	}
	if len(courses) > 0 {
		for i, _ := range courses {
			err := db.Debug().Model(&User{}).Where("id = ?", courses[i].UserID).Take(&courses[i].User).Error
			if err != nil {
				return &[]Course{}, err
			}
		}
	}
	return &courses, nil
}

func (p *Course) FindAllCoursesFree(db *gorm.DB) (*[]Course, error) {
	var err error
	courses := []Course{}
	err = db.Debug().Model(&Course{}).Where("price =? or isFree=?", 0, true).Limit(100).Find(&courses).Error
	if err != nil {
		return &[]Course{}, err
	}
	if len(courses) > 0 {
		for i, _ := range courses {
			err := db.Debug().Model(&User{}).Where("id = ?", courses[i].UserID).Take(&courses[i].User).Error
			if err != nil {
				return &[]Course{}, err
			}
		}
	}
	return &courses, nil
}

func (p *Course) SearchCourseName(db *gorm.DB, name string) (*Course, error) {
	var err error
	err = db.Debug().Model(&Course{}).Where("name = ?", name).Take(&p).Error
	if err != nil {
		return &Course{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &Course{}, err
		}
	}
	return p, nil
}

func (p *Course) FindCourseByID(db *gorm.DB, pid uint64) (*Course, error) {
	var err error
	err = db.Debug().Model(&Course{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Course{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &Course{}, err
		}
	}
	return p, nil
}

func (p *Course) UpdateACourse(db *gorm.DB) (*Course, error) {

	var err error
	err = db.Debug().Model(&Course{}).Where("id = ?", p.ID).Updates(Course{Name: p.Name, Description: p.Description, Price: p.Price, IsFree: p.IsFree, Duration: p.Duration, OpeningImage: p.OpeningImage, IsOnline: p.IsOnline,UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Course{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err != nil {
			return &Course{}, err
		}
	}
	return p, nil
}

func (p *Course) DeleteACourse(db *gorm.DB, pid uint64, uid uint32) (*Course, error) {

	var err error
	err = db.Debug().Model(&Course{}).Where("id = ?", p.ID).Updates(Course{SoftDelete: true, DeleteAt: time.Now()}).Error
	if err.Error != nil {
		return &Course{}, db.Error
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.UserID).Take(&p.User).Error
		if err.Error != nil {
			return &Course{}, db.Error
		}
	}
	return p, nil
}