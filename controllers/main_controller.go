package controllers

import (
	"Pureblog/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/mysql"
	"github.com/astaxie/beego/utils/captcha"
	"github.com/astaxie/beego/validation"
	"html/template"
	"strconv"
	"time"
)

var cpt *captcha.Captcha

//验证码模块
func init() {
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
	cpt.ChallengeNums = 4
	cpt.StdWidth = 100
	cpt.StdHeight = 40
}

type MainController struct {
	beego.Controller
}

func (c *MainController) Blog_index() {
	o := orm.NewOrm()
	qs := o.QueryTable("blog_article")
	var CategoryList []orm.ParamsList
	_, Cerror := qs.GroupBy("category").ValuesList(&CategoryList, "category")
	if Cerror == nil {
		c.Data["name"] = beego.AppConfig.String("blog_name")
		c.Data["second_name"] = beego.AppConfig.String("blog_second_name")
		c.Data["navitem_github_link"] = beego.AppConfig.String("navitem_github_link")
		c.Data["navitem_githubio_link"] = beego.AppConfig.String("navitem_githubio_link")
		c.Data["xsrf_token"] = c.XSRFToken()
		c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
		c.Data["Category"] = CategoryList
		c.TplName = "index.html"
	} else {
		logs.Error(Cerror.Error())
	}
}

func (c *MainController) Blog_category() {
	cate := c.Ctx.Input.Param(":cate")
	if FilteredSQLInject(cate) {
		c.Redirect("/", 302)
	}
	o := orm.NewOrm()
	qs := o.QueryTable("blog_article")
	var CategoryList []orm.ParamsList
	var Postlist []models.BlogArticle
	_, Cerror := qs.GroupBy("category").ValuesList(&CategoryList, "category")
	_, Perror := qs.Filter("category__iexact", cate).All(&Postlist)
	if Cerror == nil && Perror == nil {
		c.Data["name"] = beego.AppConfig.String("blog_name")
		c.Data["second_name"] = beego.AppConfig.String("blog_second_name")
		c.Data["navitem_github_link"] = beego.AppConfig.String("navitem_github_link")
		c.Data["navitem_githubio_link"] = beego.AppConfig.String("navitem_githubio_link")
		c.Data["xsrf_token"] = c.XSRFToken()
		c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
		c.Data["Category"] = CategoryList
		c.Data["Postlist"] = &Postlist
		c.TplName = "category.html"
	} else {
		logs.Error(Cerror.Error())
		logs.Error(Perror.Error())
	}
}

func (c *MainController) Blog_aboutme() {
	c.TplName = "aboutme.html"
}

func (c *MainController) Blog_article() {
	o := orm.NewOrm()
	blogid := c.Ctx.Input.Param(":id")
	if FilteredSQLInject(blogid) {
		c.Redirect("/", 302)
	} else {
		article := models.BlogArticle{}
		article.Id, _ = strconv.ParseInt(blogid, 10, 64)
		var CategoryList []orm.ParamsList
		Oerror := o.Read(&article)
		qs := o.QueryTable("blog_article")
		_, Cerror := qs.GroupBy("category").ValuesList(&CategoryList, "category")
		if Oerror != nil && Cerror != nil {
			logs.Error(Cerror.Error())
			logs.Error(Oerror.Error())
			c.Data["Errormes"] = "SORRY，没有该文章！！！"
			c.TplName = "502.html"
		} else {
			c.Data["title"] = article.Title
			c.Data["author"] = article.Author
			c.Data["subtitle"] = "--write by " + article.Author + " " + article.Addtime.String()
			c.Data["content"] = article.Content.Value()
			c.Data["Category"] = CategoryList
			c.TplName = "article.html"
		}
	}
}

func (c *MainController) Blog_getarticle_json() {
	if id := c.Ctx.Input.Param(":id"); id != "" {
		o := orm.NewOrm()
		qs := o.QueryTable("blog_article")
		if id == "0" {
			var ArticleList []orm.ParamsList
			_, Aerror := qs.Filter("id__isnull", false).OrderBy("-top").ValuesList(&ArticleList, "id", "author", "top", "title", "morecontent", "category", "addtime")
			if Aerror == nil {
				c.Data["json"] = &ArticleList
				c.ServeJSON()
			}else {
				logs.Error(Aerror.Error())
			}
		} else {
			var ArticleList []orm.ParamsList
			sid, _ := strconv.Atoi(id)
			_, Aerror := qs.Filter("id", sid).ValuesList(&ArticleList)
			if Aerror == nil {
				c.Data["json"] = &ArticleList
				c.ServeJSON()
			}else {
				logs.Error(Aerror.Error())
			}
		}
	}
}

