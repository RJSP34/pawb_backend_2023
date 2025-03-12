package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
	"pawb_backend/config"
	"pawb_backend/controller"
	"pawb_backend/middleware"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	config.ConnectDB()
	defer config.CloseDb()

	config.SetupRedisClient()

	router := gin.Default()
	router.Use(CORSMiddleware())

	v1 := router.Group("/api/v1")
	{
		v1.GET("", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello World")
		})

		auth := v1.Group("/auth")
		{
			auth.POST("/login", controller.Login)
			auth.POST("/register", controller.Register)
			auth.DELETE("/logout", middleware.Authorized(), controller.Logout)
		}

		user := v1.Group("/user", middleware.Authorized("user"))
		{
			user.POST("/images/submit", controller.SubmitImage)
			user.GET("/images/:id", controller.GetImage)
			user.GET("/images", controller.GetMyImages)
			user.PATCH("/images", controller.UpdateDescriptionImage)
			user.DELETE("/images/:id", controller.DeleteImage)
			user.PUT("/clinicians", controller.UpdateCliniciansPermissions)
			user.GET("/allowed_clinicians", controller.GetAllowedClinicians)
			user.GET("/clinicians", controller.GetClinicians)
		}

		bodyParts := v1.Group("/body_parts", middleware.Authorized())
		{
			bodyParts.GET("", controller.GetBodyParts)
		}

		clinician := v1.Group("/clinicians", middleware.Authorized("clinician"))
		{
			clinician.GET("/myfeedback", controller.GetFeedbackByClinician)
			clinician.GET("/patient/:id", controller.GetAllPatientsImages)
			clinician.GET("/patient", controller.GetAllAuthorizedPatients)
			clinician.GET("/image/:id", controller.GetImageByClinician)
			clinician.GET("/images", controller.GetAllImagesByClinician)
			clinician.PATCH("/feedback", controller.UpdateFeedback)
			clinician.DELETE("/feedback/:id", controller.RemoveFeedback)
			clinician.POST("/feedback", controller.SubmitFeedback)
		}

		feedback := v1.Group("/feedback", middleware.Authorized())
		{
			feedback.GET("/image/:id", controller.GetFeedbackByImage)
		}
		profile := v1.Group("/profile", middleware.Authorized("user", "clinician"))
		{
			profile.GET("", controller.Profile)
		}
	}

	err = router.Run(":3000")
	if err != nil {
		return
	}
}
