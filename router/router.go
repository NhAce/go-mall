package router

import (
	"go-mall/controller"
	"go-mall/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 创建控制器实例
	userController := &controller.UserController{}
	productController := &controller.ProductController{}

	// API路由组
	api := r.Group("/api")
	{
		// 用户相关路由（无需认证）
		user := api.Group("/user")
		{
			user.POST("/login", userController.Login)
		}

		// 商品相关路由（无需认证）
		product := api.Group("/product")
		{
			product.GET("/list", productController.GetProductList)
			product.GET("/detail/:id", productController.GetProductDetail)
		}

		// 分类相关路由（无需认证）
		category := api.Group("/category")
		{
			category.GET("/list", productController.GetCategoryList)
		}

		// 需要认证的路由
		auth := api.Group("")
		auth.Use(middleware.AuthMiddleware())
		{
			// 用户相关路由（需要认证）
			auth.GET("/user/info", userController.GetUserInfo)
			auth.POST("/user/update", userController.UpdateUserInfo)

			// TODO: 添加购物车、订单、地址等需要认证的路由
		}
	}

	return r
} 