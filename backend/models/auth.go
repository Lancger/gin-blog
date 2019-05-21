package models

import (
	"gin-blog/pkg/e"
	"gin-blog/pkg/util"

	"net/http"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	if auth.ID > 0 {
		return true
	}

	return false
}

// AccountInfo 用户信息
func AccountInfo(ctx *gin.Context) {
	type AccountMsg struct {
		Name   string   `json:"name"`
		Roles  []string `json:"roles"`
		Perms  []string `json:"perms"`
		Avatar string   `json:"avatar"`
	}

	accountMsg := AccountMsg{
		Name:   "nil",
		Roles:  []string{},
		Perms:  []string{},
		Avatar: "",
	}
	var auth Auth

	token := ctx.GetHeader("X-Token")
	if token == `` {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": e.ERROR_AUTH,
			"data":    accountMsg,
		})
		return
	}
	claims, err := util.ParseToken(token)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    e.ERROR,
			"message": err,
			"data":    accountMsg,
		})
		return
	}
	accountMsg.Name = claims.Username
	// 获取用户组信息
	if err = db.Set("gorm:auto_preload", true).Model(&Auth{}).Where("username = ?", accountMsg.Name).First(&auth).Error; err == nil {
		// for _, i := range u.SystemGroups {
		// 	accountMsg.Roles = append(accountMsg.Roles, i.GroupName)
		// }
		// accountMsg.Perms, _ = u.getPermission()
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    e.SUCCESS,
		"message": "success",
		"data":    accountMsg,
	})
	return
}
