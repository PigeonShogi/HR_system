package api

import (
	"context"
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

func (server *Server) createEmployee(ctx *gin.Context) {
	var identityId int32
	var req createEmployeeRequest
	fmt.Printf("請求%v", req)
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
