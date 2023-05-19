package api

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/PigeonShogi/HR_system/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createEmployeeRequest struct {
	Code     string `json:"code" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
	Identity string `json:"identity" binding:"oneof=employee HR-Admin"`
}

type getEmployeeRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) createEmployee(ctx *gin.Context) {
	var identityId int32
	var req createEmployeeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if req.Identity == "employee" {
		identityId = getEmployeeFromIdentities(server).ID
	} else if req.Identity == "HR-Admin" {
		identityId = getHrAdminFromIdentities(server).ID
	}

	arg := db.CreateEmployeeParams{
		IdentityID: identityId,
		Code:       req.Code,
		FullName:   req.FullName,
	}

	employee, err := server.store.CreateEmployee(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, employee)
}

func (server *Server) getEmployee(ctx *gin.Context) {
	var req getEmployeeRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	employee, err := server.store.GetEmployee(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, employee)
}

func getEmployeeFromIdentities(server *Server) (employee db.Identity) {
	employee, err := server.store.GetEmployeeFromIdentities(context.Background())
	if err != nil {
		fmt.Printf("Cannot get employee from identities. Error: %v", err)
	}
	return employee
}

func getHrAdminFromIdentities(server *Server) (hrAdm db.Identity) {
	hrAdmin, err := server.store.GetHrAdminFromIdentities(context.Background())
	if err != nil {
		fmt.Printf("Cannot get HR-Admin from identities. Error: %v", err)
	}
	return hrAdmin
}

type listEmployeeRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listEmployee(ctx *gin.Context) {
	var req listEmployeeRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListEmployeesParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	employees, err := server.store.ListEmployees(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, employees)
}
