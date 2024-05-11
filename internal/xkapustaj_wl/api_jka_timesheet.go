/*
 * Waiting List Api
 *
 * Hospital Employee List management for Web-In-Cloud system
 *
 * API version: 1.0.0
 * Contact: kapusta.jergus@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

 package xkapustaj_wl

import (
   "net/http"

   "github.com/gin-gonic/gin"
)

type JkaTimesheetAPI interface {

   // internal registration of api routes
   addRoutes(routerGroup *gin.RouterGroup)

    // AddNewTimesheetEntry - Adds a new timesheet entry
   AddNewTimesheetEntry(ctx *gin.Context)

    // DeleteTimesheetEntry - Deletes specific timesheet
   DeleteTimesheetEntry(ctx *gin.Context)

    // UpdateEmployeeTimesheet - Updates specific entry
   UpdateEmployeeTimesheet(ctx *gin.Context)

 }

// partial implementation of JkaTimesheetAPI - all functions must be implemented in add on files
type implJkaTimesheetAPI struct {

}

func newJkaTimesheetAPI() JkaTimesheetAPI {
  return &implJkaTimesheetAPI{}
}

func (this *implJkaTimesheetAPI) addRoutes(routerGroup *gin.RouterGroup) {
  routerGroup.Handle( http.MethodPost, "/timesheet/:ambulanceId/employee/:employeeId", this.AddNewTimesheetEntry)
  routerGroup.Handle( http.MethodDelete, "/timesheet/:ambulanceId/:timesheetId", this.DeleteTimesheetEntry)
  routerGroup.Handle( http.MethodPut, "/timesheet/:ambulanceId/:timesheetId", this.UpdateEmployeeTimesheet)
}

// Copy following section to separate file, uncomment, and implement accordingly
// // AddNewTimesheetEntry - Adds a new timesheet entry
// func (this *implJkaTimesheetAPI) AddNewTimesheetEntry(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // DeleteTimesheetEntry - Deletes specific timesheet
// func (this *implJkaTimesheetAPI) DeleteTimesheetEntry(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // UpdateEmployeeTimesheet - Updates specific entry
// func (this *implJkaTimesheetAPI) UpdateEmployeeTimesheet(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//

