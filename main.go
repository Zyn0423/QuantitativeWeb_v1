package main
//
import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)


func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./views/start", ".html"))

	app.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})

	//　使用网络地址启动服务
	app.Run(iris.Addr("localhost:8080"))
}



// ExampleController serves the "/", "/ping" and "/hello".
type ExampleController struct {}

// Get serves
// Method:   GET
// Resource: http://localhost:8080
func (c *ExampleController) Get() mvc.Result {
	return mvc.Response{
		ContentType: "text/html",
		Text:        "<h1>Welcome</h1>",
	}
}

GetPing serves
Method:   GET
Resource: http://localhost:8080/ping
func (c *ExampleController) GetPing() string {
	return "pong"
}

func (c *ExampleController) GetHello() interface{} {

	return map[string]string{"message": "Hello Iris!"}
}

func (c *ExampleController) GetInfo() interface{}  {
	return map[string]string{"name" :"zxn"}
}