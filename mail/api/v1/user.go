package v1

import (
	"GOproject/GIT/mail/pkg/util"
	"GOproject/GIT/mail/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegister(c *gin.Context) {
	//相当于创建了一个UserRegisterService对象，调用这个对象中的Register方法。
	var userRegisterService service.UserService
	if err := c.ShouldBind(&userRegisterService); err == nil {
		res := userRegisterService.Register(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusNotFound, err)
	}
}
func UserLogin(c *gin.Context) {
	var userLogin service.UserService
	if err := c.ShouldBind(&userLogin); err == nil {
		res := userLogin.Login(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusNotFound, err)
	}
}
func UserUpdate(c *gin.Context) {
	var userUpdate service.UserService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userUpdate); err == nil {
		res := userUpdate.Update(c.Request.Context(), claims.Id)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusNotFound, err)
	}
}
func UploadAvatar(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	filesize := fileHeader.Size
	var uploadAvatar service.UserService
	claims, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&uploadAvatar); err == nil {
		res := uploadAvatar.Post(c.Request.Context(), claims.Id, file, filesize)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusNotFound, err)
	}

}
