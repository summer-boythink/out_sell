package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var id int64  //最终返回的ID
var json1 UserLogin //刻印发送过来的json
var u1 UserInfo//查找用户ID

//放json
type UserLogin struct {
	NickName string `form:"nickName" json:"nickName"` //用户名
	AvatarUrl string `form:"avatarUrl" json:"avatarUrl"` //头像
}

//放ID
type UserInfo struct {
	ID int `form:"id" json:"id"`
}

//插入
func insertInfo(nickname string,avatarurl string,db *sql.DB) (id int64) {
	sqlStr := `insert into user(nickname,avatarurl) values(?,?)`
	//exec
	ret,err := db.Exec(sqlStr,nickname,avatarurl)
	if err != nil {
		fmt.Println(err)
		return
	}
	//获取新数据id
	id,err = ret.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(id)
	return id
}

//查找
func queryInfo(nickname string,avatarurl string,db *sql.DB,u1 UserInfo) (Isexist bool) {
	sqlStr := `select id from user where nickname=? and avatarurl=?;`
	rowObj := db.QueryRow(sqlStr,nickname,avatarurl)
	//把结果打印到u1里面 必须scan 因为Scan可以释放连接
	err := rowObj.Scan(&u1.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	//设置数据库连接池的最大的连接数
	db.SetMaxOpenConns(10)
	if u1.ID > 0 {
		id = int64(u1.ID)
		return true
	}else {
		return false
	}
}

//main
func CheckLogin(client *sql.DB) gin.HandlerFunc{
	return func(c *gin.Context) {
		if err := c.ShouldBindJSON(&json1); err != nil {
			// 返回错误信息
			// gin.H封装了生成json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//fmt.Println(json)
		if !queryInfo(json1.NickName,json1.AvatarUrl,client,u1) {
			id = insertInfo(json1.NickName,json1.AvatarUrl,client)
		}

		fmt.Println(id)
		c.JSON(http.StatusOK,gin.H{"status":200,"userId":id})
	}
}