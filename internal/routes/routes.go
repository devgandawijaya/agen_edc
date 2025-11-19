package routes

import (
	"agen_edc/config"
	"agen_edc/internal/controllers"
	"agen_edc/internal/repositories"
	"agen_edc/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg *config.Config) *gin.Engine {
	db := config.NewGormDB(cfg)

	// Initialize repos
	userRepo := repositories.NewUserRepository(db)
	agentRepo := repositories.NewAgentRepo(db)
	acquisitionRepo := repositories.NewAcquisitionRepo(db)
	ownerRepo := repositories.NewOwnerRepo(db)
	businessProfileRepo := repositories.NewBusinessProfileRepo(db)
	uploadedDocumentRepo := repositories.NewUploadedDocumentRepo(db)
	bankInfoRepo := repositories.NewBankInfoRepo(db)
	signatureRepo := repositories.NewSignatureRepo(db)
	auditLogRepo := repositories.NewAuditLogRepo(db)

	// Initialize services
	userService := services.NewUserService(userRepo, cfg.JWTSecret)
	agentService := services.NewAgentService(agentRepo, auditLogRepo)
	acquisitionService := services.NewAcquisitionService(acquisitionRepo)
	ownerService := services.NewOwnerService(ownerRepo)
	businessProfileService := services.NewBusinessProfileService(businessProfileRepo)
	uploadedDocumentService := services.NewUploadedDocumentService(uploadedDocumentRepo)
	bankInfoService := services.NewBankInfoService(bankInfoRepo)
	signatureService := services.NewSignatureService(signatureRepo)
	auditLogService := services.NewAuditLogService(auditLogRepo)

	// Initialize controllers
	userCtrl := controllers.NewUserController(userService)
	agentCtrl := controllers.NewAgentController(agentService)
	acquisitionCtrl := controllers.NewAcquisitionController(acquisitionService)
	ownerCtrl := controllers.NewOwnerController(ownerService)
	businessProfileCtrl := controllers.NewBusinessProfileController(businessProfileService)
	uploadedDocumentCtrl := controllers.NewUploadedDocumentController(uploadedDocumentService)
	bankInfoCtrl := controllers.NewBankInfoController(bankInfoService)
	signatureCtrl := controllers.NewSignatureController(signatureService)
	auditLogCtrl := controllers.NewAuditLogController(auditLogService)

	r := gin.Default()
	SetupRoutes(r, userCtrl, agentCtrl, acquisitionCtrl, ownerCtrl, businessProfileCtrl, uploadedDocumentCtrl, bankInfoCtrl, signatureCtrl, auditLogCtrl)
	return r
}

