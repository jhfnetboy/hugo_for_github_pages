package main

// jhfnetboy的本地web服务，V2
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.DebugMode) //全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
	router := gin.Default()    //获得路由实例

	//添加中间件
	// router.Use(Middleware)
	//注册接口
	//router.GET("/simple/server/get", GetHandler)
	// router.POST("/simple/server/post", PostHandler)
	// router.PUT("/simple/server/put", PutHandler)
	// router.DELETE("/simple/server/delete", DeleteHandler)

	//加载全局模版
	// router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	// router.GET("/index", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.tmpl", gin.H{
	// 		"title": "焦会峰的首页",
	// 	})
	// }
	// contentTmp := ReadListFile("/list.log")

	//指定静态目录
	// router.Static("/tla", "./tla")
	router.Static("/", "./public")
	//监听端口
	http.ListenAndServe(":888", router)

}

// func ReadListFile(filename) {

// }

// func Middleware(c *gin.Context) {
// 	fmt.Println("this is a middleware!")
// }

// func GetHandler(c *gin.Context) {
// 	value, exist := c.GetQuery("key")
// 	if !exist {
// 		value = "the key is not exist!"
// 	}
// 	c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("get success! %s\n", value)))
// 	return
// }
// func PostHandler(c *gin.Context) {
// 	type JsonHolder struct {
// 		Id   int    `json:"id"`
// 		Name string `json:"name"`
// 	}
// 	holder := JsonHolder{Id: 1, Name: "my name"}
// 	//若返回json数据，可以直接使用gin封装好的JSON方法
// 	c.JSON(http.StatusOK, holder)
// 	return
// }
// func PutHandler(c *gin.Context) {
// 	c.Data(http.StatusOK, "text/plain", []byte("put success!\n"))
// 	return
// }
// func DeleteHandler(c *gin.Context) {
// 	c.Data(http.StatusOK, "text/plain", []byte("delete success!\n"))
// 	return
// }

// package main

// import (
// 	"net/http"
// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	router := gin.Default()
// 	router.LoadHTMLGlob("templates/*")
// 	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
// 	router.GET("/index", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "index.tmpl", gin.H{
// 			"title": "焦会峰的首页",
// 		})
// 	})
// 	router.Run() // listen and serve on 0.0.0.0:8080
// }
