package handlers

import (
	"net/http"
	"strconv"

	"unit-api/internal/models"
	"unit-api/internal/repository"

	"github.com/gin-gonic/gin"
)

func GetUnitsHandler(c *gin.Context) {

	status := c.Query("status")

	units, err := repository.GetUnits(status)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, units)
}

func GetUnitByIDHandler(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	unit, err := repository.GetUnitByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Unit not found",
		})
		return
	}

	c.JSON(http.StatusOK, unit)
}

func CreateUnitHandler(c *gin.Context) {
	var unit models.Unit

	if err := c.ShouldBindJSON(&unit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := repository.CreateUnit(&unit)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Unit created successfully",
	})
}
	

func UpdateUnitStatusHandler(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	var body struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Validate status value
	if body.Status != models.StatusAvailable &&
		body.Status != models.StatusOccupied &&
		body.Status != models.StatusCleaning &&
		body.Status != models.StatusMaintenance {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid status value",
		})
		return
}

	unit, err := repository.GetUnitByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if unit == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Unit not found",
		})
		return
	}

	// Occupied to Available must go via "Cleaning In Progress" or "Maintenance Needed"
	if unit.Status == models.StatusOccupied && body.Status == models.StatusAvailable {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot change unit status directly from Occupied to Available.",
		})
		return
	}

	err = repository.UpdateUnitStatus(id, body.Status)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Unit status updated successfully",
	})
}