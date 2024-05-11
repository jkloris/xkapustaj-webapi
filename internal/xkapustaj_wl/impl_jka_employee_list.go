package xkapustaj_wl

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
    "slices"
)
 // Get /api/employee-list/:ambulanceId/entries
 // Provides the ambulance waiting list 
func (this *implJkaEmployeeListAPI) GetEmployeeListEntries(ctx *gin.Context) {
	updateAmbulanceFunc(ctx, func(c *gin.Context,  hospital *Hospital) (*Hospital, interface{}, int) {
        result := hospital.Employees
        if result == nil {
            result = []EmployeeListEntry{}
        }
        return nil, result, http.StatusOK
    })
}

// // AddNewEmployee - Add  new employee
func (this *implJkaEmployeeListAPI) AddNewEmployee(ctx *gin.Context) {
	updateHospitalFunc(ctx, func(c *gin.Context, hospital *Hospital) (*Hospital,  interface{},  int){
        var entry EmployeeListEntry

        if err := c.ShouldBindJSON(&entry); err != nil {
            return nil, gin.H{
                "status": http.StatusBadRequest,
                "message": "Invalid request body",
                "error": err.Error(),
            }, http.StatusBadRequest
        }

        if entry.PatientId == "" {
            return nil, gin.H{
                "status": http.StatusBadRequest,
                "message": "Patient ID is required",
            }, http.StatusBadRequest
        }

        if entry.Id == "" || entry.Id == "@new" {
            entry.Id = uuid.NewString()
        }

        conflictIndx := slices.IndexFunc( hospital.Employees, func(empl EmployeeListEntry) bool {
            return entry.Id == empl.Id || entry.PatientId == empl.PatientId
        })

        if conflictIndx >= 0 {
            return nil, gin.H{
                "status": http.StatusConflict,
                "message": "Entry already exists",
            }, http.StatusConflict
        }

        hospital.Employees = append(hospital.Employees, entry)
        // hospital.reconcileWaitingList()
        // entry was copied by value return reconciled value from the list
        entryIndx := slices.IndexFunc( hospital.Employees, func(empl EmployeeListEntry) bool {
            return entry.Id == empl.Id
        })
        if entryIndx < 0 {
            return nil, gin.H{
                "status": http.StatusInternalServerError,
                "message": "Failed to save entry",
            }, http.StatusInternalServerError
        }
        return hospital, hospital.Employees[entryIndx], http.StatusOK
    })
}

