package response

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

type Responsefiles struct {
	Message       string `json:"message"`
	PersonalFiles any    `json:"personal_files,omitempty"`
	GroupFiles    any    `json:"group_files,omitempty"`
	Error         string `json:"error,omitempty"`
}

func ParseUintParam(c *gin.Context, paramName string) (uint64, error) {
	paramValue := c.Param(paramName)
	id, err := strconv.ParseUint(paramValue, 10, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func RespondWithError(c *gin.Context, status int, err error) {
	response := Response{
		Error: err.Error(),
	}
	c.JSON(status, response)
}

func RespondWithInvalid(c *gin.Context, status int, message string) {
	response := Response{
		Error: message,
	}
	c.JSON(status, response)
}

func RespondWithNotFound(c *gin.Context) {
	response := Response{
		Error: "Not Find",
	}
	c.JSON(404, response)
}

func RespondWithSuccess(c *gin.Context, message string, data any) {
	response := Response{
		Message: message,
		Data:    data,
	}
	c.JSON(200, response)
}

func RespondWithSuccessForFiles(c *gin.Context, message string, personalfiles any, groupfile any) {
	response := Responsefiles{
		Message:       message,
		PersonalFiles: personalfiles,
		GroupFiles:    groupfile,
	}
	c.JSON(200, response)
}

func RespondWithSuccessCreation(c *gin.Context, message string, data any) {
	response := Response{
		Message: message,
		Data:    data,
	}
	c.JSON(201, response)
}
