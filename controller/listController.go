package controller

import (
	"net/http"
	"strconv"
	"your_project/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ListController struct {
	ListService *service.ListService
}

// InsertFriend 添加好友
// @Summary 添加好友
// @Description 用户添加好友
// @Tags 好友管理
// @Accept json
// @Produce json
// @Param friendUid formData int true "好友 UID"
// @Param userUid formData int true "用户 UID"
// @Success 200 {object} entity.ResultEntity
// @Router /list/add [post]
func (lc *ListController) InsertFriend(c *gin.Context) {
	friendUid, _ := strconv.Atoi(c.PostForm("friendUid"))
	userUid, _ := strconv.Atoi(c.PostForm("userUid"))

	db := c.MustGet("db").(*gorm.DB)
	result, err := lc.ListService.InsertFriend(db, friendUid, userUid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// UpdateRemark 更新好友备注
// @Summary 更新好友备注
// @Description 用户更新好友备注
// @Tags 好友管理
// @Accept json
// @Produce json
// @Param userUid formData int true "用户 UID"
// @Param friendUid formData int true "好友 UID"
// @Param remark formData string true "备注名称"
// @Success 200 {object} entity.ResultEntity
// @Router /list/updateRemark [post]
func (lc *ListController) UpdateRemark(c *gin.Context) {
	userUid, _ := strconv.Atoi(c.PostForm("userUid"))
	friendUid, _ := strconv.Atoi(c.PostForm("friendUid"))
	remark := c.PostForm("remark")

	db := c.MustGet("db").(*gorm.DB)
	result, err := lc.ListService.UpdateRemark(db, userUid, friendUid, remark)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// SetRelationship 设置亲密关系
// @Summary 设置亲密关系
// @Description 用户设置好友亲密关系
// @Tags 好友管理
// @Accept json
// @Produce json
// @Param userUid formData int true "用户 UID"
// @Param friendUid formData int true "好友 UID"
// @Param relationship formData string true "亲密关系名称"
// @Success 200 {object} entity.ResultEntity
// @Router /list/setRelationship [post]
func (lc *ListController) SetRelationship(c *gin.Context) {
	userUid, _ := strconv.Atoi(c.PostForm("userUid"))
	friendUid, _ := strconv.Atoi(c.PostForm("friendUid"))
	relationship := c.PostForm("relationship")

	db := c.MustGet("db").(*gorm.DB)
	result, err := lc.ListService.SetRelationship(db, userUid, friendUid, relationship)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// DeleteFriend 假删除好友
// @Summary 假删除好友
// @Description 用户假删除好友
// @Tags 好友管理
// @Accept json
// @Produce json
// @Param userUid formData int true "用户 UID"
// @Param friendUid formData int true "好友 UID"
// @Success 200 {object} entity.ResultEntity
// @Router /list/deleteFriend [post]
func (lc *ListController) DeleteFriend(c *gin.Context) {
	userUid, _ := strconv.Atoi(c.PostForm("userUid"))
	friendUid, _ := strconv.Atoi(c.PostForm("friendUid"))

	db := c.MustGet("db").(*gorm.DB)
	result, err := lc.ListService.DeleteFriend(db, userUid, friendUid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// CheckWhoAdd 查看好友申请
// @Summary 查看好友申请
// @Description 查看哪些用户申请添加当前用户为好友
// @Tags 好友管理
// @Accept json
// @Produce json
// @Param userUid query int true "用户 UID"
// @Success 200 {object} entity.ResultEntity
// @Router /list/checkWhoAdd [get]
func (lc *ListController) CheckWhoAdd(c *gin.Context) {
	userUid, _ := strconv.Atoi(c.Query("userUid"))

	db := c.MustGet("db").(*gorm.DB)
	result, err := lc.ListService.CheckWhoAdd(db, userUid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// AgreeFriendAdd 同意好友申请
// @Summary 同意好友申请
// @Description 用户同意好友申请
// @Tags 好友管理
// @Accept json
// @Produce json
// @Param userUid formData int true "用户 UID"
// @Param friendUid formData int true "好友 UID"
// @Success 200 {object} entity.ResultEntity
// @Router /list/agreeFriendAdd [post]
func (lc *ListController) AgreeFriendAdd(c *gin.Context) {
	userUid, _ := strconv.Atoi(c.PostForm("userUid"))
	friendUid, _ := strconv.Atoi(c.PostForm("friendUid"))

	db := c.MustGet("db").(*gorm.DB)
	result, err := lc.ListService.AgreeFriendAdd(db, userUid, friendUid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// ShowAllFriend 查看好友列表
// @Summary 查看好友列表
// @Description 用户分页查看好友列表
// @Tags 好友管理
// @Accept json
// @Produce json
// @Param userUid query int true "用户 UID"
// @Param pageNum query int true "页码（每页 10 个好友）"
// @Success 200 {object} entity.ResultEntity
// @Router /list/showAllFriend [get]
func (lc *ListController) ShowAllFriend(c *gin.Context) {
	userUid, _ := strconv.Atoi(c.Query("userUid"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))

	db := c.MustGet("db").(*gorm.DB)
	result, err := lc.ListService.ShowAllFriend(db, userUid, pageNum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
