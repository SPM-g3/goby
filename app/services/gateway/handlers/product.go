package handlers

import (
	"context"
	"strconv"
	"time"

	rpc_product "github.com/bitdance-panic/gobuy/app/rpc/kitex_gen/product"
	clients "github.com/bitdance-panic/gobuy/app/services/gateway/biz/clients"
	"github.com/bitdance-panic/gobuy/app/services/gateway/biz/dal/tidb"
	"github.com/bitdance-panic/gobuy/app/services/gateway/biz/dao"
	"github.com/bitdance-panic/gobuy/app/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/client/callopt"
)

// handleProductPut 这是更新商品
// @Summary 这是一段Summary
// @Description 这是一段Description
// @Accept application/json
// @Produce application/json
// @Router /product [put]
func HandleUpdateProduct(ctx context.Context, c *app.RequestContext) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Fail(c, err.Error())
		return
	}
	var req rpc_product.UpdateProductReq
	if err := c.BindAndValidate(&req); err != nil {
		utils.Fail(c, err.Error())
		return
	}
	req.Id = int32(id)
	resp, err := clients.ProductClient.UpdateProduct(context.Background(), &req, callopt.WithRPCTimeout(5*time.Second))
	if err != nil {
		utils.Fail(c, err.Error())
		return
	}
	utils.Success(c, utils.H{"newProduct": resp.Product})
}

// handleProductPost 这是创建商品
// @Summary 这是一段Summary
// @Description 这是一段Description
// @Accept application/json
// @Produce application/json
// @Router /product [post]
func HandleCreateProduct(ctx context.Context, c *app.RequestContext) {
	var req rpc_product.CreateProductReq
	if err := c.Bind(&req); err != nil {
		utils.Fail(c, err.Error())
		return
	}
	resp, err := clients.ProductClient.CreateProduct(context.Background(), &req, callopt.WithRPCTimeout(10*time.Second))
	if err != nil {
		utils.Fail(c, err.Error())
		return
	}
	utils.Success(c, utils.H{"product": resp.Product})
}

// handleProductDELETE 这是删除商品
// @Summary 这是一段Summary
// @Description 这是一段Description
// @Accept application/json
// @Produce application/json
// @Router /product [delete]
func HandleRemoveProduct(ctx context.Context, c *app.RequestContext) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Fail(c, err.Error())
		return
	}
	req := rpc_product.RemoveProductReq{
		Id: int32(id),
	}
	resp, err := clients.ProductClient.RemoveProduct(context.Background(), &req, callopt.WithRPCTimeout(10*time.Second))
	if err != nil {
		utils.Fail(c, err.Error())
		return
	}
	if resp.Success {
		utils.Success(c, nil)
	} else {
		utils.Fail(c, "删除失败")
	}
}

// handleProductGet 这是获取一个商品
// @Summary 这是一段Summary
// @Description 这是一段Description
// @Accept application/json
// @Produce application/json
// @Router /product [get]
func HandleGetProduct(ctx context.Context, c *app.RequestContext) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Fail(c, err.Error())
		return
	}
	req := rpc_product.GetProductByIDReq{
		Id: int32(id),
	}
	resp, err := clients.ProductClient.GetProductByID(context.Background(), &req, callopt.WithRPCTimeout(5*time.Second))
	if err != nil {
		utils.Fail(c, err.Error())
		return
	}
	utils.Success(c, utils.H{"product": resp.Product})
}

// handleProductSearch 这是模糊查询商品
// @Summary 这是一段Summary
// @Description 这是一段Description
// @Accept application/json
// @Produce application/json
// @Router /product/search [get]
func HandleSearchProducts(ctx context.Context, c *app.RequestContext) {
	pageNumStr := c.Query("page")
	pageSizeStr := c.Query("size")
	pageNum, err := strconv.Atoi(pageNumStr)
	if err != nil {
		utils.Fail(c, err.Error())
		return
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		utils.Fail(c, err.Error())
		return
	}
	// var body struct {
	// 	Query string `json:"query"`
	// }
	// err = c.Bind(&body)
	// if err != nil {
	// 	utils.Fail(c, err.Error())
	// 	return
	// }
	query := c.Query("query")
	category := c.Query("category")
	req := rpc_product.SearchProductsReq{
		Query:    query,
		Category: category,
		PageNum:  int32(pageNum),
		PageSize: int32(pageSize),
	}
	resp, err := clients.ProductClient.SearchProducts(context.Background(), &req, callopt.WithRPCTimeout(10*time.Second))
	if err != nil {
		utils.Fail(c, err.Error())
		return
	}
	utils.Success(c, utils.H{"products": resp.Products, "total_count": resp.TotalCount})
}

func HandleListAllProduct(ctx context.Context, c *app.RequestContext) {
	resp, err := clients.ProductClient.ListProduct(context.Background(), callopt.WithRPCTimeout(10*time.Second))
	if err != nil {
		utils.Fail(c, err.Error())
		return
	}
	utils.Success(c, utils.H{"products": resp.Products, "total_count": resp.TotalCount})
}

