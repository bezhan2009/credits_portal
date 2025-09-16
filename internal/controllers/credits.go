package controllers

import (
	"Credits/internal/app/models"
	"Credits/internal/app/service"
	"Credits/internal/app/service/validators"
	"Credits/internal/controllers/middlewares"
	"Credits/pkg/errs"
	"Credits/pkg/logger"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllCredits(c *gin.Context) {
	statusID := c.Query("status_id")
	isValid, month := validators.ValidateMonth(c)
	if !isValid {
		HandleError(c, errs.ErrInvalidMonth)
		return
	}

	isValid, year := validators.ValidateYear(c)
	if !isValid {
		HandleError(c, errs.ErrInvalidYear)
		return
	}

	afterID := c.Query("after")

	credit, err := service.GetAllCredits(month, year, afterID, statusID)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, credit)
}

func GetCreditById(c *gin.Context) {
	creditIdStr := c.Param("id")
	creditId, err := strconv.Atoi(creditIdStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	credit, err := service.GetCreditById(creditId)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, credit)
}

func CreateCredit(c *gin.Context) {
	var credit models.Credits

	if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
		logger.Error.Print("Parse multipart form error:", err.Error())
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	if err := c.ShouldBind(&credit); err != nil {
		logger.Error.Print("Should bind error:", err.Error())
		HandleError(c, errs.ErrValidationFailed)
		return
	}

	frontPath := c.GetString(middlewares.FrontSideOfThePassportPath)
	backPath := c.GetString(middlewares.BackSideOfThePassportPath)
	selfiePath := c.GetString(middlewares.SelfieWithPassportPath)

	credit.FrontSideOfThePassport = frontPath
	credit.BackSideOfThePassport = backPath
	credit.SelfieWithPassport = selfiePath

	credit.RequestCreator = "Мобильный банк"

	app, err := service.CreateCredit(&credit)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, app)
}

func UpdateCredit(c *gin.Context) {
	creditIdStr := c.Param("id")
	creditId, err := strconv.Atoi(creditIdStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	var credit models.Credits
	if err = c.ShouldBind(&credit); err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	credit.ID = uint(creditId)

	err = service.UpdateCredit(credit)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "credit updated successfully",
	})
}

func DeleteCredit(c *gin.Context) {
	creditIdStr := c.Param("id")
	creditId, err := strconv.Atoi(creditIdStr)
	if err != nil {
		HandleError(c, errs.ErrInvalidID)
		return
	}

	err = service.DeleteCredit(uint(creditId))
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "credit deleted successfully",
	})
}
