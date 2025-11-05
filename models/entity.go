package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	Name         string `json:"name"`
	Email        string `json:"email" gorm:"unique;not null"`
	PasswordHash string `json:"-" gorm:"not null"`
	Role         string `json:"role" gorm:"type:enum('admin','member');default:'member'"`

	// Relasi
	Submissions []Submission `json:"submissions,omitempty" gorm:"foreignKey:UserID"`
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

	TrackName   string `json:"track_name"`
	Description string `json:"description" gorm:"type:text"`
	CreatedByID uint   `json:"created_by_id"`

	// Relasi
	CreatedBy User     `json:"created_by" gorm:"foreignKey:CreatedByID"`
	Series    []Series `json:"series,omitempty" gorm:"foreignKey:TrackID"`
}

type Series struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`

	TrackID          uint      `json:"track_id"`
	SeriesName       string    `json:"series_name"`
	Description      string    `json:"description" gorm:"type:text"`
	Deadline         time.Time `json:"deadline"`
	OrderIndex       int       `json:"order_index"`
	VerificationCode string    `json:"-" gorm:"varchar(10);null"`
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

	UserID   uint   `json:"user_id"`
	SeriesID uint   `json:"series_id"`
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
