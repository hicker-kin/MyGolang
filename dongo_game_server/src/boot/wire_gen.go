// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package boot

import (
	"dongo_game_server/src/config"
	"dongo_game_server/src/grpc"
	"dongo_game_server/src/support"
	"dongo_game_server/src/web"
	"dongo_game_server/src/web/controller"
	"dongo_game_server/src/web/service"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitWeb() (*web.WebApp, error) {
	configConfig := config.DefaultConfig()
	userServiceClient := config.Grpc_DefaultUserService(configConfig)
	db := config.NewDatabase_Web(configConfig)
	baseHdl := &controller.BaseHdl{
		DB: db,
	}
	captchaHdl := &controller.CaptchaHdl{}
	jwtHdl := &controller.JWTHdl{}
	managerService := &service.ManagerService{
		DB: db,
	}
	managerHdl := &controller.ManagerHdl{
		Service: managerService,
	}
	projectService := &service.ProjectService{
		DB: db,
	}
	projectHdl := &controller.ProjectHdl{
		Service: projectService,
	}
	resourceHdl := &controller.ResourceHdl{
		DB: db,
	}
	rpcHdl := &controller.RpcHdl{
		UserService: userServiceClient,
	}
	socketService := &service.SocketService{
		DB: db,
	}
	socketHdl := &controller.SocketHdl{
		Service: socketService,
		Project: projectService,
	}
	emailConfig := config.DefaultEmailConfig(configConfig)
	toolHdl := &controller.ToolHdl{
		DB:    db,
		Email: emailConfig,
	}
	trackHdl := &controller.TrackHdl{
		DB: db,
	}
	webApp := &web.WebApp{
		Config:      configConfig,
		UserService: userServiceClient,
		Base:        baseHdl,
		Captcha:     captchaHdl,
		JWT:         jwtHdl,
		Manager:     managerHdl,
		Project:     projectHdl,
		Resource:    resourceHdl,
		RPC:         rpcHdl,
		Socket:      socketHdl,
		Tool:        toolHdl,
		Track:       trackHdl,
	}
	return webApp, nil
}

func InitGrpc() (*grpc.GrpcApp, error) {
	configConfig := config.DefaultConfig()
	db := config.NewDatabase_Grpc(configConfig)
	grpcConfig := config.DefaultGrpcConfig(configConfig)
	grpcApp := &grpc.GrpcApp{
		DB:              db,
		GrpcUserService: grpcConfig,
	}
	return grpcApp, nil
}

func InitSupport() (*support.SupportApp, error) {
	configConfig := config.DefaultConfig()
	userServiceClient := config.Grpc_DefaultUserService(configConfig)
	db := config.NewDatabase_Web(configConfig)
	supportApp := &support.SupportApp{
		Config:      configConfig,
		UserService: userServiceClient,
		DB:          db,
	}
	return supportApp, nil
}

// wire.go:

var configSet = wire.NewSet(config.DefaultConfig, config.DefaultEmailConfig, config.DefaultGrpcConfig, config.Grpc_DefaultUserService, config.DefaultMemory)

var webSet = wire.NewSet(wire.Struct(new(controller.BaseHdl), "*"), wire.Struct(new(controller.CaptchaHdl), "*"), wire.Struct(new(controller.JWTHdl), "*"), wire.Struct(new(controller.ManagerHdl), "*"), wire.Struct(new(controller.ProjectHdl), "*"), wire.Struct(new(controller.ResourceHdl), "*"), wire.Struct(new(controller.RpcHdl), "*"), wire.Struct(new(controller.SocketHdl), "*"), wire.Struct(new(controller.ToolHdl), "*"), wire.Struct(new(controller.TrackHdl), "*"), wire.Struct(new(service.ManagerService), "*"), wire.Struct(new(service.SocketService), "*"), wire.Struct(new(service.ProjectService), "*"))
