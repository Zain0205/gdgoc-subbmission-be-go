package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID           uint           `json:"id" gorm:"primarykey"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
	Name         string         `json:"name"`
	Email        string         `json:"email" gorm:"unique;not null"`
	PasswordHash string         `json:"-" gorm:"not null"`
	Role         string         `json:"role" gorm:"type:enum('admin','member');default:'member'"`
	AvatarURL    string         `json:"avatar_url" gorm:"size:255"`
	Submissions  []Submission   `json:"submissions,omitempty" gorm:"foreignKey:UserID"`
}

func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
}

type Track struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	TrackName   string `json:"track_name" gorm:"unique;size:255"`
	Description string `json:"description" gorm:"type:text"`
	CreatedByID uint   `json:"created_by_id"`

	TrackType string `json:"track_type" gorm:"size:50;default:'STUDY_JAM'"`

	// Relasi
	CreatedBy User     `json:"created_by" gorm:"foreignKey:CreatedByID"`
	Series    []Series `json:"series,omitempty" gorm:"foreignKey:TrackID"`
}

type Series struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	TrackID    uint   `json:"track_id" gorm:"uniqueIndex:idx_track_series"`
	SeriesName string `json:"series_name" gorm:"uniqueIndex:idx_track_series;size:255"`

	IsCompetition bool `json:"is_competition" gorm:"default:false"`

	Description      string    `json:"description" gorm:"type:text"`
	Deadline         time.Time `json:"deadline"`
	OrderIndex       int       `json:"order_index"`
	VerificationCode string    `json:"verification_code" gorm:"varchar(10);null"`
}

type UserSeriesVerification struct {
	UserID     uint      `json:"user_id" gorm:"primaryKey"`
	SeriesID   uint      `json:"series_id" gorm:"primaryKey"`
	VerifiedAt time.Time `json:"verified_at"`

	User   User   `json:"user" gorm:"foreignKey:UserID"`
	Series Series `json:"series" gorm:"foreignKey:SeriesID"`
}

type Submission struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	UserID   uint   `json:"user_id" gorm:"uniqueIndex:idx_user_series"`
	SeriesID uint   `json:"series_id" gorm:"uniqueIndex:idx_user_series"`
	FileURL  string `json:"file_url" gorm:"type:text"`
	Score    int    `json:"score"`
	Feedback string `json:"feedback" gorm:"type:text"`

	// Relasi
	User   User   `json:"user" gorm:"foreignKey:UserID"`
	Series Series `json:"series" gorm:"foreignKey:SeriesID"`
}

type Leaderboard struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	TrackID    uint    `json:"track_id"`
	UserID     uint    `json:"user_id"`
	TotalScore float64 `json:"total_score"`
	Rank       int     `json:"rank"`

	User  User  `json:"user" gorm:"foreignKey:UserID"`
	Track Track `json:"track" gorm:"foreignKey:TrackID"`
}

type AchievementType struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Name      string    `json:"name" gorm:"size:100;unique;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Achievement struct {
	ID                uint      `json:"id" gorm:"primarykey"`
	Name              string    `json:"name" gorm:"size:255;unique;not null"`
	Description       string    `json:"description" gorm:"type:text"`
	IconURL           string    `json:"icon_url" gorm:"type:text"`
	AchievementTypeID uint      `json:"achievement_type_id"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`

	// Relasi
	Type AchievementType `json:"type" gorm:"foreignKey:AchievementTypeID"`
}

type UserAchievement struct {
	UserID        uint      `json:"user_id" gorm:"primaryKey"`
	AchievementID uint      `json:"achievement_id" gorm:"primaryKey"`
	EarnedAt      time.Time `json:"earned_at"`

	// Relasi
	User        User        `json:"user" gorm:"foreignKey:UserID"`
	Achievement Achievement `json:"achievement" gorm:"foreignKey:AchievementID"`
}

type Notification struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	UserID    uint      `json:"user_id" gorm:"index"`
	Title     string    `json:"title" gorm:"type:varchar(255);not null"`
	Message   string    `json:"message" gorm:"type:text"`
	Type      string    `json:"type" gorm:"default:'grade'"`
	RelatedID *uint     `json:"related_id"`
	IsRead    bool      `json:"is_read" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at"`
}
