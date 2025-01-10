package controller

import (
	"net/http"
	"strconv"
	"your_project/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GroupController struct {
	GroupService *service.GroupService
}

// InsertGroup 创建群组
// @Summary 创建群组
// @Description 用户创建群组
// @Tags 用户群组管理
// @Accept json
// @Produce json
// @Param userUid formData int true "用户 UID"
// @Param friendUid formData int true "好友 UID"
// @Param groupName formData string true "群组名称"
// @Success 200 {object} entity.ResultEntity
// @Router /group/insertGroup [post]
func (gc *GroupController) InsertGroup(c *gin.Context) {
	userUid, _ := strconv.Atoi(c.PostForm("userUid"))
	friendUid, _ := strconv.Atoi(c.PostForm("friendUid"))
	groupName := c.PostForm("groupName")

	db := c.MustGet("db").(*gorm.DB)
	result, err := gc.GroupService.InsertGroup(db, userUid, friendUid, groupName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// SendEmailToAll 群发消息
// @Summary 群发消息
// @Description 群组内群发消息
// @Tags 用户群组管理
// @Accept json
// @Produce json
// @Param userUid formData int true "用户 UID"
// @Param groupName formData string true "群组名称"
// @Param title formData string true "邮件标题"
// @Param content formData string true "邮件内容"
// @Success 200 {object} entity.ResultEntity
// @Router /group/sendEmailToAll [post]
func (gc *GroupController) SendEmailToAll(c *gin.Context) {
	userUid, _ := strconv.Atoi(c.PostForm("userUid"))
	groupName := c.PostForm("groupName")
	title := c.PostForm("title")
	content := c.PostForm("content")

	db := c.MustGet("db").(*gorm.DB)
	result, err := gc.GroupService.SendEmailToAll(db, userUid, groupName, title, content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
