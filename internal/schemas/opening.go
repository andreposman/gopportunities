package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role        string
	Description string
	Company     string
	IsRemote    bool
	Link        string
	Salary      int64
}

type OpeningResponse struct {
	ID          uint      `json: "id"`
	CreatedAt   time.Time `json: "createdAt"`
	UpdatedAt   time.Time `json: "updatedAt"`
	DeletedAt   time.Time `json: "deletedAt,omitempty"`
	Role        string    `json: "role"`
	Description string    `json: "description"`
	Company     string    `json: "company"`
	IsRemote    bool      `json: "isRemote"`
	Link        string    `json: "link"`
	Salary      int64     `json: "salary"`
}
