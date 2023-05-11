package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// EmployeeFunc is main repo for employee
type EmployeeFunc interface {
	CreateEmployee(c *gin.Context)
	GetAllEmployee(c *gin.Context)
	GetEmployeeById(c *gin.Context)
	UpdateEmployee(c *gin.Context)
	RemoveEmployee(c *gin.Context)
}

// EmployeeConroller is main controller
type EmployeeConroller struct {
	db *mongo.Client
}

// NewEmployeeConroller new handle
func NewEmployeeConroller(db *mongo.Client) *EmployeeConroller {
	return &EmployeeConroller{
		db: db,
	}
}

// Login is for route employee
func (e *EmployeeConroller) Login(c *gin.Context) {

}

// CreateEmployee is for route employee
func (e *EmployeeConroller) CreateEmployee(c *gin.Context) {

}

// GetAllEmployee is for route employee
func (e *EmployeeConroller) GetAllEmployee(c *gin.Context) {

}

// GetEmployeeByID is for route employee
func (e *EmployeeConroller) GetEmployeeByID(c *gin.Context) {

}

// UpdateEmployee is for route employee
func (e *EmployeeConroller) UpdateEmployee(c *gin.Context) {

}

// RemoveEmployee is for route employee
func (e *EmployeeConroller) RemoveEmployee(c *gin.Context) {

}
