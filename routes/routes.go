package routes

import (
	"psy-service/controllers"
	"psy-service/middlewares"
	"psy-service/repository"
	"psy-service/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Initialize repositories and services
	calendarRepo := &repository.CalendarRepository{}
	calendarService := &services.CalendarService{Repo: calendarRepo}
	calendarController := &controllers.CalendarController{Service: calendarService}

	patientRepo := &repository.PatientRepository{}
	patientService := &services.PatientService{Repo: patientRepo}
	patientController := &controllers.PatientController{Service: patientService}

	paymentRepo := &repository.PaymentRepository{}
	paymentService := &services.PaymentService{Repo: paymentRepo}
	paymentController := &controllers.PaymentController{Service: paymentService}

	authService := &services.AuthService{}
	authController := &controllers.AuthController{Service: authService}

	// Public Routes
	router.POST("/register", authController.Register)
	router.POST("/login", authController.Login)

	// Secure Routes
	api := router.Group("/api")
	api.Use(middlewares.JwtVerify())

	// Calendar Routes
	api.POST("/calendar", calendarController.AddCalendarEvent)
	api.PUT("/calendar/:id", calendarController.UpdateCalendarEvent)
	api.DELETE("/calendar/:id", calendarController.DeleteCalendarEvent)
	api.GET("/calendar", calendarController.GetCalendarEvents)

	// Patient Routes
	api.POST("/patients", patientController.CreatePatient)
	api.PUT("/patients/:id", patientController.UpdatePatient)
	api.DELETE("/patients/:id", patientController.DeletePatient)
	api.GET("/patients", patientController.GetPatients)

	// Payment Routes
	api.POST("/payments", paymentController.CreatePayment)
	api.PUT("/payments/:id", paymentController.UpdatePayment)
	api.DELETE("/payments/:id", paymentController.DeletePayment)
	api.GET("/payments", paymentController.GetPayments)

	return router
}
