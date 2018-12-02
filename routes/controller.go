package routes

import (
	"github.com/gin-gonic/gin"
)

type Controller struct{}

func NewController() *Controller { return &Controller{} }

type APIMessage struct{ Message string `json:"message"` }

func (e *APIMessage) Throw(c *gin.Context) { c.JSON(200, e) }

func NewAPIMessage(message string) *APIMessage { return &APIMessage{Message: message} }

type APIErrorMessage struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}

func NewAPIError(message string, error error) *APIErrorMessage { return &APIErrorMessage{Message: message, Error: error} }

func (e *APIErrorMessage) Throw(c *gin.Context, status int) { c.JSON(status, e) }
