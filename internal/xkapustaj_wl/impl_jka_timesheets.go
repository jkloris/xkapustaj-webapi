package xkapustaj_wl

import (
	"net/http"
  
	"github.com/gin-gonic/gin"
  )


// Get /api/timesheet/:ambulanceId/employee/:employeeId
// GetEmployeeTimesheet - Provides the timesheet
func (this *implJkaTimesheetsAPI) GetEmployeeTimesheet(ctx *gin.Context) {
	updateAmbulanceFunc(ctx, func(c *gin.Context,  hospital *Hospital) (*Hospital, interface{}, int) {
		employeeId := ctx.Param("employeeId")
		if employeeId == "" {
            return nil, gin.H{
                "status":  http.StatusBadRequest,
                "message": "Entry ID is required",
            }, http.StatusBadRequest
        }

        entryIndx := slices.IndexFunc(hospital.Timesheets, func(t Timesheet) bool {
            return employeeId == t.EmployeeId
        })

        if entryIndx < 0 {
            return nil, gin.H{
                "status":  http.StatusNotFound,
                "message": "Entry not found",
            }, http.StatusNotFound
        }

        // return nil ambulance - no need to update it in db
        return nil,hospital.Timesheets[entryIndx], http.StatusOK
    })
}

