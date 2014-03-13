package models

import (
	//_ "code.google.com/p/go-mysql-driver/mysql"
	"database/sql"
	"github.com/astaxie/beedb"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

const (
	mysql_user    = "root"
	mysql_pass    = ""
	mysql_host    = "localhost"
	mysql_port    = "3306"
	database_name = "myblog"
)

//admin表
type Admin struct {
	Id       int `PK`
	Username string
	Password string
}

//数据库连接
func GetLink() beedb.Model {
	//beedb.OnDebug = true
	db, err := sql.Open("mysql", mysql_user+":"+mysql_pass+"@tcp("+mysql_host+":"+mysql_port+")/"+database_name)
	if err != nil {
		panic(err)
	}

	orm := beedb.New(db)
	return orm
}

//根据用户名查询admin信息
func GetAdminInfo(username string) (admin Admin) {
	db := GetLink()
	db.Where("username=?", strings.TrimSpace(username)).Find(&admin)
	defer db.Db.Close()
	return
}
