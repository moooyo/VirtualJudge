package src

import (
	"github.com/gin-gonic/gin"
	"github.com/moooyo/VirtualJudge/dispatch"
	"github.com/moooyo/VirtualJudge/tools"
	"net/http"
)

func GetProblemInfo(r *gin.Context) {
	var args dispatch.ProblemArgs
	if err := r.ShouldBindUri(&args); err != nil {
		r.AbortWithStatusJSON(http.StatusOK, gin.H{
			"status":  tools.ConnectOJError,
			"message": err.Error(),
		})
		return
	}
	resp, err := dispatch.GetProblemInfo(&args)
	if err != nil {
		r.AbortWithStatusJSON(http.StatusOK, gin.H{
			"status":  tools.ConnectOJError,
			"message": err.Error(),
		})
		return
	}

	r.JSON(http.StatusOK, gin.H{
		"status":  tools.OK,
		"message": resp,
	})
	return
}
