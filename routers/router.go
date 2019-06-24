package routers

import (
	"Pureblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Blog_index")
	beego.Router("/category/:cate*.*", &controllers.MainController{}, "get:Blog_category")
	beego.Router("/aboutme", &controllers.MainController{}, "get:Blog_aboutme")
	beego.Router("/article", &controllers.MainController{}, "get:Blog_article")
	beego.Router("/article/update", &controllers.MainController{}, "post:Blgo_articleupdate")
	beego.Router("/article/delete", &controllers.MainController{}, "post:Blgo_articledelete")
	beego.Router("/article/add", &controllers.MainController{}, "post:Blgo_articleadd_post")
	beego.Router("/article/add", &controllers.MainController{}, "get:Blgo_articleadd")
	beego.Router("/article/:id:int", &controllers.MainController{}, "get:Blog_article")
	beego.Router("/article/:id", &controllers.MainController{}, "post:Blog_getarticle_json")
	beego.Router("/admin", &controllers.MainController{}, "get:Blog_admin")
	beego.Router("/login", &controllers.MainController{}, "get:Blog_login")
	beego.Router("/login", &controllers.MainController{}, "post:Blog_checkpwd")
}
