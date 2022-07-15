package controllers

import (
	"demoProject/models"
	"strings"

	// "github.com/astaxie/beego"
	"github.com/beego/beego/v2/server/web"
)

type ContactusController struct {
	web.Controller
}

// 預設表單頁面
func (this *ContactusController) Get() {
	this.TplName = "contactusTemplate.html"

	result := models.GetUserData()

	if len(result) <= 0 {
		this.Data["error"] = "查無資料"
		return
	}
	this.Data["data"] = result
}

// 感謝頁面
func (this *ContactusController) Thankyou() {
	this.TplName = "thankyouTemplate.html"
}

// 新增user
func (this *ContactusController) InsertUser() {
	email := strings.TrimSpace(this.GetString("email"))
	name := strings.TrimSpace(this.GetString("name"))
	password := strings.TrimSpace(this.GetString("password"))

	result, err := models.InsertUser(name, password, email)

	if err != nil || result <= 0 {
		this.TplName = "contactusTemplate.html"
		this.Data["error"] = "資料新增失敗"
		return
	}
	this.Redirect("/contactus", 302)
}

// 更新user資料
func (this *ContactusController) UpdateUser() {

	id := strings.TrimSpace(this.GetString("id")) // [id]：獲取input資料，[:id]：獲取路由參數資料
	// id := strings.TrimSpace(this.Input().Get("id")) 	//只獲取input資料
	name := strings.TrimSpace(this.GetString("name"))

	userInfo := make(map[string]string)
	userInfo["id"] = id
	userInfo["username"] = name

	num, err := models.UpdateUserArray(userInfo)

	if err == nil && num <= 0 {
		this.TplName = "contactusTemplate.html"
		this.Data["error"] = "資料更新失敗"
		return
	}
	this.Redirect("/contactus", 302)
}

// 刪除 user
func (this *ContactusController) DelUser() {
	id := this.GetString(":delId")
	result, err := models.DeleteUser(id)

	if err != nil || result <= 0 {
		this.TplName = "contactusTemplate.html"
		this.Data["error"] = "資料刪除失敗"
		return
	}

	this.Redirect("/contactus", 302)
}
