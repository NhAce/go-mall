package controller

import (
	"go-mall/model"
	"go-mall/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct{}

// GetProductList 获取商品列表
func (pc *ProductController) GetProductList(c *gin.Context) {
	// 获取查询参数
	categoryID := c.Query("category_id")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	
	if page < 1 {
		page = 1
	}
	if size < 1 || size > 100 {
		size = 10
	}

	offset := (page - 1) * size

	var products []model.Product
	var total int64

	// 构建查询
	query := utils.DB.Model(&model.Product{}).Where("status = ?", 1)
	
	if categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}

	// 获取总数
	query.Count(&total)

	// 获取商品列表
	err := query.Preload("Category").
		Order("sort DESC, created_at DESC").
		Offset(offset).
		Limit(size).
		Find(&products).Error

	if err != nil {
		utils.ServerError(c, "获取商品列表失败")
		return
	}

	utils.Success(c, gin.H{
		"list":  products,
		"total": total,
		"page":  page,
		"size":  size,
	})
}

// GetProductDetail 获取商品详情
func (pc *ProductController) GetProductDetail(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.BadRequest(c, "商品ID不能为空")
		return
	}

	var product model.Product
	err := utils.DB.Preload("Category").First(&product, id).Error
	if err != nil {
		utils.NotFound(c, "商品不存在")
		return
	}

	utils.Success(c, product)
}

// GetCategoryList 获取分类列表
func (pc *ProductController) GetCategoryList(c *gin.Context) {
	var categories []model.Category
	err := utils.DB.Where("status = ?", 1).
		Order("sort ASC").
		Find(&categories).Error

	if err != nil {
		utils.ServerError(c, "获取分类列表失败")
		return
	}

	utils.Success(c, categories)
} 