package src

import (
	"github.com/gin-gonic/gin"
)

func problemRouter(r *gin.Engine) {
	/*
		submit := r.Group("/submit")
		submit.POST("", SubmitProblem)
	*/
}

func submitRouter(r *gin.Engine) {
	router := r.Group("/submit")
	router.POST("", SubmitProblem)
}

func RegisterRouter(r *gin.Engine) {
	problemRouter(r)
	submitRouter(r)
}
