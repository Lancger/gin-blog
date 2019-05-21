package api

import (
	"gin-blog/backend/pkg/logging"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"gin-blog/backend/models"
	"gin-blog/backend/pkg/e"
	"gin-blog/backend/pkg/util"
)

// GetAuth 登录
func GetAuth(c *gin.Context) {

	var (
		req models.Auth
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
