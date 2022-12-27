package app

import (
	"desi-stonks/pkg/aws"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router        *gin.Engine
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run() error {
	r := gin.Default()

  r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"data": "hello world"})    
  })

	r.GET("/get-secret", func(c *gin.Context) {
		secretName := c.DefaultQuery("secretname", "")
		secretVal, err := aws.GetSecretValue(secretName);
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": secretVal})    
		}
  })

  r.Run()

	return nil
}