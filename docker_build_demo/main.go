package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// 项目配置
type AppConfig struct {
	AppName     string
	Port        int
	Description string
}

var conf AppConfig

// 初始化配置文件
func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("/etc/appname/")  // 查找配置文件所在路径
	viper.AddConfigPath("$HOME/.appname") // 多次调用AddConfigPath，可以添加多个搜索路径
	viper.AddConfigPath(".")              // 还可以在工作目录中搜索配置文件
	viper.AddConfigPath("./conf")         // 还可以在工作目录中搜索配置文件
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("read config failed: %v \r", err)
	}
	if err := viper.Unmarshal(&conf); err != nil {
		fmt.Println(err)
	}
}

// 使用 gin 创建一个 http 服务
func main() {
	r := gin.Default()
	r.Any("/", func(c *gin.Context) {
		message := strings.Join([]string{"Hello", conf.AppName, conf.Description}, " ")
		c.JSON(200, gin.H{
			"message": message,
		})
	})
	appPort := strings.Join([]string{":", strconv.Itoa(conf.Port)}, "")
	r.Run(appPort)
}

func main1() {
	// 1、创建路由
	router := gin.Default()

	// 2、绑定规则
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello World!")
		context.Abort()
	})

	// 启动监听
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
