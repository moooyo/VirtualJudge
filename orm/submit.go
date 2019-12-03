package orm

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Submit struct {
	ID            int     `gorm:"AUTO_INCREMENT; primary_key"` // submit id
	SubmitProblem Problem `gorm:"foreignkey:ProblemID"`
	ProblemID     int
	RunID         int
	UID           int
	Source        string
	Language      int
	Status        int
	CreateAt      time.Time
}
