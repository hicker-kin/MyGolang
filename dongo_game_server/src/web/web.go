package web

import (
	"dongo_game_server/service/inf"
	"dongo_game_server/src/config"
	"dongo_game_server/src/database"
	"dongo_game_server/src/util"
	"dongo_game_server/src/web/controller"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type WebApp struct {
	Config      *config.Config
	UserService inf.UserServiceClient
	DB          *database.DB

	Base     *controller.BaseHdl
	Captcha  *controller.CaptchaHdl
	JWT      *controller.JWTHdl
	Manager  *controller.ManagerHdl
	Project  *controller.ProjectHdl
	Resource *controller.ResourceHdl
	RPC      *controller.RpcHdl
	Socket   *controller.SocketHdl
	Tool     *controller.ToolHdl
	Track    *controller.TrackHdl
}

func (p *WebApp) Start() {
	router := gin.New()
	router.Use(ServeRecover)
	// router.LoadHTMLGlob("./resources")
	routerGroup := router.Group("")
	p.Mount(routerGroup)

	err := router.Run(fmt.Sprintf(`:%s`, p.Config.Web.Addr))
	if err != nil {
		log.Fatalln(err)
		util.Chk(err)
	}

	log.Println("web serve running")
}

func (p *WebApp) Mount(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/version", p.Base.GetVersion())
	routerGroup.GET("/grpc/user/:user_id", p.RPC.GetGrpcUser)

	captcha := routerGroup.Group("/captcha")
	{
		captcha.GET("", p.Captcha.GetCaptcha)
		captcha.GET("/image/:captchaId", p.Captcha.GetCaptchaImg)
		captcha.POST("/verify/:captchaId/:value", p.Captcha.VerifyCaptcha)
	}

	routerGroup.GET("/email", p.Tool.SendEmail)
}
