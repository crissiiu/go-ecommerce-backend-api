package user

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (us *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	//public router
	userRouterPulblic := Router.Group("/user")
	{
		userRouterPulblic.POST("/register")
		userRouterPulblic.POST("/otp")
	}

	//private router
	userRouterPrivate := Router.Group("/user")
	// userRouterPrivate.Use(Limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.GET("/get_info")
	}
}
