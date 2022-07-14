package app

import (
	httpSwagger "github.com/swaggo/http-swagger"
)

func (a *App) initializeRoutes() {
	var apiV1 = "/api/v1"

	// tickets routes
	a.Router.HandleFunc(apiV1+"/tickets", a.getAllTickets).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/tickets/{id}", a.getSingleTicket).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/tickets", a.createTicket).Methods("POST", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/tickets/{id}", a.deleteTicket).Methods("DELETE", "OPTIONS")

	// search
	a.Router.HandleFunc(apiV1+"/search", a.searchTickets).Methods("GET", "OPTIONS")

	// stats
	a.Router.HandleFunc(apiV1+"/stats/agent/{agent_id}", a.getStatsForAgent).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/stats/group/{group_id}", a.getStatsForGroup).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/stats/date/{created_at}", a.getStatsForTicketsSinceDate).Methods("GET", "OPTIONS")

	// agents
	a.Router.HandleFunc(apiV1+"/agents", a.getAgents).Methods("GET", "OPTIONS")

	// groups
	a.Router.HandleFunc(apiV1+"/groups", a.getGroups).Methods("GET", "OPTIONS")

	// accounts
	a.Router.HandleFunc(apiV1+"/account", a.getAccounts).Methods("GET", "OPTIONS")

	// Swagger
	a.Router.PathPrefix("/docs").Handler(httpSwagger.WrapHandler)
}
