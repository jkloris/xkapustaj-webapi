package xkapustaj_wl

import (
  "net/http"
  "slices"
  "github.com/gin-gonic/gin"
)

// Get /api/employee-list/:ambulanceId/entries/:employeeId
// Provides the employee 
func (this *implJkaEmployeeAPI) GetEmployee(ctx *gin.Context) {
	updateHospitalFunc(ctx, func(c *gin.Context, hospital *Hospital) (*Hospital, interface{}, int) {
        entryId := ctx.Param("employeeId")

        if entryId == "" {
            return nil, gin.H{
                "status":  http.StatusBadRequest,
                "message": "Entry ID is required",
            }, http.StatusBadRequest
        }

        entryIndx := slices.IndexFunc(hospital.Employees, func(empl EmployeeListEntry) bool {
            return entryId == empl.Id
        })

        if entryIndx < 0 {
            return nil, gin.H{
                "status":  http.StatusNotFound,
                "message": "Entry not found",
            }, http.StatusNotFound
        }

        
        return nil, hospital.Employees[entryIndx], http.StatusOK
    })
}

