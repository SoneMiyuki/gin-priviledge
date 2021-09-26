package controller

import (
	"github.com/jinzhu/gorm"
	"net/http"
	"github.com/gin-gonic/gin"
	"priviledge/common"
	"priviledge/model"
	"priviledge/util"
)

func Register(c *gin.Context) {
	DB := common.GetDB()
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
		name = util.RandomString(10);
	}
	// 手机存在性判断

	if isTelephoneExist(DB, telephone){
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code" : 422,
			"msg": "用户存在",
		})
		return
	}

	// 持久化
	newUser :=model.User{
		Name:	name,
		Telephone: telephone,
		Password: password,
	}
	DB.Create(&newUser)

	// 返回结果
	c.JSON(http.StatusOK, gin.H{
		"message" : "注册成功",
	})

}



func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?",telephone).First(&user)
	if user.ID != 0{
		return true
	}
	return false
}