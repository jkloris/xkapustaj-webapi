package xkapustaj_wl

import (
	"net/http"
	"github.com/gin-gonic/gin"
  )


// Get /api/timesheet/:ambulanceId/employee/:employeeId
// GetEmployeeTimesheet - Provides the timesheet
func (this *implJkaTimesheetsAPI) GetEmployeeTimesheet(ctx *gin.Context) {
	updateHospitalFunc(ctx, func(c *gin.Context,  hospital *Hospital) (*Hospital, interface{}, int) {
		employeeId := ctx.Param("employeeId")
		if employeeId == "" {
            return nil, gin.H{
                "status":  http.StatusBadRequest,
                "message": "Entry ID is required",
            }, http.StatusBadRequest
        }

		filteredTimesheets := []Timesheet{}
		for _, value := range hospital.Timesheets {
			 if value.EmployeeId == employeeId {
				 filteredTimesheets = append(filteredTimesheets, value)
			 }
		}

    

        // return nil ambulance - no need to update it in db
        return nil,filteredTimesheets, http.StatusOK
    })
}

