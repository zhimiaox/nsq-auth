package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type api struct {
}

func (a *api) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (a *api) Auth(c *gin.Context) {
	req := &AuthReq{}
	err := c.ShouldBindQuery(req)
	if err != nil {
		return
	}
	auth := GetStorage().Get(req.Secret)
	if auth == nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "NOT_AUTHORIZED",
		})
		return
	}
	c.JSON(http.StatusOK, AuthResp{
		TTL:            TTL,
		Identity:       Identity,
		IdentityURL:    IdentityURL,
		Authorizations: auth,
	})
}

func (a api) Refresh(c *gin.Context) {
	GetStorage().Refresh()
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
