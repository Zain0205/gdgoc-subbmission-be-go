package dto

import "time"

type RegisterInput struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type CreateSubmissionInput struct {
	SeriesID uint   `json:"series_id" validate:"required,gt=0"`
	FileURL  string `json:"file_url" validate:"required,url"`
}

type CreateScoreInput struct {
	SubmissionID uint   `json:"submission_id" validate:"required,gt=0"`
	Score        int    `json:"score" validate:"required,gte=0,lte=100"`
	Feedback     string `json:"feedback" validate:"max=1000"`
}

type SetVerificationCodeInput struct {
	Code string `json:"code" validate:"required,min=4,max=10"`
}

type VerifyCodeInput struct {
	Code string `json:"code" validate:"required"`
}

type LeaderboardResult struct {
	UserID     uint    `json:"user_id"`
	Name       string  `json:"name"`
	TotalScore float64 `json:"total_score"`
	Rank       int     `json:"rank"`
}

type UpdateUserRoleInput struct {
	Role string `json:"role" validate:"required,oneof=admin member"`
}

type CreateTrackInput struct {
	TrackName   string `json:"track_name" validate:"required,min=3,max=255"`
	Description string `json:"description" validate:"max=1000"`
	TrackType   string `json:"track_type" validate:"oneof=STUDY_JAM HACKATHON"`
}

type CreateSeriesInput struct {
	TrackID       uint      `json:"track_id" validate:"required,gt=0"`
	SeriesName    string    `json:"series_name" validate:"required,min=3,max=255"`
	Description   string    `json:"description" validate:"max=1000"`
	Deadline      time.Time `json:"deadline" validate:"required"`
	OrderIndex    int       `json:"order_index" validate:"gte=0"`
	IsCompetition bool      `json:"is_competition"`
}

type CreateAchievementTypeInput struct {
	Name string `json:"name" validate:"required,min=2,max=100"`
}

type UpdateAchievementTypeInput struct {
	Name string `json:"name" validate:"required,min=2,max=100"`
}

type CreateAchievementInput struct {
	Name              string `json:"name" validate:"required,min=3,max=255"`
	Description       string `json:"description" validate:"required,min=10,max=1000"`
	IconURL           string `json:"icon_url" validate:"omitempty,url"`
	AchievementTypeID uint   `json:"achievement_type_id" validate:"required,gt=0"`
}

type UpdateAchievementInput struct {
	Name              string `json:"name" validate:"omitempty,min=3,max=255"`
	Description       string `json:"description" validate:"omitempty,min=10,max=1000"`
	IconURL           string `json:"icon_url" validate:"omitempty,url"`
	AchievementTypeID uint   `json:"achievement_type_id" validate:"omitempty,gt=0"`
}

type AwardAchievementInput struct {
	UserID        uint `json:"user_id" validate:"required,gt=0"`
	AchievementID uint `json:"achievement_id" validate:"required,gt=0"`
}

type UpdateProfileInput struct {
	Name string `json:"name" validate:"required,min=2,max=50"`
}

type ChangePasswordInput struct {
	CurrentPassword string `json:"current_password" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required,min=6"`
}
