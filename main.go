package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/moooyo/VirtualJudge/orm"
	"github.com/moooyo/VirtualJudge/src"
	"github.com/moooyo/VirtualJudge/tools"
)

func main() {
	config := tools.GetConfig()
	r := gin.Default()
	src.RegisterRouter(r)
	address := fmt.Sprintf("%s:%d", config.WebConfig.Address, config.WebConfig.Port)
	println(address)
	orm.DbInit()
	println(address)
	r.Run(address)
}
