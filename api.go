package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nsq-auth/models"
)

type api struct {
}

func (a *api) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (a *api) Auth(c *gin.Context) {
	req := &models.AuthReq{}
	err := c.ShouldBindQuery(req)
	if err != nil {
		return
	}

	if len(resp.Authorizations) <= 0 {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "NOT_AUTHORIZED",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (a api) Refresh(c *gin.Context) {

}

// StartAPI 启动web服务
func StartAPI() {
	r := gin.Default()
	rImpl := api{}
	r.GET("/ping", rImpl.Ping)
	r.GET("/auth", rImpl.Auth)
	r.GET("/refresh", rImpl.Refresh)
	if err := r.Run(APIAddr); err != nil {
		panic(err)
	}
}
