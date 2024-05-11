package xkapustaj_wl

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/google/uuid"
  "github.com/jkloris/xkapustaj-webapi/internal/db_service"
)


// CreateHospital - Saves new hospital definition
func (this *implHospitalsAPI) CreateHospital(ctx *gin.Context) {
  value, exists := ctx.Get("db_service")
  if !exists {
      ctx.JSON(
          http.StatusInternalServerError,
          gin.H{
              "status":  "Internal Server Error",
              "message": "db not found",
              "error":   "db not found",
          })
      return
  }

  db, ok := value.(db_service.DbService[Hospital])
  if !ok {
      ctx.JSON(
          http.StatusInternalServerError,
          gin.H{
              "status":  "Internal Server Error",
              "message": "db context is not of required type",
              "error":   "cannot cast db context to db_service.DbService",
          })
      return
  }

  hospital := Hospital{}
  err := ctx.BindJSON(&hospital)
  if err != nil {
      ctx.JSON(
          http.StatusBadRequest,
          gin.H{
              "status":  "Bad Request",
              "message": "Invalid request body",
              "error":   err.Error(),
          })
      return
  }

  if hospital.Id == "" {
    hospital.Id = uuid.New().String()
  }

  err = db.CreateDocument(ctx, hospital.Id, &hospital)

  switch err {
  case nil:
      ctx.JSON(
          http.StatusCreated,
          hospital,
      )
  case db_service.ErrConflict:
      ctx.JSON(
          http.StatusConflict,
          gin.H{
              "status":  "Conflict",
              "message": "hospital already exists",
              "error":   err.Error(),
          },
      )
  default:
      ctx.JSON(
          http.StatusBadGateway,
          gin.H{
              "status":  "Bad Gateway",
              "message": "Failed to create hospital in database",
              "error":   err.Error(),
          },
      )
  }
}

// DeleteHospital - Deletes specific hospital
func (this *implHospitalsAPI) DeleteHospital(ctx *gin.Context) {
  value, exists := ctx.Get("db_service")
  if !exists {
      ctx.JSON(
          http.StatusInternalServerError,
          gin.H{
              "status":  "Internal Server Error",
              "message": "db_service not found",
              "error":   "db_service not found",
          })
      return
  }

  db, ok := value.(db_service.DbService[Hospital])
  if !ok {
      ctx.JSON(
          http.StatusInternalServerError,
          gin.H{
              "status":  "Internal Server Error",
              "message": "db_service context is not of type db_service.DbService",
              "error":   "cannot cast db_service context to db_service.DbService",
          })
      return
  }

  ambulanceId := ctx.Param("ambulanceId")
  err := db.DeleteDocument(ctx, ambulanceId)

  switch err {
  case nil:
      ctx.AbortWithStatus(http.StatusNoContent)
  case db_service.ErrNotFound:
      ctx.JSON(
          http.StatusNotFound,
          gin.H{
              "status":  "Not Found",
              "message": "Hospital not found",
              "error":   err.Error(),
          },
      )
  default:
      ctx.JSON(
          http.StatusBadGateway,
          gin.H{
              "status":  "Bad Gateway",
              "message": "Failed to delete hospital from database",
              "error":   err.Error(),
          })
  }
}