func SetupRoutes(r *gin.Engine,
	userCtrl *controllers.UserController,
	agentCtrl *controllers.AgentController,
	acquisitionCtrl *controllers.AcquisitionController,
	ownerCtrl *controllers.OwnerController,
	businessProfileCtrl *controllers.BusinessProfileController,
	uploadedDocumentCtrl *controllers.UploadedDocumentController,
	bankInfoCtrl *controllers.BankInfoController,
	signatureCtrl *controllers.SignatureController,
	auditLogCtrl *controllers.AuditLogController) {

	api := r.Group("/api")
	{
		users := api.Group("/users")
		{
			users.POST("/register", userCtrl.RegisterUser)
			users.POST("/login", userCtrl.LoginUser)
			users.POST("/forgot-password", userCtrl.ForgotPasswordUser)
			users.POST("/reset-password", userCtrl.ResetPasswordUser)
			users.GET("", userCtrl.GetUsers)
			users.GET("/search", userCtrl.SearchUsers)
			users.GET("/:id", userCtrl.GetUser)
			users.PUT("/:id", userCtrl.UpdateUser)
			users.DELETE("/:id", userCtrl.DeleteUser)
		}

		agents := api.Group("/agents")
		{
			agents.POST("", agentCtrl.Create)
			agents.GET("/:id", agentCtrl.Get)
			agents.GET("/:id/full", agentCtrl.GetFull)
			agents.GET("", agentCtrl.Search)
			agents.PUT("/:id", agentCtrl.Update)
			agents.DELETE("/:id", agentCtrl.Delete)
			agents.POST("/:id/documents", agentCtrl.UploadDocuments)
		}

		acquisitions := api.Group("/acquisitions")
		{
			acquisitions.POST("", acquisitionCtrl.Create)
			acquisitions.GET("/:id", acquisitionCtrl.Get)
			acquisitions.GET("/agent/:agent_id", acquisitionCtrl.GetByAgentID)
			acquisitions.PUT("/:id", acquisitionCtrl.Update)
			acquisitions.DELETE("/:id", acquisitionCtrl.Delete)
			acquisitions.GET("", acquisitionCtrl.GetAll)
		}

		owners := api.Group("/owners")
		{
			owners.POST("", ownerCtrl.Create)
			owners.GET("/:id", ownerCtrl.Get)
			owners.GET("/agent/:agent_id", ownerCtrl.GetByAgentID)
			owners.PUT("/:id", ownerCtrl.Update)
			owners.DELETE("/:id", ownerCtrl.Delete)
			owners.GET("", ownerCtrl.GetAll)
		}

		businessProfiles := api.Group("/business_profiles")
		{
			businessProfiles.POST("", businessProfileCtrl.Create)
			businessProfiles.GET("/:id", businessProfileCtrl.Get)
			businessProfiles.GET("/agent/:agent_id", businessProfileCtrl.GetByAgentID)
			businessProfiles.PUT("/:id", businessProfileCtrl.Update)
			businessProfiles.DELETE("/:id", businessProfileCtrl.Delete)
			businessProfiles.GET("", businessProfileCtrl.GetAll)
		}

		uploadedDocuments := api.Group("/uploaded_documents")
		{
			uploadedDocuments.POST("", uploadedDocumentCtrl.Create)
			uploadedDocuments.GET("/:id", uploadedDocumentCtrl.Get)
			uploadedDocuments.GET("/agent/:agent_id", uploadedDocumentCtrl.GetByAgentID)
			uploadedDocuments.PUT("/:id", uploadedDocumentCtrl.Update)
			uploadedDocuments.DELETE("/:id", uploadedDocumentCtrl.Delete)
			uploadedDocuments.GET("", uploadedDocumentCtrl.GetAll)
		}

		bankInfos := api.Group("/bank_infos")
		{
			bankInfos.POST("", bankInfoCtrl.Create)
			bankInfos.GET("/:id", bankInfoCtrl.Get)
			bankInfos.GET("/agent/:agent_id", bankInfoCtrl.GetByAgentID)
			bankInfos.PUT("/:id", bankInfoCtrl.Update)
			bankInfos.DELETE("/:id", bankInfoCtrl.Delete)
			bankInfos.GET("", bankInfoCtrl.GetAll)
		}

		signatures := api.Group("/signatures")
		{
			signatures.POST("", signatureCtrl.Create)
			signatures.GET("/:id", signatureCtrl.Get)
			signatures.GET("/agent/:agent_id", signatureCtrl.GetByAgentID)
			signatures.PUT("/:id", signatureCtrl.Update)
			signatures.DELETE("/:id", signatureCtrl.Delete)
			signatures.GET("", signatureCtrl.GetAll)
		}

		auditLogs := api.Group("/audit_logs")
		{
			auditLogs.POST("", auditLogCtrl.Create)
			auditLogs.GET("/:id", auditLogCtrl.Get)
			auditLogs.GET("/agent/:agent_id", auditLogCtrl.GetByAgentID)
			auditLogs.PUT("/:id", auditLogCtrl.Update)
			auditLogs.DELETE("/:id", auditLogCtrl.Delete)
			auditLogs.GET("", auditLogCtrl.GetAll)
		}
	}
}
