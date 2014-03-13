package admin

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/session"
	_ "github.com/astaxie/session/providers/memory"
	"github.com/royburns/goblogs/models"
	"github.com/royburns/goblogs/utils"
	"time"
	//"fmt"
)

type AddBlogController struct {
	beego.Controller
}

//然后在init函数中初始化
func init() {
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	go globalSessions.GC()
}

func (this *AddBlogController) Prepare() {
	sess := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
	sess_uid := sess.Get("userid")
	sess_username := sess.Get("username")
	if sess_uid == nil {
		this.Ctx.Redirect(302, "/admin/login")
		return
	}
	this.Data["Username"] = sess_username
}

func (this *AddBlogController) Get() {
	this.Layout = "admin/layout.html"
	this.TplNames = "admin/addblog.tpl"
}

func (this *AddBlogController) Post() {
	this.Ctx.Request.ParseForm()
	title := this.Ctx.Request.Form.Get("title")
	content := this.Ctx.Request.Form.Get("content")
	//打印生成日志
	defer utils.Info("addblog: ", "title:"+title, "content:"+content)
	var data models.Blogs
	data.Title = title
	data.Content = content
	//获取系统当前时间
	now := time.Now().Format("2006-01-02 15:04:05")
	data.Created = now
	models.InsertBlogs(data)
	this.Ctx.Redirect(302, "/admin/index")
}
