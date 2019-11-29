package src

import (
	"github.com/gin-gonic/gin"
	"github.com/moooyo/VirtualJudge/tools"
	"log"
	"net/http"
)

type UserSubmit struct {
	UID      int    `json:"uid"`
	SID      int    `json:"sid"`
	OJ       int    `json:"oj"`
	Problem  int    `json:"problem"`
	Language int    `json:"language"`
	Source   string `json:"source"`
}

func SubmitProblem(r *gin.Context) {
	var submit UserSubmit
	if r.ShouldBindJSON(&submit) == nil {
		r.AbortWithStatusJSON(http.StatusOK, gin.H{
			"status":  tools.PostDataInvalid,
			"message": "Post data invalid.",
		})
		log.Fatal("being here")
	}

}
