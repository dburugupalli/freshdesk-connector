package models

import "time"

type Group struct {
	ID                              int64      `json:"id"`
	Name                            string     `json:"name"`
	Description                     string     `json:"description"`
	EscalateTo                      int        `json:"escalate_to"`
	UnassignedFor                   string     `json:"unassigned_for"`
	AgentIds                        []int64    `json:"agent_ids"`
	CreatedAt                       *time.Time `json:"created_at"`
	UpdatedAt                       *time.Time `json:"updated_at"`
	BusinessCalendarID              int        `json:"business_calendar_id"`
	Type                            string     `json:"type"`
	AllowAgentsToChangeAvailability bool       `json:"allow_agents_to_change_availability"`
	AutomaticAgentAssignment        struct {
		Enabled bool `json:"enabled"`
	} `json:"automatic_agent_assignment"`
}
