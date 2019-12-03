package orm

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Problem struct {
	ID           int `gorm:"AUTO_INCREMENT; primary_key"`
	OJ           int
	ProblemID    int
	ProblemName  string
	Description  string
	Input        []string
	Output       []string
	SampleInput  []string
	SampleOutput []string
	Language     []int
	UpdateAt     time.Time
}
