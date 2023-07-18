package main

import (
	"github.com/gin-contrib/cors"
	"project/model"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user	=postgres password=teptom0792. dbname=test port=5432 sslmode=prefer connect_timeout=10 TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:8081"},
		AllowMethods:  []string{"PUT", "POST", "GET", "DELETE"},
		AllowHeaders:  []string{"Content-Type", "Content-Length", "Accept-Encoding"},
		ExposeHeaders: []string{"Content-Length", "Content-Type"},
	}))
	api := router.Group("/api")
	{
		api.GET("getUser/", func(ctx *gin.Context) {
			var user []model.User
			if err := db.Table("users").Find(&user).Error; err != nil {
				ctx.JSON(200, gin.H{
					"err": err.Error(),
				})
				return
			}
			ctx.JSON(200, gin.H{
				"data": user,
			})
		})
		api.GET("getUser/Internal", func(ctx *gin.Context) {
			var user []model.User
			if err := db.Where("status = ?", "Internal").Table("users").Find(&user).Error; err != nil {
				ctx.JSON(200, gin.H{
					"err": err.Error(),
				})
				return
			}
			ctx.JSON(200, gin.H{
				"data": user,
			})
		})
		api.GET("getUser/Pending", func(ctx *gin.Context) {
			var user []model.User
			if err := db.Where("status = ?", "Pending").Table("users").Find(&user).Error; err != nil {
				ctx.JSON(200, gin.H{
					"err": err.Error(),
				})
				return
			}
			ctx.JSON(200, gin.H{
				"data": user,
			})
		})
		api.GET("getUser/:id", func(ctx *gin.Context) {
			id := ctx.Params.ByName("id")
			var user model.User
			if err := db.Where("id = ?", id).Table("users").First(&user).Error; err != nil {
				ctx.JSON(200, gin.H{
					"err": err.Error(),
				})
				return
			}
			ctx.JSON(200, gin.H{
				"data": user,
			})

		})
		api.GET("getCompanyInformation/:control_number", func(ctx *gin.Context) {
			control_number := ctx.Params.ByName("control_number")
			var company model.Company
			if err := db.Where("control_number = ?", control_number).Table("company").First(&company).Error; err != nil {
				ctx.JSON(200, gin.H{
					"err": err.Error(),
				})
				return
			}
			ctx.JSON(200, gin.H{
				"data": company,
			})

		})
		api.POST("createUser", func(ctx *gin.Context) {
			var user model.User
			ctx.ShouldBind(&user)

			if err := db.Table("users").Create(&user).Error; err != nil {
				ctx.JSON(200, gin.H{
					"err": err.Error(),
				})
				return
			}
			ctx.JSON(200, gin.H{
				"data": user,
			})
		})
		api.POST("createCompany", func(ctx *gin.Context) {
			var company model.Company
			ctx.ShouldBind(&company)

			if err := db.Table("company").Create(&company).Error; err != nil {
				ctx.JSON(200, gin.H{
					"err": err.Error(),
				})
				return
			}
			ctx.JSON(200, gin.H{
				"data": company,
			})
		})
		api.POST("createLogin", func(ctx *gin.Context) {
			var login model.Login
			ctx.ShouldBind(&login)

			if err := db.Table("login").Create(&login).Error; err != nil {
				ctx.JSON(200, gin.H{
					"err": err.Error(),
				})
				return
			}
			ctx.JSON(200, gin.H{
				"data": login,
			})
		})
		api.PUT("updatePassword/:id", func(ctx *gin.Context) {
			id := ctx.Params.ByName("id")
			var login model.Login
			if err := db.Where("user_id = ?", id).Table("login").First(&login).Error; err != nil {
				ctx.JSON(200, gin.H{
					"err": err.Error(),
				})
				return
			}
			ctx.ShouldBind(&login)
			db.Save(&login)
			ctx.JSON(200, gin.H{
				"data": login,
			})
		})
		api.PUT("updateUser/:id", func(ctx *gin.Context) {
			id := ctx.Params.ByName("id")
			var user model.User
			if err := db.Where("id = ?", id).Table("users").First(&user).Error; err != nil {
				ctx.JSON(200, gin.H{
					"err": err.Error(),
				})
				return
			}
			ctx.ShouldBind(&user)
			db.Save(&user)
			ctx.JSON(200, gin.H{
				"data": user,
			})
		})
		api.PUT("updateCompany/:control_number", func(ctx *gin.Context) {
			control_number := ctx.Param("control_number")
			var company model.CompanyBody
			ctx.ShouldBind(&company)
			db.Where("control_number = ?", control_number).Table("company").Save(&company.Body)
			ctx.JSON(200, gin.H{
				"data": company.Body,
			})
		})
		api.DELETE("deleteUser/:id", func(ctx *gin.Context) {
			id := ctx.Params.ByName("id")
			var user model.User
			if err := db.Where("id = ?", id).Table("users").Delete(&user).Error; err != nil {
				ctx.JSON(200, gin.H{
					"err": err.Error(),
				})
				return
			}
			ctx.JSON(200, gin.H{
				"data": user,
			})
		})

	}
	router.Run(":8080")
}
