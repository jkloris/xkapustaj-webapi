package xkapustaj_wl

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

// Get /api/employee-list/:ambulanceId/entries/:employeeId
// Provides the employee 
func (this *implJkaEmployeeAPI) GetEmployee(ctx *gin.Context) {
 	ctx.AbortWithStatus(http.StatusNotImplemented)
}

