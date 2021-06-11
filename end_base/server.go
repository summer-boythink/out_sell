package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	serverlessplus "github.com/serverlessplus/go"
	"net"
	"net/http"
	apiController "out_sell/controller/api"
	"out_sell/exec_client"
	"github.com/tencentyun/scf-go-lib/cloudfunction"
)

const (
	port = 8080
)

var handler *serverlessplus.Handler

func init() {
	mysqlClient := exec_client.GetMysql()
	r := gin.Default()
	api := r.Group("/helloworld-1614945200")
	{
		food := api.Group("/food")
		{
			food.GET("/checktake", apiController.CheckTake(mysqlClient))
			food.GET("/changetake",apiController.ChangeTake(mysqlClient))
			food.GET("/putfood",apiController.PutFood(mysqlClient))
		}

		user := api.Group("/user")
		{
			user.POST("/login", apiController.CheckLogin(mysqlClient))
		}
	}

	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", serverlessplus.Host, port))
	if err != nil {
		fmt.Printf("failed to listen on port %d: %v\n", port, err)
		// panic to force the runtime to restart
		panic(err)
	}
	go http.Serve(l, r)

	// setup handler
	types := []string{"image/png"}
	handler = serverlessplus.NewHandler(port).WithBinaryMIMETypes(types)
}

func entry(ctx context.Context, req *serverlessplus.APIGatewayRequest) (*serverlessplus.APIGatewayResponse, error) {
	return handler.Handle(ctx, req)
}

func main() {
	cloudfunction.Start(entry)
}
