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

type JkaTimesheetsAPI interface {

   // internal registration of api routes
   addRoutes(routerGroup *gin.RouterGroup)

    // GetEmployeeTimesheet - Provides the timesheet
   GetEmployeeTimesheet(ctx *gin.Context)

 }

// partial implementation of JkaTimesheetsAPI - all functions must be implemented in add on files
type implJkaTimesheetsAPI struct {

}

func newJkaTimesheetsAPI() JkaTimesheetsAPI {
  return &implJkaTimesheetsAPI{}
}

func (this *implJkaTimesheetsAPI) addRoutes(routerGroup *gin.RouterGroup) {
  routerGroup.Handle( http.MethodGet, "/timesheet/:ambulanceId/employee/:employeeId", this.GetEmployeeTimesheet)
}

// Copy following section to separate file, uncomment, and implement accordingly
// // GetEmployeeTimesheet - Provides the timesheet
// func (this *implJkaTimesheetsAPI) GetEmployeeTimesheet(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//

