package api

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"                  // Swagger 静态文件
	ginSwagger "github.com/swaggo/gin-swagger" // Swagger 路由
	"gorm.io/gorm"
	"your_project/controller"
	"your_project/mapper"
	"your_project/service"
)

func InitRouter(db *gorm.DB) *gin.Engine {
	// 实例化服务和控制器
	userMapper := &mapper.UserMapper{}
	listMapper := &mapper.ListMapper{}
	groupMapper := &mapper.GroupMapper{}
	emailMapper := &mapper.EmailMapper{}

	loginService := &service.LoginService{UserMapper: userMapper}
	listService := &service.ListService{ListMapper: listMapper, UserMapper: userMapper}
	groupService := &service.GroupService{GroupMapper: groupMapper, EmailMapper: emailMapper}
	emailService := &service.EmailService{EmailMapper: emailMapper}

	loginController := &controller.LoginController{LoginService: loginService}
	listController := &controller.ListController{ListService: listService}
	groupController := &controller.GroupController{GroupService: groupService}
	emailController := &controller.EmailController{EmailService: emailService}
	// 创建 Gin 路由
	r := gin.Default()

	// 将数据库连接绑定到上下文
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// 定义路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiGroup := r.Group("/login")
	{
		apiGroup.POST("/saveUser", loginController.SaveUser) //注册
		apiGroup.POST("/log", loginController.Login)         // 登录
	}
	// 好友管理路由
	listGroup := r.Group("/list")
	{
		listGroup.POST("/add", listController.InsertFriend)                // 添加好友
		listGroup.POST("/updateRemark", listController.UpdateRemark)       // 更改备注
		listGroup.POST("/setRelationship", listController.SetRelationship) // 设置亲密关系
		listGroup.POST("/deleteFriend", listController.DeleteFriend)       // 假删除好友
		listGroup.GET("/checkWhoAdd", listController.CheckWhoAdd)          // 查看好友申请
		listGroup.POST("/agreeFriendAdd", listController.AgreeFriendAdd)   // 同意好友申请
		listGroup.GET("/showAllFriend", listController.ShowAllFriend)      // 查看好友列表
	}
	// 用户群组管理路由
	groupGroup := r.Group("/group")
	{
		groupGroup.POST("/insertGroup", groupController.InsertGroup)       // 创建群组
		groupGroup.POST("/sendEmailToAll", groupController.SendEmailToAll) // 群发消息
	}
	emailGroup := r.Group("/email")
	{
		emailGroup.POST("/sendEmail", emailController.SendEmail)      // 发送邮件
		emailGroup.GET("/receiveEmail", emailController.ReceiveEmail) // 接收邮件
		emailGroup.GET("/selectEmail", emailController.SelectEmail)   // 模糊查询邮件
		emailGroup.POST("/driftEmail", emailController.DriftEmail)    // 漂流瓶功能
		emailGroup.POST("/timeEmail", emailController.TimeEmail)      // 定时发送邮件
		emailGroup.POST("/deleteEmail", emailController.DeleteEmail)  // 删除邮件
	}

	return r
}