// handleProductSearch 这是模糊查询商品
// @Summary 这是一段Summary
// @Description 这是一段Description
// @Accept application/json
// @Produce application/json
// @Router /product/search [get]
func HandleAdminListProduct(ctx context.Context, c *app.RequestContext) {
	pageNum, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		utils.Fail(c, err.Error())
		return
	}
	pageSize, err := strconv.Atoi(c.Query("size"))
	if err != nil {
		utils.Fail(c, err.Error())
		return
	}
	req := rpc_product.ListProductReq{
		PageNum:  int32(pageNum),
		PageSize: int32(pageSize),
	}
	resp, err := clients.ProductClient.AdminListProduct(context.Background(), &req, callopt.WithRPCTimeout(10*time.Second))
	if err != nil {
		utils.Fail(c, err.Error())
		return
	}
	utils.Success(c, utils.H{"products": resp.Products, "total_count": resp.TotalCount})
}

func HandleGetProductReviews(ctx context.Context, c *app.RequestContext) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Fail(c, err.Error())
		return
	}
	req := rpc_product.GetProductReviewsReq{
		ProductId: int32(id),
	}
	resp, err := clients.ProductClient.GetProductReviews(context.Background(), &req, callopt.WithRPCTimeout(10*time.Second))
	if err != nil {
		utils.Fail(c, err.Error())
		return
	}
	utils.Success(c, utils.H{"reviews": resp.Reviews})
}

func HandlePostProductReviews(ctx context.Context, c *app.RequestContext) {
	var body struct {
		ProductID int    `json:"product_id"`
		UserID    int    `json:"user_id"`
		Rating    int    `json:"rating"`
		Comment   string `json:"comment"`
	}
	if err := c.Bind(&body); err != nil {
		utils.Fail(c, err.Error())
		return
	}
	userName, err := dao.GetUserName(tidb.DB, body.UserID)
	if err != nil {
		utils.Fail(c, err.Error())
		return
	}
	req := rpc_product.CreateProductReviewReq{
		ProductId: int32(body.ProductID),
		UserId:    int32(body.UserID),
		UserName:  userName,
		Rating:    int32(body.Rating),
		Comment:   body.Comment,
	}
	resp, err := clients.ProductClient.CreateProductReview(context.Background(), &req, callopt.WithRPCTimeout(5*time.Second))
	if err != nil {
		utils.Fail(c, err.Error())
		return
	}
	utils.Success(c, utils.H{"success": resp.Success})
}

func HandleCreatePromotion(ctx context.Context, c *app.RequestContext) {
	var body struct {
		Type         string `json:"type"`
		ProductId    string `json:"product_id"`
		DiscountRate string `json:"discount_rate"`
		CouponCode   string `json:"coupon_code"`
		CouponType   string `json:"coupon_type"`
		CouponValue  string `json:"coupon_value"`
		MinPurchase  string `json:"min_purchase"`
		UsageLimit   string `json:"usage_limit"`
		StartDate    string `json:"start_date"`
		EndDate      string `json:"end_date"`
	}
	if err := c.Bind(&body); err != nil {
		utils.Fail(c, err.Error())
		return
	}

	productId, _ := strconv.Atoi(body.ProductId)
	discountRate, _ := strconv.ParseFloat(body.DiscountRate, 64)
	couponValue, _ := strconv.ParseFloat(body.CouponValue, 64)
	minPurchase, _ := strconv.ParseFloat(body.MinPurchase, 64)
	usageLimit, _ := strconv.Atoi(body.UsageLimit)

	req := rpc_product.CreatePromotionReq{
		Type:         body.Type,
		ProductId:    int32(productId),
		DiscountRate: discountRate,
		CouponCode:   body.CouponCode,
		CouponType:   body.CouponType,
		CouponValue:  couponValue,
		MinPurchase:  minPurchase,
		UsageLimit:   int32(usageLimit),
		StartDate:    body.StartDate,
		EndDate:      body.EndDate,
	}
	resp, err := clients.ProductClient.CreatePromotion(context.Background(), &req, callopt.WithRPCTimeout(5*time.Second))
	if err != nil {
		utils.Fail(c, err.Error())
		return
	}
	utils.Success(c, utils.H{"success": resp.Success})
}

func HandleGetActivePromotions(ctx context.Context, c *app.RequestContext) {
	resp, err := clients.ProductClient.GetActivePromotions(context.Background(), callopt.WithRPCTimeout(5*time.Second))
	if err != nil {
		utils.Fail(c, err.Error())
		return
	}
	utils.Success(c, utils.H{"promotions": resp.Promotions})
}

func HandleDeletePromotion(ctx context.Context, c *app.RequestContext) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Fail(c, err.Error())
		return
	}
	req := rpc_product.DeletePromotionReq{
		Id: int32(id),
	}
	resp, err := clients.ProductClient.DeletePromotion(context.Background(), &req, callopt.WithRPCTimeout(5*time.Second))
	if err != nil {
		utils.Fail(c, err.Error())
		return
	}
	utils.Success(c, utils.H{"success": resp.Success})
}
