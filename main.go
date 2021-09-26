package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	Name string `grom:"type:varchar(20);not null"`
	Telephone string `gorm:"varchar(110);not null;unique"`
	Password string `gorm:"size: 255;not null"`
}

func main(){
	db := InitDB()
	defer
		db.Close()

	r := gin.Default()
	r.POST("/api/auth/register", func(c *gin.Context) {

		// 获取参数
		name := c.PostForm("name")
		telephone := c.PostForm("telephone")
		password := c.PostForm("password")
		// 数据验证
		if len(telephone) != 11 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code" : 422,
				"msg": "手机号码不符合格式",
			})
			return
		}


		if len(password) < 6{
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code" : 422,
				"msg": "密码不符合格式",
			})
			return
		}

		if len(name) == 0{
			name = RandomString(10);
		}
		// 手机存在性判断

		if isTelephoneExist(db, telephone){
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code" : 422,
				"msg": "用户存在",
			})
			return
		}

		// 持久化
		newUser :=User{
			Name:	name,
			Telephone: telephone,
			Password: password,
		}
		db.Create(&newUser)

		// 返回结果
		c.JSON(http.StatusOK, gin.H{
			"message" : "注册成功",
		})

	})
	panic(r.Run())
}

func RandomString(n int) string{
	var letters = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result{
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func InitDB() *gorm.DB{
	driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "ginvue"
	username := "root"
	password := "123456"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)

	db, err := gorm.Open(driverName, args)
	if err != nil{
		panic("fail to connect database, err: " + err.Error())
	}
	db.AutoMigrate(&User{})

	return db
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone = ?",telephone).First(&user)
	if user.ID != 0{
		return true
	}
	return false
}