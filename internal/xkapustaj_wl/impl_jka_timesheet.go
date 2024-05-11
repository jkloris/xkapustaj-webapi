package xkapustaj_wl

import (
  "net/http"

  "github.com/gin-gonic/gin"
)


//Post /api/timesheet/:ambulanceId/employee/:employeeId
func (this *implJkaTimesheetAPI) AddNewTimesheetEntry(ctx *gin.Context) {
 	ctx.AbortWithStatus(http.StatusNotImplemented)
}

// DeleteTimesheetEntry - Deletes specific timesheet
func (this *implJkaTimesheetAPI) DeleteTimesheetEntry(ctx *gin.Context) {
 	ctx.AbortWithStatus(http.StatusNotImplemented)
}

// UpdateEmployeeTimesheet - Updates specific entry
func (this *implJkaTimesheetAPI) UpdateEmployeeTimesheet(ctx *gin.Context) {
 	ctx.AbortWithStatus(http.StatusNotImplemented)
}


