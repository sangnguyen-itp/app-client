package main

import (
	"gRPC/client/middleware"
	proto "gRPC/client/proto"
	"gRPC/client/resource"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net/http"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	authClient := proto.NewAuthenticateClient(conn)
	appClient := proto.NewAppServiceClient(conn)
	g := gin.Default()

	g.POST("/login", func(context *gin.Context) {

		var request resource.RLoginRequest
		if err := context.ShouldBindJSON(&request); err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}

		req := &proto.LoginRequest{UserName: request.Username, Password: request.Password, Platform: request.Platform}
		res, err := authClient.Login(context, req)
		if err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}

		context.JSON(http.StatusOK, res)
		return
	})

	g.POST("/logout", func(context *gin.Context) {

		var request resource.RLogoutRequest
		if err := context.ShouldBindJSON(&request); err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}

		req := &proto.LogoutRequest{AccessToken: middleware.GetAuthorizationToken(context), Platform: request.Platform}
		res, err := authClient.Logout(context, req)
		if err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}

		context.JSON(http.StatusOK, res)
		return
	})

	g.POST("/app", func(context *gin.Context) {

		var request resource.RAppRequest
		if err := context.ShouldBindJSON(&request); err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}

		req := &proto.AppRequest{
			Name:          request.Name,
			LatestVersion: request.LatestVersion,
			RunningStatus: request.RunningStatus,
			Type:          request.Type}

		res, err := appClient.CreateApp(context, req)
		if err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}

		context.JSON(http.StatusOK, res)
		return
	})

	g.GET("/app/:id", func(context *gin.Context) {

		id := context.Param("id")

		req := &proto.GetRequest{Uuid: id}

		res, err := appClient.GetApp(context, req)
		if err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}

		context.JSON(http.StatusOK, res)
		return
	})

	g.POST("/app/:id/firewall", func(context *gin.Context) {

		id := context.Param("id")

		var request resource.RFirewallRequest
		if err := context.ShouldBindJSON(&request); err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if id != request.AppUUID {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}

		req := &proto.FirewallRequest{
			AppUuid: request.AppUUID,
			Host:    request.Host,
			Port:    request.Port,
		}

		res, err := appClient.CreateFirewall(context, req)
		if err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}

		context.JSON(http.StatusOK, res)
		return
	})

	g.GET("/app/:id/firewall", func(context *gin.Context) {

		id := context.Param("id")

		req := &proto.GetRequest{Uuid: id}

		res, err := appClient.GetFirewall(context, req)
		if err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}

		context.JSON(http.StatusOK, res)
		return
	})

	g.Run(":8002")
}
