package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"your_project/service"
)

// EmailController 邮件控制器
type EmailController struct {
	EmailService *service.EmailService
}

// SendEmail 发送邮件接口
// @Summary 发送邮件
// @Description 发送一封邮件
// @Tags Email
// @Accept json
// @Produce json
// @Param senderUid body int true "发信者uid"
// @Param receiverUid body int true "收信者uid"
// @Param title body string true "信件标题"
// @Param content body string true "信件内容"
// @Success 200 {object} entity.ResultEntity
// @Failure 400 {object} entity.ResultEntity
// @Router /email/sendEmail [post]
func (ctrl *EmailController) SendEmail(c *gin.Context) {
	senderUid, _ := strconv.Atoi(c.DefaultPostForm("senderUid", "0"))
	receiverUid, _ := strconv.Atoi(c.DefaultPostForm("receiverUid", "0"))
	title := c.DefaultPostForm("title", "")
	content := c.DefaultPostForm("content", "")

	message, status := ctrl.EmailService.SendEmail(senderUid, receiverUid, title, content)
	c.JSON(http.StatusOK, gin.H{"message": message, "status": status})
}

// ReceiveEmail 接收邮件接口并分页展示
// @Summary 接收邮件并分页展示
// @Description 根据发信者UID分页展示邮件
// @Tags Email
// @Accept json
// @Produce json
// @Param receiverUid query int true "收信者uid"
// @Param pageNum query int true "页码"
// @Success 200 {object} entity.ResultEntity
// @Failure 400 {object} entity.ResultEntity
// @Router /email/receiveEmail [get]
func (ctrl *EmailController) ReceiveEmail(c *gin.Context) {
	receiverUid, _ := strconv.Atoi(c.DefaultQuery("receiverUid", "0"))
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))

	emails, message, status := ctrl.EmailService.ReceiveEmail(receiverUid, pageNum)
	c.JSON(http.StatusOK, gin.H{"message": message, "status": status, "emails": emails})
}

// SelectEmail 根据标题和内容模糊查询
// @Summary 根据标题和内容模糊查询邮件
// @Description 根据关键词查询邮件
// @Tags Email
// @Accept json
// @Produce json
// @Param receiverUid query int true "收信者uid"
// @Param keyword query string true "查询内容"
// @Success 200 {object} entity.ResultEntity
// @Failure 400 {object} entity.ResultEntity
// @Router /email/selectEmail [get]
func (ctrl *EmailController) SelectEmail(c *gin.Context) {
	receiverUid, _ := strconv.Atoi(c.DefaultQuery("receiverUid", "0"))
	keyword := c.DefaultQuery("keyword", "")

	emails, message, status := ctrl.EmailService.SelectEmail(receiverUid, keyword)
	c.JSON(http.StatusOK, gin.H{"message": message, "status": status, "emails": emails})
}

// DriftEmail 漂流瓶功能接口
// @Summary 漂流瓶功能
// @Description 发送漂流瓶邮件
// @Tags Email
// @Accept json
// @Produce json
// @Param senderUid body int true "发信者uid"
// @Param title body string true "信件标题"
// @Param content body string true "信件内容"
// @Success 200 {object} entity.ResultEntity
// @Failure 400 {object} entity.ResultEntity
// @Router /email/driftEmail [post]
func (ctrl *EmailController) DriftEmail(c *gin.Context) {
	senderUid, _ := strconv.Atoi(c.DefaultPostForm("senderUid", "0"))
	title := c.DefaultPostForm("title", "")
	content := c.DefaultPostForm("content", "")

	message, status := ctrl.EmailService.DriftEmail(senderUid, title, content)
	c.JSON(http.StatusOK, gin.H{"message": message, "status": status})
}

// TimeEmail 定时发送邮件接口
// @Summary 定时发送邮件
// @Description 发送定时邮件
// @Tags Email
// @Accept json
// @Produce json
// @Param senderUid body int true "发信者uid"
// @Param receiverUid body int true "收信者uid"
// @Param title body string true "信件标题"
// @Param content body string true "信件内容"
// @Param time body int true "多少分钟后发送"
// @Success 200 {object} entity.ResultEntity
// @Failure 400 {object} entity.ResultEntity
// @Router /email/timeEmail [post]
func (ctrl *EmailController) TimeEmail(c *gin.Context) {
	senderUid, _ := strconv.Atoi(c.DefaultPostForm("senderUid", "0"))
	receiverUid, _ := strconv.Atoi(c.DefaultPostForm("receiverUid", "0"))
	title := c.DefaultPostForm("title", "")
	content := c.DefaultPostForm("content", "")
	time, _ := strconv.Atoi(c.DefaultPostForm("time", "0"))

	message, status := ctrl.EmailService.TimeEmail(senderUid, receiverUid, title, content, time)
	c.JSON(http.StatusOK, gin.H{"message": message, "status": status})
}

// DeleteEmail 删除邮件接口
// @Summary 删除邮件
// @Description 根据发信者UID和收信者UID删除邮件
// @Tags Email
// @Accept json
// @Produce json
// @Param senderUid body int true "发信者uid"
// @Param receiverUid body int true "收信者uid"
// @Param title body string true "信件标题"
// @Success 200 {object} entity.ResultEntity
// @Failure 400 {object} entity.ResultEntity
// @Router /email/deleteEmail [post]
func (ctrl *EmailController) DeleteEmail(c *gin.Context) {
	senderUid, _ := strconv.Atoi(c.DefaultPostForm("senderUid", "0"))
	receiverUid, _ := strconv.Atoi(c.DefaultPostForm("receiverUid", "0"))
	title := c.DefaultPostForm("title", "")

	message, status := ctrl.EmailService.DeleteEmail(senderUid, receiverUid, title)
	c.JSON(http.StatusOK, gin.H{"message": message, "status": status})
}
