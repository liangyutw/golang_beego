使用 golang + beego framework 實作 crud 的內容

# 開發環境
windows 10 WSL2 <BR>
Docker Desktop: 4.9.1 <BR>
# 使用版本
golang: 1.18.3 [go](https://go.dev/doc/) <BR>
beego framework: 2.0.4 [beego](https://beego.vip/docs/intro/) <BR>
# 教學文章
## beego framework
https://www.golangprograms.com/golang/beego-setup-installation/ <BR>
當中的 `go install github.com/astaxie/beego` 改成 `go install github.com/beego/beego/v2/server/web` <BR>

## 資料庫連線
```
mysql 要先在 docker 運作成功，並且能成功連線建立資料庫
```
https://developpaper.com/go-practical-project-basic-use-of-beegos-orm/ <br>

## 密碼加密
https://gowebexamples.com/password-hashing/ <br>

# 實際操作檔案
* controllers/contactus.go
* models/user.go
* routers/router.go
* views/contactusTemplate.html
* database/migrations/20220713_164011_user_table.go
