package controller

import (
	"go-mall/model"
	"go-mall/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct{}

// LoginRequest 登录请求
type LoginRequest struct {
	Code      string `json:"code" binding:"required"`
	Nickname  string `json:"nickname"`
	Avatar    string `json:"avatar"`
	Gender    int    `json:"gender"`
	Country   string `json:"country"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Language  string `json:"language"`
}

// Login 微信小程序登录
func (uc *UserController) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "参数错误")
		return
	}

	// TODO: 调用微信API获取openid
	// 这里简化处理，实际需要调用微信接口
	openid := req.Code // 临时使用code作为openid

	// 查找或创建用户
	var user model.User
	err := utils.DB.Where("openid = ?", openid).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 创建新用户
			user = model.User{
				OpenID:   openid,
				Nickname: req.Nickname,
				Avatar:   req.Avatar,
				Gender:   req.Gender,
				Country:  req.Country,
				Province: req.Province,
				City:     req.City,
				Language: req.Language,
			}
			if err := utils.DB.Create(&user).Error; err != nil {
				utils.ServerError(c, "创建用户失败")
				return
			}
		} else {
			utils.ServerError(c, "查询用户失败")
			return
		}
	} else {
		// 更新用户信息
		user.Nickname = req.Nickname
		user.Avatar = req.Avatar
		user.Gender = req.Gender
		user.Country = req.Country
		user.Province = req.Province
		user.City = req.City
		user.Language = req.Language
		utils.DB.Save(&user)
	}

	// 生成JWT token
	token, err := utils.GenerateToken(user.ID, user.OpenID)
	if err != nil {
		utils.ServerError(c, "生成token失败")
		return
	}

	utils.Success(c, gin.H{
		"token": token,
		"user":  user,
	})
}

// GetUserInfo 获取用户信息
func (uc *UserController) GetUserInfo(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.Unauthorized(c, "用户未登录")
		return
	}

	var user model.User
	if err := utils.DB.First(&user, userID).Error; err != nil {
		utils.NotFound(c, "用户不存在")
		return
	}

	utils.Success(c, user)
}

// UpdateUserInfo 更新用户信息
func (uc *UserController) UpdateUserInfo(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.Unauthorized(c, "用户未登录")
		return
	}

	var req struct {
		Nickname string `json:"nickname"`
		Avatar   string `json:"avatar"`
		Phone    string `json:"phone"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "参数错误")
		return
	}

	var user model.User
	if err := utils.DB.First(&user, userID).Error; err != nil {
		utils.NotFound(c, "用户不存在")
		return
	}

	// 更新用户信息
	if req.Nickname != "" {
		user.Nickname = req.Nickname
	}
	if req.Avatar != "" {
		user.Avatar = req.Avatar
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}

	if err := utils.DB.Save(&user).Error; err != nil {
		utils.ServerError(c, "更新用户信息失败")
		return
	}

	utils.Success(c, user)
} 