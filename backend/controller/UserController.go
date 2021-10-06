package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"priviledge/common"
	"priviledge/dto"
	"priviledge/model"
	"priviledge/response"
	"priviledge/util"
)

func Register(c *gin.Context) {
	DB := common.GetDB()
	//var requestMap = make(map[string]string)
	//json.NewDecoder(c.Request.Body).Decode(&requestMap)
	var requestUser = model.User{}
	c.Bind(&requestUser)

	// 获取参数
	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password
	// 数据验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号码不符合格式")
		return
	}

	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不符合格式")
		return
	}

	if len(name) == 0 {
		name = util.RandomString(10)
	}
	// 手机存在性判断

	if isTelephoneExist(DB, telephone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户已存在")
		return
	}

	// 持久化
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	DB.Create(&newUser)
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "数据库内部错误")
		log.Printf("Token generate error : %v", err)
		return
	}
	response.Success(c, gin.H{
		"token": token,
		"user": newUser,
	}, "登录成功")

}

func Login(c *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	var requestUser = model.User{}
	c.Bind(&requestUser)

	// 获取参数
	telephone := requestUser.Telephone
	password := requestUser.Password

	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号码不符合格式")
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不符合格式")
		return
	}
	// 判断手机号是否存在
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)

	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(c, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "数据库内部错误")
		log.Printf("Token generate error : %v", err)
		return
	}
	// 返回结果

	response.Success(c, gin.H{
		"token": token,
	}, "登录成功")

}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func Info(c *gin.Context) {
	user, _ := c.Get("user")
	response.Success(c, gin.H{
		"user": dto.ToUserDto(user.(model.User)),
	}, "成功")
}
