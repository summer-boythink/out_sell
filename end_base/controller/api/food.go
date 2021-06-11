package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Food struct {
	Id int `json:"id"`
	Name string `json:"name"`
	ImgUrl string `json:"imgurl"`
	Remark string `json:"remark"`
}


func CheckTake(client *sql.DB) gin.HandlerFunc{
	return func(c *gin.Context) {
		sqlStr := `select id,name,imgurl,remark from food where status=? and userid=?;`
		rows ,err := client.Query(sqlStr,c.Query("status"),c.Query("userid"))
		if err != nil {
			fmt.Println(err)
			return
		}
		//关闭 释放内存
		defer rows.Close()
		//fmt.Println(rows)
		//循环读取数据
		var foods = make([]Food,0)

		for rows.Next(){
			var food Food
			err := rows.Scan(&food.Id,&food.Name,&food.ImgUrl,&food.Remark)
			if err != nil {
				fmt.Println(err)
				return
			}
			foods = append(foods,food)

		}
		if len(foods) > 0 {
			c.JSON(http.StatusOK,gin.H{"status":200,"food":foods})
		}else {
			c.JSON(http.StatusOK,gin.H{"status":204,"message":"no food","food":foods})
		}

	}
}

//传id
func ChangeTake(client *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		sqlStr := `update food set status=1 where id=?`
		rows,err := client.Exec(sqlStr,c.Query("id"))
		if err != nil {
			fmt.Println(err)
			return
		}

		id,err := rows.RowsAffected()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(id)
		if id > 0 {
			c.JSON(http.StatusOK,gin.H{"status":200,"message":"ok change","id":c.Query("id")})
		}else {
			c.JSON(http.StatusOK,gin.H{"status":204,"message":"not change"})
		}
	}
}

//userid tele
//TODO 交给node了
func PutFood(client *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//sqlStr := `insert into food (name,remark,userid,code,status) values (?)`
		//		//rows,err := client.Exec(sqlStr,c.Query("id"))
		//		//if err != nil {
		//		//	fmt.Println(err)
		//		//	return
		//		//}
		//		//
		//		//id,err := rows.RowsAffected()
		//		//if err != nil {
		//		//	fmt.Println(err)
		//		//	return
		//		//}
		//		//fmt.Println(id)
		//		//if id > 0 {
		//		//	c.JSON(http.StatusOK,gin.H{"status":200,"message":"ok change","id":c.Query("id")})
		//		//}else {
		//		//	c.JSON(http.StatusOK,gin.H{"status":204,"message":"not change"})
		//		//}
	}
}