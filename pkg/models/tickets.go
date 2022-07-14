package models

import "time"

type Ticket struct {
	CcEmails        []string    `json:"cc_emails,omitempty"`
	FwdEmails       []string    `json:"fwd_emails,omitempty"`
	ReplyCcEmails   []string    `json:"reply_cc_emails,omitempty"`
	TicketCcEmails  []string    `json:"ticket_cc_emails,omitempty"`
	FrEscalated     bool        `json:"fr_escalated,omitempty"`
	Spam            bool        `json:"spam,omitempty"`
	EmailConfigID   int         `json:"email_config_id,omitempty"`
	GroupID         int64       `json:"group_id,omitempty"`
	Email           string      `json:"email,omitempty"`
	Priority        interface{} `json:"priority,omitempty"`
	RequesterID     int64       `json:"requester_id,omitempty"`
	ResponderID     int         `json:"responder_id,omitempty"`
	Source          interface{} `json:"source,omitempty"`
	CompanyID       int         `json:"company_id,omitempty"`
	Status          interface{} `json:"status,omitempty"`
	Subject         string      `json:"subject,omitempty"`
	AssociationType interface{} `json:"association_type,omitempty"`
	SupportEmail    interface{} `json:"support_email,omitempty"`
	ToEmails        []string    `json:"to_emails,omitempty"`
	ProductID       int         `json:"product_id,omitempty"`
	ID              int         `json:"id,omitempty"`
	Type            string      `json:"type,omitempty"`
	DueBy           *time.Time  `json:"due_by,omitempty"`
	FrDueBy         *time.Time  `json:"fr_due_by,omitempty"`
	IsEscalated     bool        `json:"is_escalated,omitempty"`
	Description     string      `json:"description,omitempty"`
	DescriptionText string      `json:"description_text,omitempty"`
	CustomFields    struct {
	} `json:"custom_fields,omitempty"`
	CreatedAt              *time.Time    `json:"created_at,omitempty"`
	UpdatedAt              *time.Time    `json:"updated_at,omitempty"`
	AssociatedTicketsCount interface{}   `json:"associated_tickets_count,omitempty"`
	Attachments            []interface{} `json:"attachments,omitempty"`
	SourceAdditionalInfo   interface{}   `json:"source_additional_info,omitempty"`
	TwitterId              string        `json:"twitter_id,omitempty"`
	Tags                   []string      `json:"tags,omitempty"`
	Requester              struct {
		ID     int64  `json:"id,omitempty"`
		Name   string `json:"name,omitempty"`
		Email  string `json:"email,omitempty"`
		Mobile int64  `json:"mobile,omitempty"`
		Phone  string `json:"phone,omitempty"`
	} `json:"requester,omitempty"`
	Stats struct {
		AgentRespondedAt     interface{} `json:"agent_responded_at,omitempty"`
		RequesterRespondedAt interface{} `json:"requester_responded_at,omitempty"`
		FirstRespondedAt     interface{} `json:"first_responded_at,omitempty"`
		StatusUpdatedAt      *time.Time  `json:"status_updated_at,omitempty"`
		ReopenedAt           interface{} `json:"reopened_at,omitempty"`
		ResolvedAt           interface{} `json:"resolved_at,omitempty"`
		ClosedAt             interface{} `json:"closed_at,omitempty"`
		PendingSince         interface{} `json:"pending_since,omitempty"`
	} `json:"stats,omitempty"`
	Company struct {
	} `json:"company,omitempty"`
	NrDueBy     interface{} `json:"nr_due_by,omitempty"`
	NrEscalated bool        `json:"nr_escalated,omitempty"`
}

type SearchResults struct {
	Results []*Ticket `json:"results,omitempty"`
	Total   int       `json:"total,omitempty"`
}

type OpenTickets struct {
	OpenTickets []*Ticket `json:"tickets,omitempty"`
	Total       int       `json:"total,omitempty"`
}

type PendingTickets struct {
	PendingTickets []*Ticket `json:"tickets,omitempty"`
	Total          int       `json:"total,omitempty"`
}

type ResolvedTickets struct {
	ResolvedTickets []*Ticket `json:"tickets,omitempty"`
	Total           int       `json:"total,omitempty"`
}

type ClosedTickets struct {
	ClosedTickets []*Ticket `json:"tickets,omitempty"`
	Total         int       `json:"total,omitempty"`
}

type AllTickets struct {
	OpenTickets     OpenTickets     `json:"open_tickets,omitempty"`
	PendingTickets  PendingTickets  `json:"pending_tickets,omitempty"`
	ResolvedTickets ResolvedTickets `json:"resolved_tickets,omitempty"`
	ClosedTickets   ClosedTickets   `json:"closed_tickets,omitempty"`
}

type Status int

const (
	Open Status = iota + 2
	Pending
	Resolved
	Closed
)

func (s Status) String() string {
	switch s {
	case Open:
		return "open"
	case Pending:
		return "pending"
	case Resolved:
		return "resolved"
	case Closed:
		return "closed"
	}
	return "unknown"
}

type Priority int

const (
	Low Priority = iota + 1
	Medium
	High
	Urgent
)

func (p Priority) String() string {
	switch p {
	case Low:
		return "low"
	case Medium:
		return "medium"
	case High:
		return "high"
	case Urgent:
		return "urgent"
	}
	return "unknown"
}

type Source int

const (
	Email Source = iota + 1
	Portal
	Phone
	Chat
	FeedBackWidget
	OutboundEmail
)

func (s Source) String() string {
	switch s {
	case Email:
		return "email"
	case Portal:
		return "portal"
	case Phone:
		return "phone"
	case Chat:
		return "chat"
	case FeedBackWidget:
		return "feedback widget"
	case OutboundEmail:
		return "outbound email"
	}
	return "unknown"
}
