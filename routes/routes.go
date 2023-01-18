package routes

import (
	"3/api"
	"3/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	// 创建基于cookie的存储引擎，something-very-secret参数是用于加密的密钥
	store := cookie.NewStore([]byte("something-very-secret"))
	// 设置session中间件，参数my_session，指的是session的名字，也是cookie的名字
	r.Use(sessions.Sessions("my_session", store))
	v1 := r.Group("/api/v1")
	{
		//用户注册
		v1.POST("/register", api.UserRegister)
		//用户登录
		v1.POST("/login", api.UserLogin)
		authed := v1.Group("/")
		//设置中间件
		authed.Use(middleware.JWT())
		{
			//添加一条代办事项
			authed.POST("/task", api.CreateTask)
			//查看一条事项
			authed.GET("/task/:id", api.ShowTask)
			//查看所有 已完成/未完成/所有 事项
			authed.GET("/tasks", api.ListTask)
			//修改所有待办事项状态
			authed.PUT("/tasks", api.UpdateAllTask)
			//修改一条待办事项状态
			authed.PUT("/task/:id", api.UpdateTask)
			//输入关键词查询事项
			authed.GET("/tasks/search", api.SearchTask)
			//删除一条事项
			authed.DELETE("/task/:id", api.DeleteTask)
			//删除 所有已经完成/所有待办/所有 事项
			authed.DELETE("/tasks", api.DeleteAllTask)
		}
	}
	return r
}
