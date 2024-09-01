package schemas

import "gorm.io/gorm"

type Opening struct {
	gorm.Model
	ID          string
	Role        string
	Description string
	Company     string
	IsRemote    bool
	Link        string
	Salary      int64
}
