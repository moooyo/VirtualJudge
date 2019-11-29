package src

import (
	"github.com/gin-gonic/gin"
)

func problemRouter(r *gin.Engine) {
	router := r.Group("/problem")
	router.POST("", SubmitProblem)
}

func RegisterRouter(r *gin.Engine) {
	problemRouter(r)
}
