package xkapustaj_wl

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/google/uuid"
  "slices"
)


//Post /api/timesheet/:ambulanceId/employee/:employeeId
func (this *implJkaTimesheetAPI) AddNewTimesheetEntry(ctx *gin.Context) {
	updateHospitalFunc(ctx, func(c *gin.Context, hospital *Hospital) (*Hospital,  interface{},  int){
        var timesheet Timesheet

        if err := c.ShouldBindJSON(&timesheet); err != nil {
            return nil, gin.H{
                "status": http.StatusBadRequest,
                "message": "Invalid request body",
                "error": err.Error(),
            }, http.StatusBadRequest
        }

        if timesheet.EmployeeId == "" {
            return nil, gin.H{
                "status": http.StatusBadRequest,
                "message": "Employee ID is required",
            }, http.StatusBadRequest
        }

        timesheet.Id = uuid.NewString()

		hospital.Timesheets = append(hospital.Timesheets, timesheet)

		entryIndx := slices.IndexFunc( hospital.Timesheets, func(t Timesheet) bool {
            return timesheet.Id == t.Id 
        })

        if entryIndx < 0 {
            return nil, gin.H{
                "status": http.StatusInternalServerError,
                "message": "Failed to save entry",
            }, http.StatusInternalServerError
        }
        return hospital, hospital.Timesheets[entryIndx], http.StatusOK
    })
}

// DeleteTimesheetEntry - Deletes specific timesheet "/timesheet/{ambulanceId}/{timesheetId}"
func (this *implJkaTimesheetAPI) DeleteTimesheetEntry(ctx *gin.Context) {
	updateHospitalFunc(ctx, func(c *gin.Context, hospital *Hospital) (*Hospital,  interface{},  int){
		entryId := ctx.Param("timesheetId")

		if entryId == "" {
            return nil, gin.H{
                "status":  http.StatusBadRequest,
                "message": "Entry ID is required",
            }, http.StatusBadRequest
        }

		entryIndx := slices.IndexFunc( hospital.Timesheets, func(t Timesheet) bool {
			return entryId == t.Id
		})

		if entryIndx < 0{
			return nil, gin.H{
				"status":  http.StatusNotFound,
				"message": "Entry not found",
			}, http.StatusNotFound
		}

		hospital.Timesheets = slices.Delete(hospital.Timesheets, entryIndx, entryIndx+1)
		return hospital, nil, http.StatusNoContent
	})
}

// UpdateEmployeeTimesheet - Updates specific entry "/timesheet/{ambulanceId}/{timesheetId}"
func (this *implJkaTimesheetAPI) UpdateEmployeeTimesheet(ctx *gin.Context) {
	updateHospitalFunc(ctx, func(c *gin.Context, hospital *Hospital) (*Hospital,  interface{},  int){
		var timesheet Timesheet

        if err := c.ShouldBindJSON(&timesheet); err != nil {
            return nil, gin.H{
                "status":  http.StatusBadRequest,
                "message": "Invalid request body",
                "error":   err.Error(),
            }, http.StatusBadRequest
        }

		entryId := ctx.Param("timesheetId")
		if entryId == "" {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Entry ID is required",
			}, http.StatusBadRequest
		}

		entryIndx := slices.IndexFunc( hospital.Timesheets, func(t Timesheet) bool {
			return entryId == t.Id
		})

		if entryIndx < 0{
			return nil, gin.H{
				"status":  http.StatusNotFound,
				"message": "Entry not found",
			}, http.StatusNotFound
		}

		
		if timesheet.Description != "" {
			hospital.Timesheets[entryIndx].Description = timesheet.Description
		}
		
		if timesheet.Hours > 0 {
			hospital.Timesheets[entryIndx].Hours = timesheet.Hours
		}
		hospital.Timesheets[entryIndx].Date = timesheet.Date
		return hospital, hospital.Timesheets[entryIndx], http.StatusOK
	})
}


