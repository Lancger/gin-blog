package api

import (
	"gin-blog/pkg/logging"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"gin-blog/models"
	"gin-blog/pkg/e"
	"gin-blog/pkg/util"
)

// type Auth struct {
// 	Username string `valid:"Required; MaxSize(50)"`
// 	Password string `valid:"Required; MaxSize(50)"`
// }

// GetAuth 登录
func GetAuth(c *gin.Context) {
	type reqPost struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var (
		req reqPost
	)
	// username := c.Query("username")
	// password := c.Query("password")

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": e.INVALID_PARAMS,
		})
		return
	}

	valid := validation.Validation{}
	a := models.Auth{Username: req.Username, Password: req.Password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		isExist := models.CheckAuth(req.Username, req.Password)
		if isExist {
			token, err := util.GenerateToken(req.Username, req.Password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token

				code = e.SUCCESS
			}

		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
