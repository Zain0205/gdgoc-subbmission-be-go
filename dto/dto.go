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

type CreateTrackInput struct {
	TrackName   string `json:"track_name" binding:"required"`
	Description string `json:"description"`
}

type CreateSeriesInput struct {
	TrackID     uint      `json:"track_id" binding:"required"`
	SeriesName  string    `json:"series_name" binding:"required"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline" binding:"required"`
	OrderIndex  int       `json:"order_index"`
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
