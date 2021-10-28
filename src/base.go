package src

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func wrapStr(context *gin.Context, str string) {
	context.String(200, "%s", str)
}

// CRUD 比较简单的RESTFul格式的请求
func CRUD() {
	router := gin.Default()
	// 每个请求一个Context
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
		wrapStr(context, "name is:"+context.Param("name"))
	})
	// 强匹配，优先于路径参数匹配，和书写顺序无关
	router.GET("/param/msl", func(context *gin.Context) {
		wrapStr(context, "just msl")
	})
	// 可为空匹配
	router.GET("/param/nullable/:name1/*name2", func(context *gin.Context) {
		wrapStr(context, "nullable name: "+context.Param("name1")+", "+context.Param("name2"))
	})
	router.Run(":8190")
}

func GetAndPost() {
	r := gin.Default()
	r.GET("/get", func(context *gin.Context) {
		// 进行参数查询，也可以设置缺省值
		name := context.DefaultQuery("name", "msl")
		age := context.Query("age")
		wrapStr(context, "name: "+name+", age: "+age)
	})
	r.POST("/post", func(context *gin.Context) {
		name := context.DefaultPostForm("name", "msl")
		age := context.PostForm("age")
		wrapStr(context, "name: "+name+", age: "+age)
	})
	// 当然，路径查询也可以和表单查询混合使用
	r.POST("/map", func(context *gin.Context) {
		// 进行map解析，要求查询参数符合map书写形式，比如：/map?ids[0]=1&ids[1]=2
		// 同时请求体：names[0]=msl;names[1]=cwb
		ids := context.QueryMap("ids")
		names := context.PostFormMap("names")
		context.JSON(200, gin.H{
			"ids":   ids,
			"names": names,
		})
	})
	r.Run(":8190")
}

func FileUpload() {
	r := gin.Default()
	// 限制文件存储使用的内存大小为8MB
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(context *gin.Context) {
		file, _ := context.FormFile("file")
		wrapStr(context, "get file: "+file.Filename+", size: "+fmt.Sprintf("%d", file.Size))
	})
	// 多文件上传
	r.POST("/uploads", func(context *gin.Context) {
		form, _ := context.MultipartForm()
		files := form.File["files"]
		stringBuilder := strings.Builder{}
		for _, file := range files {
			// 保存文件
			// context.SaveUploadedFile(file, "")
			stringBuilder.WriteString(file.Filename)
			stringBuilder.WriteString(", ")
		}
		wrapStr(context, stringBuilder.String())
	})
	r.Run(":8190")
}

func MiddleWare() {
	r := gin.New()
	r.GET("/test1", func(context *gin.Context) {
		wrapStr(context, "ok")
	})
	// 对所有/a开头的请求进行拦截
	auth := r.Group("/a")
	// 类似于添加请求拦截器
	auth.Use(func(context *gin.Context) {
		fmt.Println("need auth")
	})
	// 这个花括号就是为了美观
	// 在这里处理所有以/a为开头的请求
	{
		auth.POST("/signIn", func(context *gin.Context) {
			username := context.PostForm("username")
			password := context.PostForm("password")
			context.JSON(200, gin.H{
				"username": username,
				"password": password,
			})
		})
	}
	// 统一拦截和书写位置无关
	r.GET("/test2", func(context *gin.Context) {
		wrapStr(context, "ok")
	})
	r.Use(gin.CustomRecovery(func(context *gin.Context, err interface{}) {
		// 在这里编写panic处理逻辑
	}))
	r.Run(":8190")
}
