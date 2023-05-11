package controllers

import (
	"admin-api/config"
	"admin-api/helper"
	"admin-api/models"
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	db         *mongo.Client
	collection *mongo.Collection
}

var employeeValidate = validator.New()

// NewEmployeeConroller new handle
func NewEmployeeConroller(db *mongo.Client) *EmployeeConroller {
	return &EmployeeConroller{
		db:         db,
		collection: config.SelectCollection(db, "person"),
	}
}

// Login is for route employee
func (e *EmployeeConroller) Login(c *gin.Context) {
	var account = models.AdminLogin{}

	if err := c.BindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	if account.Username != os.Getenv("ADMIN_USERNAME") && account.Password != os.Getenv("ADMIN_PASSWORD") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password."})
		return
	}

	token, refreshToken, _ := helper.GenerateAllTokens(account.Username)

	c.JSON(http.StatusOK, gin.H{"response": gin.H{
		"access_token":         token,
		"refresh_access_token": refreshToken},
	})
}

// CreateEmployee is for route employee
func (e *EmployeeConroller) CreateEmployee(c *gin.Context) {
	var employee = models.EmployeeAccount{}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	if err := c.BindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	employee.ID = primitive.NewObjectID()
	employee.EmployeeID = primitive.NewObjectID().Hex()
	employee.CreateAt = time.Now()
	employee.UpdateAt = time.Now()

	validationErr := employeeValidate.Struct(employee)

	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		fmt.Println(validationErr)
		return
	}

	result, insertErr := e.db.Database("Employee").Collection("person").InsertOne(ctx, employee)
	fmt.Println(result)

	if insertErr != nil {
		msg := fmt.Sprintf("order item was not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		fmt.Println(result)
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": gin.H{"employee": employee.ID}})
	return
}

// GetAllEmployee is for route employee
func (e *EmployeeConroller) GetAllEmployee(c *gin.Context) {
	var employees []bson.M

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := e.collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	if err = result.All(ctx, &employees); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	defer cancel()
	c.JSON(http.StatusOK, gin.H{"response": employees})
	return
}

// GetEmployeeByID is for route employee
func (e *EmployeeConroller) GetEmployeeByID(c *gin.Context) {
	var employee bson.M

	employeeid := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := e.collection.Find(ctx, bson.M{"person": employeeid})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	if err = result.All(ctx, &employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, employee)
	return
}

// UpdateEmployee is for route employee
func (e *EmployeeConroller) UpdateEmployee(c *gin.Context) {
	var employee = models.EmployeeAccount{}

	employeeid := c.Param("id")
	eid, _ := primitive.ObjectIDFromHex(employeeid)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	if err := c.BindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	result, err := e.collection.UpdateOne(ctx, bson.M{"_id": eid},
		bson.D{bson.E{Key: "$set", Value: bson.D{
			bson.E{Key: "firstname", Value: employee.Firstname},
			bson.E{Key: "lastname", Value: employee.Lastname},
			bson.E{Key: "position", Value: employee.Position},
			bson.E{Key: "salary", Value: employee.Salary},
			bson.E{Key: "year_experience", Value: employee.Yearexperience},
			bson.E{Key: "update_at", Value: employee.Yearexperience},
		}}})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": "employee" + strconv.Itoa(int(result.ModifiedCount)) + "has update"})
	return
}

// RemoveEmployee is for route employee
func (e *EmployeeConroller) RemoveEmployee(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	employeeid := c.Param("id")
	mid, _ := primitive.ObjectIDFromHex(employeeid)
	result, err := e.collection.DeleteOne(ctx, bson.M{"_id": mid})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, result.DeletedCount)
	return
}
