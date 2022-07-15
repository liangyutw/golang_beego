package main

import (
	_ "demoProject/routers"
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

// 註冊sql驅動
func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	maxIdle := 10
	maxConn := 100
	err := orm.RegisterDataBase("default", "mysql", "root:password@/goblog?charset=utf8", orm.MaxIdleConnections(maxIdle), orm.MaxOpenConnections(maxConn))

	// 正常連線下出錯，才會顯示在終端機
	if err != nil {
		fmt.Println("RegisterDataBase err: ", err)
	}
}

func main() {
	web.Run()
}
