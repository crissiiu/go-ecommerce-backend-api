package user

import "github.com/gin-gonic/gin"

type ProductRouter struct{}

func (pr *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	//public router
	productRouterPulblic := Router.Group("/product")
	{
		productRouterPulblic.GET("/search")
		productRouterPulblic.GET("/detail/:id")
	}

	//private router

}
