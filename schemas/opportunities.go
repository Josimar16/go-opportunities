package schemas

import "gorm.io/gorm"

type Opportunities struct {
	gorm.Model
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   bool   `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}
