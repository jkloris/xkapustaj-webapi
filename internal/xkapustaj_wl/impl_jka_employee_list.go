package xkapustaj_wl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
 // Get /api/employee-list/:ambulanceId/entries
 // Provides the ambulance waiting list 
func (this *implJkaEmployeeListAPI) GetEmployeeListEntries(ctx *gin.Context) {
 	ctx.AbortWithStatus(http.StatusNotImplemented)
}
