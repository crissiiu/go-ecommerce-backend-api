package manage

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (ad *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {

	//public router
	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.POST("/login")
	}

	//private router
	adminRouterPrivate := Router.Group("/admin")
	// userRouterPrivate.Use(Limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		adminRouterPrivate.POST("/active_user")
	}
}
