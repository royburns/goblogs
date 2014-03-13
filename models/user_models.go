package models

type User struct {
	Id             int `PK`
	Username       string
	Password       string
	Created        string
	Last_logintime string
}

//用户注册
func AddUser(users User) User {
	db := GetLink()
	db.Save(&users)
	defer db.Db.Close()
	return users
}

//查询用户
func GetUserInfo(username string) (users User) {
	db := GetLink()
	db.Where("username=?", username).Find(&users)
	defer db.Db.Close()
	return
}

//更新数据
func UpdateUserInfo(users User) (User, error) {
	db := GetLink()
	err := db.Save(&users)
	defer db.Db.Close()
	return users, err
}
