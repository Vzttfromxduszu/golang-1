package initialize

import (
	"fmt"

	"github.com/Vzttfromxduszu/golang-1.git/common/global"
	"github.com/Vzttfromxduszu/golang-1.git/controller"
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

func Router() {
	engine := gin.Default()
	engine.Static("/assets", "./assets")

	engine.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:         "templates/frontend",
		Extension:    ".html",
		Master:       "layouts/master",
		DisableCache: true,
	}) // HTMLRender设置模板引擎, gintemplate用于简化 Gin 的模板渲染过程,Root是模板文件路径,Extension是模板文件后缀,Master是模板文件的主题,DisableCache设置是否缓存模板

	// 前台路由
	engine.GET("/", controller.Index)

	mw := gintemplate.NewMiddleware(gintemplate.TemplateConfig{
		Root:         "templates/backend",
		Extension:    ".html",
		Master:       "layouts/master",
		DisableCache: true,
	})

	// 后台管理员前端接口
	web := engine.Group("/admin", mw) // Group 是 gin 的路由组方法，可以将相同前缀的路由分组，这里是将后台路由分组, mw是中间件,用于设置模板引擎,会应用于该路由组中的所有路由。
	{
		//index
		web.GET("/", controller.AdminIndex)
	}
	{
		//用户登录api
		web.GET("/channel/list", controller.ListChannel)
		web.GET("/channel/view", controller.ViewChannel)
		web.GET("/channel/delete", controller.DeleteChannel)
		web.POST("/channel/save", controller.SaveChannel)
	}

	// 启动服务
	post := fmt.Sprintf(":%s", global.Config.Server.Post)
	if err := engine.Run(post); err != nil {
		fmt.Printf("启动失败: %v", err)
	}

}