func (c *MainController) Blgo_articleupdate() {
	csession := c.GetSession("Pureblog")
	if csession == nil {
		c.Redirect("/login", 302)
	} else {
		uparticle := models.BlogArticle{}
		Cerror := c.ParseForm(&uparticle)
		if Cerror != nil {
			logs.Error(Cerror)
			c.Data["Errormes"] = "您输入的信息有误请重新输入"
			c.TplName = "502.html"
		} else {
			if len(uparticle.Content) > 200 {
				uparticle.Morecontent = uparticle.Content[0:200]
			} else {
				uparticle.Morecontent = uparticle.Content
			}
			o := orm.NewOrm()
			_, Oerror := o.Update(&uparticle, "title", "content", "category", "top", "morecontent")
			if Oerror != nil {
				logs.Error(Oerror.Error())
				c.Data["Errormes"] = "文章更新失败!"
				c.TplName = "502.html"
			} else {
				c.Redirect("/admin", 302)
			}
		}
	}
}

func (c *MainController) Blgo_articledelete() {
	csession := c.GetSession("Pureblog")
	if csession == nil {
		c.Redirect("/login", 302)
	} else {
		delarticle := models.BlogArticle{}
		Cerror := c.ParseForm(&delarticle)
		if Cerror != nil {
			logs.Error(Cerror.Error())
			c.Data["Errormes"] = "您输入的信息有误请重新输入"
			c.TplName = "502.html"
		} else {
			o := orm.NewOrm()
			_, Oerror := o.Delete(&delarticle, "id")
			if Oerror != nil {
				logs.Error(Oerror.Error())
				c.Data["Errormes"] = "文章删除失败!"
				c.TplName = "502.html"
			} else {
				c.Redirect("/admin", 302)
			}
		}
	}
}

func (c *MainController) Blgo_articleadd_post() {
	csession := c.GetSession("Pureblog")
	if csession == nil {
		c.Redirect("/login", 302)
	} else {
		addarticle := models.BlogArticle{}
		Cerror := c.ParseForm(&addarticle)
		username := c.GetSession("Pureblog")
		addarticle.Author = username.(string)
		addarticle.Addtime = orm.DateTimeField(time.Now())
		if len(addarticle.Content) > 200 {
			addarticle.Morecontent = addarticle.Content[0:200]
		} else {
			addarticle.Morecontent = addarticle.Content
		}
		if Cerror != nil {
			logs.Error(Cerror.Error())
			c.Data["Errormes"] = "您输入的信息有误请重新输入"
			c.TplName = "502.html"
		} else {
			o := orm.NewOrm()
			_, Oerror := o.Insert(&addarticle)
			if Oerror != nil {
				logs.Error(Oerror.Error())
				c.Data["Errormes"] = "文章添加失败!"
				c.TplName = "502.html"
			} else {
				c.Redirect("/admin", 302)
			}
		}
	}
}

func (c *MainController) Blgo_articleadd() {
	csession := c.GetSession("Pureblog")
	if csession == nil {
		c.Redirect("/login", 302)
	} else {
		c.Data["xsrf_token"] = c.XSRFToken()
		c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
		c.TplName = "addarticle.html"
	}
}

func (c *MainController) Blog_admin() {
	csession := c.GetSession("Pureblog")
	if csession == nil {
		c.Redirect("/login", 302)
	} else {
		o := orm.NewOrm()
		qs := o.QueryTable("blog_article")
		var CategoryList []orm.ParamsList
		_, Cerror := qs.GroupBy("category").ValuesList(&CategoryList, "category")
		if Cerror == nil {
			c.Data["Category"] = CategoryList
			c.Data["xsrf_token"] = c.XSRFToken()
			c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
		}else {
			logs.Error(Cerror.Error())
		}
		c.TplName = "admin.html"
	}
}

func (c *MainController) Blog_login() {
	csession := c.GetSession("Pureblog")
	if csession == nil {
		c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
		c.Data["xsrf_token"] = c.XSRFToken()
		c.TplName = "login.html"
	} else {
		c.Redirect("/admin", 302)
	}
}

func (c *MainController) Blog_checkpwd() {
	//获得form表单的数据结构
	type Check_User struct {
		User string `form:"User"`
		Pwd  string `form:"Pwd"`
		Key  string `form:"captcha"`
	}
	cuser := Check_User{}
	Cerror := c.ParseForm(&cuser)    //将form表单数据储存到cuser中
	valid := validation.Validation{} //表单验证
	valid.MaxSize(cuser.User, 10, "User")
	valid.Length(cuser.Pwd, 32, "Pwd")
	o := orm.NewOrm()                                                  //获得orm
	checkpass := models.Admin{Username: cuser.User, Pappwd: cuser.Pwd} //初始化结构，用于账号和密码校验
	Oerror := o.Read(&checkpass)                                       //进行用户查询

	if !cpt.VerifyReq(c.Ctx.Request) || Oerror != nil || Cerror != nil || valid.HasErrors() || FilteredSQLInject(cuser.User+cuser.Pwd+cuser.Key) {
		c.Data["Errormes"] = "您输入的信息有误请重新输入"
		c.TplName = "502.html"
		logs.Error("表单验证：")
		logs.Error(valid.Errors)
		logs.Error("用户密码校验：")
		logs.Error(Oerror.Error())
	} else {
		csession := c.GetSession("Pureblog")
		if csession == nil {
			c.SetSession("Pureblog", cuser.User)
		}
		c.Redirect("/admin", 302)
	}
}
