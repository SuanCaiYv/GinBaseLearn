package src

import "github.com/gin-gonic/gin"

func wrapStr(context *gin.Context, str string) {
	context.String(200, "%s", str)
}

// CRUD 比较简单的RESTFul格式的请求
func CRUD() {
	router := gin.Default()
	router.GET("/isGet", func(context *gin.Context) {
		context.String(200, "%s", "ok")
	})
	router.POST("/isPost", func(context *gin.Context) {
		context.String(200, "%s", "ok")
	})
	router.DELETE("/isDelete", func(context *gin.Context) {
		context.String(200, "%s", "ok")
	})
	router.PUT("isPut", func(context *gin.Context) {
		context.String(200, "%s", "ok")
	})
	router.Run("127.0.0.1:8190")
}

func PathVariable() {
	router := gin.Default()
	// 路径参数
	router.GET("/param/:name", func(context *gin.Context) {
		wrapStr(context, "name is:" + context.Param("name"))
	})
	// 强匹配，优先于路径参数匹配，和书写顺序无关
	router.GET("/param/msl", func(context *gin.Context) {
		wrapStr(context, "just msl")
	})
	// 可为空匹配
	router.GET("/param/nullable/:name1/*name2", func(context *gin.Context) {
		wrapStr(context, "nullable name: " + context.Param("name1") + ", " + context.Param("name2"))
	})
	router.Run(":8190")
}

func GetAndPost() {
	r := gin.Default()
	r.GET("/get", func(context *gin.Context) {
		name := context.DefaultQuery("name", "msl")
		age := context.Query("age")
		wrapStr(context, "name: " + name + ", age: " + age)
	})
}
