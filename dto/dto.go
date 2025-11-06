package dto

import "time"

type RegisterInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
type CreateSubmissionInput struct {
	SeriesID uint   `json:"series_id" binding:"required"`
	FileURL  string `json:"file_url" binding:"required"`
}
type CreateScoreInput struct {
	SubmissionID uint   `json:"submission_id" binding:"required"`
	Score        int    `json:"score" binding:"required,min=0,max=100"`
	Feedback     string `json:"feedback"`
}
type SetVerificationCodeInput struct {
	Code string `json:"code" binding:"required,min=4,max=10"`
}
type VerifyCodeInput struct {
	Code string `json:"code" binding:"required"`
}
type LeaderboardResult struct {
	UserID     uint    `json:"user_id"`
	Name       string  `json:"name"`
	TotalScore float64 `json:"total_score"`
	Rank       int     `json:"rank"`
}
type UpdateUserRoleInput struct {
	Role string `json:"role" binding:"required,oneof=admin member"`
}

type CreateTrackInput struct {
	TrackName   string `json:"track_name" binding:"required"`
	Description string `json:"description"`
	// Admin can set this, e.g., "STUDY_JAM" or "HACKATHON"
	TrackType string `json:"track_type"`
}

type CreateSeriesInput struct {
	TrackID     uint      `json:"track_id" binding:"required"`
	SeriesName  string    `json:"series_name" binding:"required"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline" binding:"required"`
	OrderIndex  int       `json:"order_index"`

	// Admin sets this: false for "Series", true for "Mini-Competition"
	IsCompetition bool `json:"is_competition"`
}

type CreateAchievementTypeInput struct {
	Name string `json:"name" binding:"required"`
}

type UpdateAchievementTypeInput struct {
	Name string `json:"name" binding:"required"`
}

type CreateAchievementInput struct {
	Name              string `json:"name" binding:"required"`
	Description       string `json:"description" binding:"required"`
	IconURL           string `json:"icon_url"`
	AchievementTypeID uint   `json:"achievement_type_id" binding:"required"`
}

type UpdateAchievementInput struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	IconURL           string `json:"icon_url"`
	AchievementTypeID uint   `json:"achievement_type_id"`
}

type AwardAchievementInput struct {
	UserID        uint `json:"user_id" binding:"required"`
	AchievementID uint `json:"achievement_id" binding:"required"`
}

