package orm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/moooyo/VirtualJudge/tools"
	"log"
)

func DbInit() {
	config := tools.GetConfig().MysqlConfig
	argsStr := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local",
		config.User, config.Password, config.Address, config.Port, config.DB)
	db, err := gorm.Open("mysql", argsStr)
	if err != nil {
		log.Fatal("connect error")
	}
	db.DB().SetMaxOpenConns(10)

}
