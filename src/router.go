package src

import (
	"github.com/gin-gonic/gin"
)

func problemsRouter(r *gin.Engine) {
	router := r.Group("/problems")
	router.GET("/:OJ/:ProblemID", GetProblemInfo)
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
	problemsRouter(r)
	submitRouter(r)
}
