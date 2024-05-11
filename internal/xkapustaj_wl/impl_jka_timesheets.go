package xkapustaj_wl

import (
	"net/http"
  
	"github.com/gin-gonic/gin"
  )


// Get /api/timesheet/:ambulanceId/employee/:employeeId
// GetEmployeeTimesheet - Provides the timesheet
func (this *implJkaTimesheetsAPI) GetEmployeeTimesheet(ctx *gin.Context) {
 	ctx.AbortWithStatus(http.StatusNotImplemented)
}

