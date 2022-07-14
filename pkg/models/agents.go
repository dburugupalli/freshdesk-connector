package models

import "time"

type Agent struct {
	Available      bool       `json:"available,omitempty"`
	Occasional     bool       `json:"occasional,omitempty"`
	ID             int64      `json:"id,omitempty"`
	TicketScope    int        `json:"ticket_scope,omitempty"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
	LastActiveAt   *time.Time `json:"last_active_at,omitempty"`
	AvailableSince *time.Time `json:"available_since,omitempty"`
	Type           string     `json:"type,omitempty"`
	Contact        struct {
		Active      bool        `json:"active,omitempty"`
		Email       string      `json:"email,omitempty"`
		JobTitle    string      `json:"job_title,omitempty"`
		Language    string      `json:"language,omitempty"`
		LastLoginAt interface{} `json:"last_login_at,omitempty"`
		Mobile      int64       `json:"mobile,omitempty"`
		Name        string      `json:"name,omitempty"`
		Phone       string      `json:"phone,omitempty"`
		TimeZone    string      `json:"time_zone,omitempty"`
		CreatedAt   *time.Time  `json:"created_at,omitempty"`
		UpdatedAt   *time.Time  `json:"updated_at,omitempty"`
	} `json:"contact,omitempty"`
	Signature string `json:"signature,omitempty"`
}
