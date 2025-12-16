package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID        int
	Name      string
	Age       int
	Email     string
	Posts     []Post
	PostCount int
}

type Post struct {
	ID           int
	Title        string
	Content      string
	UserID       int
	Comments     []Comment
	CommentCount int
	Status       string `gorm:"default:'无评论'"`
}

// 增加文章自动给用户文章数+1
func (pn *Post) AfterCreate(tx *gorm.DB) (err error) {
	if pn.UserID > 0 {
		tx.Model(&User{}).Where("id = ?", pn.UserID).Update("post_count", gorm.Expr("post_count + ?", 1))
	}
	return
}

type Comment struct {
	ID      int
	Content string
	PostID  int
}

// 增加评论后自动给相应的评论+1,如果第一次评论给文字状态改为有评论
func (cn *Comment) AfterCreate(tx *gorm.DB) (err error) {
	if cn.PostID > 0 {
		tx.Model(&Post{}).Where("id = ?", cn.PostID).Update("comment_count", gorm.Expr("comment_count + ?", 1))
	}
	var post Post
	tx.Model(&Post{}).Where("id = ?", cn.PostID).First(&post)
	if post.Status == "无评论" {
		tx.Model(&Post{}).Where("id = ?", cn.PostID).Update("status", "有评论")
	}
	return
}

// 删除评论后给文章的评论数-1，如果评论数为0，则将文章的status改为无评论
func (cn *Comment) AfterDelete(tx *gorm.DB) (err error) {
	if cn.PostID > 0 {
		tx.Model(&Post{}).Where("id = ?", cn.PostID).Update("comment_count", gorm.Expr("comment_count - ?", 1))
	}
	var post Post
	tx.Model(&Post{}).Where("id = ?", cn.PostID).Find(&post)
	//如果评论数为0，则将文章的status改为无评论
	if post.CommentCount <= 0 {
		tx.Model(&Post{}).Where("id = ?", cn.PostID).Update("status", "无评论")
	}
	return
}

func RunGorm(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
	db.Create(&User{
		Name:  "xiaohan",
		Email: "xiaohan@gmail.com",
		Age:   18,
		Posts: []Post{
			{
				Title:   "post1",
				Content: "content1",
				Comments: []Comment{
					{
						Content: "post1 comment1",
					},
					{
						Content: "post1 comment2",
					},
				},
			},
			{
				Title:   "post2",
				Content: "content2",
				Comments: []Comment{
					{
						Content: "post2 comment1",
					},
					{
						Content: "post2 comment2",
					},
					{
						Content: "post2 comment3",
					},
				},
			},
		},
	})

	// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	var user User
	db.Debug().Preload("Posts.Comments").Find(&user, 1) //直接预加载文章和文章下的评论
	fmt.Println(user)

	//编写Go代码，使用Gorm查询评论数量最多的文章信息。
	var post Post
	db.Debug().Select("posts.*,count(comments.id) as comment_count").Joins("left join comments on posts.id = comments.post_id").
		Group("posts.id").Order("comment_count desc").Limit(1).Find(&post)
	fmt.Println(post) //输出评论数最高的文章

	//在评论删除时触发钩子函数，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
	//但是实测下面一次性删除的方式无法触发钩子函数
	//db.Debug().Delete(&Comment{}, []int{1, 2})

	//需要一个一个删除才可以出发钩子函数
	var comments []Comment
	db.Where("id IN ?", []int{1, 2}).Find(&comments)
	for _, comment := range comments {
		db.Delete(&comment)
	}

}

func main() {
	//连接数据库
	dsn := "root:xiaohan@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	RunGorm(db)
}
