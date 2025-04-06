package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents the user data in the system
type User struct {
	ID                primitive.ObjectID    `bson:"_id,omitempty" json:"id"`
	Username          string                `bson:"username" json:"username"`                   // Mandatory field
	Email             string                `bson:"email,omitempty" json:"email,omitempty"`     // Mandatory field
	Password          string                `bson:"password" json:"password"`                   // Mandatory field
	Profile           *UserProfile          `bson:"profile,omitempty" json:"profile,omitempty"` // Optional field, but recommended
	Phone             string                `bson:"phone,omitempty" json:"phone,omitempty"`
	Preferences       UserPreferences       `bson:"preferences,omitempty" json:"preferences,omitempty"`
	Roles             []string              `bson:"roles,omitempty" json:"roles,omitempty"`
	Activity          UserActivity          `bson:"activity,omitempty" json:"activity,omitempty"`
	SocialConnections UserSocialConnections `bson:"socialConnections,omitempty" json:"socialConnections,omitempty"`
	Subscriptions     []UserSubscription    `bson:"subscriptions,omitempty" json:"subscriptions,omitempty"`
	Settings          UserSettings          `bson:"settings,omitempty" json:"settings,omitempty"`
	AuditLogs         []UserAuditLog        `bson:"auditLogs,omitempty" json:"auditLogs,omitempty"`
}

// UserProfile represents personal information of the user
type UserProfile struct {
	FirstName      string `bson:"firstName,omitempty" json:"firstName,omitempty"`
	LastName       string `bson:"lastName,omitempty" json:"lastName,omitempty"`
	Dob            string `bson:"dob,omitempty" json:"dob,omitempty"`
	ProfilePicture string `bson:"profilePicture,omitempty" json:"profilePicture,omitempty"`
	Bio            string `bson:"bio,omitempty" json:"bio,omitempty"`
	Location       string `bson:"location,omitempty" json:"location,omitempty"`
}

// DefaultProfile returns a profile with default values
func DefaultProfile() *UserProfile {
	return &UserProfile{
		FirstName:      "Unknown",
		LastName:       "Unknown",
		Dob:            "",
		ProfilePicture: "",
		Bio:            "",
		Location:       "",
	}
}

// UserPreferences stores the user's settings and preferences
type UserPreferences struct {
	Language        string              `bson:"language,omitempty" json:"language,omitempty"`
	Theme           string              `bson:"theme,omitempty" json:"theme,omitempty"`
	Notifications   UserNotifications   `bson:"notifications,omitempty" json:"notifications,omitempty"`
	PrivacySettings UserPrivacySettings `bson:"privacySettings,omitempty" json:"privacySettings,omitempty"`
}

// UserNotifications represents the user's notification preferences
type UserNotifications struct {
	EmailNotifications bool `bson:"emailNotifications,omitempty" json:"emailNotifications,omitempty"`
	SMSNotifications   bool `bson:"smsNotifications,omitempty" json:"smsNotifications,omitempty"`
}

// UserPrivacySettings holds privacy-related preferences
type UserPrivacySettings struct {
	ProfileVisibility string `bson:"profileVisibility,omitempty" json:"profileVisibility,omitempty"`
	ActivityStatus    string `bson:"activityStatus,omitempty" json:"activityStatus,omitempty"`
	ShareLocation     bool   `bson:"shareLocation,omitempty" json:"shareLocation,omitempty"`
}

// UserActivity stores the user's login and activity history
type UserActivity struct {
	LastLogin    time.Time           `bson:"lastLogin,omitempty" json:"lastLogin,omitempty"`
	LastActivity time.Time           `bson:"lastActivity,omitempty" json:"lastActivity,omitempty"`
	LoginHistory []LoginHistoryEntry `bson:"loginHistory,omitempty" json:"loginHistory,omitempty"`
}

// LoginHistoryEntry represents an individual login entry
type LoginHistoryEntry struct {
	LoginTime time.Time `bson:"loginTime,omitempty" json:"loginTime,omitempty"`
	IP        string    `bson:"ip,omitempty" json:"ip,omitempty"`
}

// UserSocialConnections represents user's social relationships
type UserSocialConnections struct {
	Friends   []string `bson:"friends,omitempty" json:"friends,omitempty"`
	Following []string `bson:"following,omitempty" json:"following,omitempty"`
	Followers []string `bson:"followers,omitempty" json:"followers,omitempty"`
}

// UserSubscription stores the user's subscription details
type UserSubscription struct {
	SubscriptionType string    `bson:"subscriptionType,omitempty" json:"subscriptionType,omitempty"`
	StartDate        time.Time `bson:"startDate,omitempty" json:"startDate,omitempty"`
	EndDate          time.Time `bson:"endDate,omitempty" json:"endDate,omitempty"`
	Status           string    `bson:"status,omitempty" json:"status,omitempty"`
}

// UserSettings represents additional settings like two-factor authentication
type UserSettings struct {
	TwoFactorEnabled bool   `bson:"twoFactorEnabled,omitempty" json:"twoFactorEnabled,omitempty"`
	Theme            string `bson:"theme,omitempty" json:"theme,omitempty"`
}

// UserAuditLog stores significant actions performed by the user
type UserAuditLog struct {
	Action    string    `bson:"action,omitempty" json:"action,omitempty"`
	Timestamp time.Time `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	IP        string    `bson:"ip,omitempty" json:"ip,omitempty"`
	Details   string    `bson:"details,omitempty" json:"details,omitempty"`
}
