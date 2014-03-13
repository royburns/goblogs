package admin

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/session"
	_ "github.com/astaxie/session/providers/memory"
	"github.com/royburns/goblogs/models"
	"io"
)

var globalSessions *session.Manager

type LoginController struct {
	beego.Controller
}

//然后在init函数中初始化
func init() {
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	go globalSessions.GC()
}

func (this *LoginController) Get() {
	this.TplNames = "admin/signin.tpl"
}

func (this *LoginController) Post() {
	//数据处理
	this.TplNames = "admin/signin.tpl"
	this.Ctx.Request.ParseForm()
	username := this.Ctx.Request.Form.Get("username")
	password := this.Ctx.Request.Form.Get("password")

	if username == "" {
		this.Data["UsernameNull"] = "username is not null"
		return
	}

	if password == "" {
		this.Data["PasswordNull"] = "password is not null"
		return
	}

	adminInfo := models.GetAdminInfo(username)

	if adminInfo.Username == "" {
		this.Data["UsernameNull"] = "用户不存在"
		return
	}

	md5Password := md5.New()
	io.WriteString(md5Password, password)
	buffer := bytes.NewBuffer(nil)
	fmt.Fprintf(buffer, "%x", md5Password.Sum(nil))
	newPass := buffer.String()

	if adminInfo.Password != newPass {
		this.Data["PasswordNull"] = "密码错误"
		return
	}

	//登录成功设置session
	sess := globalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
	sess.Set("userid", adminInfo.Id)
	sess.Set("username", adminInfo.Username)

	this.Ctx.Redirect(302, "/admin/index")
}
