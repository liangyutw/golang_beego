package models

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int
	Username string
	Email    string
	Password string
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func init() {
	// 需要在init中註冊model
	orm.RegisterModel(new(User))
}

//新增USER操作
func InsertUser(name string, password string, email string) (int64, error) {

	var userInfo User
	hash, _ := HashPassword(password)
	userInfo.Username = name
	userInfo.Password = hash
	userInfo.Email = email
	id, err := orm.NewOrm().Insert(&userInfo)
	return id, err
}

//取User資料
//返回结果集的 key => value 值
func GetUserData() []orm.Params {
	var maps []orm.Params
	_, err := orm.NewOrm().QueryTable("user").Values(&maps)

	if err != nil {
		fmt.Printf("Something went wrong %d", err)
	}

	return maps
}

//更新USER操作
func UpdateUserdDefault(id string, name string) (int64, error) {

	params := make(map[string]interface{})
	params["id"] = id
	params["username"] = name

	num, err := orm.NewOrm().QueryTable("user").Filter("id", id).Update(params)

	return num, err
}

// 多個傳遞參數
func UpdateUserdMultiParameters(args ...string) (int64, error) {

	params := make(map[string]interface{})

	params["id"] = args[0]
	params["username"] = args[1]

	num, err := orm.NewOrm().QueryTable("user").Filter("id", params["id"]).Update(params)

	return num, err
}

// 傳遞陣列
func UpdateUserArray(array map[string]string) (int64, error) { //陣列型態傳遞
	params := make(map[string]interface{})

	for k, v := range array {
		params[k] = v
	}
	fmt.Println(params)

	num, err := orm.NewOrm().QueryTable("user").Filter("id", params["id"]).Update(params)

	return num, err
}

//刪除USER操作
func DeleteUser(id string) (int64, error) {

	num, err := orm.NewOrm().QueryTable("user").Filter("id", id).Delete()

	return num, err
}
