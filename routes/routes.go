package routes

import (
	"github.com/Zain0205/gdgoc-subbmission-be-go/controllers"
	"github.com/Zain0205/gdgoc-subbmission-be-go/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	api := r.Group("/")
	{
		// 1. Auth (Public)
		auth := api.Group("/auth")
		{
			auth.POST("/register", controllers.Register)
			auth.POST("/login", controllers.Login)
		}

		// 2. Public Routes (All authenticated users)
		public := api.Group("/")
		public.Use(middleware.AuthMiddleware())
		{
			public.GET("/tracks", controllers.GetAllTracks)
			public.GET("/tracks/:id", controllers.GetTrackWithSeries)
			public.GET("/leaderboard/track/:trackId", controllers.GetLeaderboardByTrack)
		}

		// 3. Member Routes
		member := api.Group("/member")
		member.Use(middleware.AuthMiddleware(), middleware.RoleMiddleware("member"))
		{
			member.POST("/submissions", controllers.CreateSubmission)
			member.POST("/series/:id/verify", controllers.VerifySeriesCode)

			member.GET("/me/achievements", controllers.GetMyAchievements)
		}

		// 4. Admin Routes
		admin := api.Group("/admin")
		admin.Use(middleware.AuthMiddleware(), middleware.RoleMiddleware("admin"))
		{
			admin.POST("/tracks", controllers.CreateTrack)
			admin.POST("/series", controllers.CreateSeries)
			admin.PATCH("/series/:id/code", controllers.SetSeriesVerificationCode)
			admin.GET("/submissions/series/:seriesId", controllers.GetSubmissionsBySeries)
			admin.POST("/submissions/grade", controllers.GradeSubmission)
			admin.PATCH("/users/:id/role", controllers.SetUserRole)

			admin.POST("/achievement-types", controllers.CreateAchievementType)
			admin.GET("/achievement-types", controllers.GetAchievementTypes)

			admin.POST("/achievements", controllers.CreateAchievement)
			admin.GET("/achievements", controllers.GetAchievements)
			admin.PUT("/achievements/:id", controllers.UpdateAchievement)

			admin.POST("/achievements/award", controllers.AwardAchievementToUser)
			admin.POST("/achievements/revoke", controllers.RevokeAchievementFromUser)
		}
	}

	return r
}

