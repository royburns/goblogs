package models

import (
	//_ "code.google.com/p/go-mysql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

//admin表
type Blogs struct {
	Id      int `PK`
	Title   string
	Content string
	Created string
}

//添加博客文章
func InsertBlogs(blog Blogs) Blogs {
	db := GetLink()
	db.Save(&blog)
	defer db.Db.Close()
	return blog
}

//查询博客
func GetBlogsList(start, offset int) (blog []Blogs) {
	db := GetLink()
	defer db.Db.Close()
	db.Limit(offset, start).FindAll(&blog)
	return
}

//查询所有博客
func GetAllBlogList() (blog []Blogs) {
	db := GetLink()
	defer db.Db.Close()
	db.FindAll(&blog)
	return
}

//获取博客总数量
func GetAllBlogsCount() (resultsSlice []map[string][]byte, err error) {
	db := GetLink()
	defer db.Db.Close()
	return db.SetTable("blogs").Select("count(*)").FindMap()
}

//更新博客信息
func UpdateBlogInfo(blog Blogs) (Blogs, error) {
	db := GetLink()
	err := db.Save(&blog)
	defer db.Db.Close()
	return blog, err
}

//根据id查询博客信息
func GetBlogInfoById(id int) (blog Blogs) {
	db := GetLink()
	db.Where("id=?", id).Find(&blog)
	defer db.Db.Close()
	return
}

//删除博客
func DelBlogById(blog Blogs) {
	db := GetLink()
	db.Delete(&blog)
	defer db.Db.Close()
	return
}
