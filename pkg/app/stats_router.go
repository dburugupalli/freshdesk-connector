package app

import (
	"github.com/gorilla/mux"
	"github.com/kosha/freshdesk-connector/pkg/httpclient"
	"github.com/kosha/freshdesk-connector/pkg/models"
	"net/http"
)

// getStatsForAgent godoc
// @Summary Get ticket statistics for a particular agent
// @Description Get ticket statistics for a particular agent id
// @Tags statistics
// @Accept  json
// @Produce  json
// @Param agent_id path string true "Enter agent id"
// @Param page query string true "Page number"
// @Success 200 {object} models.AllTickets
// @Router /api/v1/stats/agent/{agent_id} [get]
func (a *App) getStatsForAgent(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	var allTickets models.AllTickets

	var openTickets []*models.Ticket
	var pendingTickets []*models.Ticket
	var resolvedTickets []*models.Ticket
	var closedTickets []*models.Ticket

	vars := mux.Vars(r)
	agentId := vars["agent_id"]
	page := r.FormValue("page")

	t := httpclient.GetTicketsPerAgent(a.Cfg.GetFreshDeskURL(), a.Cfg.GetApiKey(), agentId, page)
	for _, ticket := range t.Results {

		ticket.Status = ticket.Status.(float64)

		if ticket.Status == float64(2) {
			openTickets = append(openTickets, ticket)
		} else if ticket.Status == float64(3) {
			pendingTickets = append(pendingTickets, ticket)
		} else if ticket.Status == float64(4) {
			resolvedTickets = append(resolvedTickets, ticket)
		} else if ticket.Status == float64(5) {
			closedTickets = append(closedTickets, ticket)
		}
	}

	var oT models.OpenTickets
	var pT models.PendingTickets
	var rT models.ResolvedTickets
	var cT models.ClosedTickets

	if len(openTickets) > 0 {
		oT = models.OpenTickets{OpenTickets: openTickets, Total: len(openTickets)}
	}
	if len(pendingTickets) > 0 {

		pT = models.PendingTickets{PendingTickets: pendingTickets, Total: len(pendingTickets)}
	}
	if len(resolvedTickets) > 0 {
		rT = models.ResolvedTickets{ResolvedTickets: resolvedTickets, Total: len(resolvedTickets)}
	}
	if len(closedTickets) > 0 {
		cT = models.ClosedTickets{ClosedTickets: closedTickets, Total: len(closedTickets)}
	}

	allTickets = models.AllTickets{
		OpenTickets:     oT,
		PendingTickets:  pT,
		ResolvedTickets: rT,
		ClosedTickets:   cT,
	}

	/*status := ticket.Status.(float64)
	ticket.Status = models.Status(status).String()

	priority := ticket.Priority.(float64)
	ticket.Priority = models.Priority(priority).String()

	source := ticket.Source.(float64)
	ticket.Source = models.Source(source).String()
	*/

	respondWithJSON(w, http.StatusOK, allTickets)
	return
}

// getStatsForGroup godoc
// @Summary Get ticket statistics for a particular group
// @Description Get ticket statistics for a particular group id
// @Tags statistics
// @Accept  json
// @Produce  json
// @Param group_id path string true "Enter group id"
// @Param page query string true "Page number"
// @Success 200 {object} models.AllTickets
// @Router /api/v1/stats/group/{group_id} [get]
func (a *App) getStatsForGroup(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	var allTickets models.AllTickets

	var openTickets []*models.Ticket
	var pendingTickets []*models.Ticket
	var resolvedTickets []*models.Ticket
	var closedTickets []*models.Ticket

	vars := mux.Vars(r)
	groupId := vars["group_id"]
	page := r.FormValue("page")

	t := httpclient.GetTicketsPerGroup(a.Cfg.GetFreshDeskURL(), a.Cfg.GetApiKey(), groupId, page)
	for _, ticket := range t.Results {

		ticket.Status = ticket.Status.(float64)

		if ticket.Status == float64(2) {
			openTickets = append(openTickets, ticket)
		} else if ticket.Status == float64(3) {
			pendingTickets = append(pendingTickets, ticket)
		} else if ticket.Status == float64(4) {
			resolvedTickets = append(resolvedTickets, ticket)
		} else if ticket.Status == float64(5) {
			closedTickets = append(closedTickets, ticket)
		}
	}

	var oT models.OpenTickets
	var pT models.PendingTickets
	var rT models.ResolvedTickets
	var cT models.ClosedTickets

	if len(openTickets) > 0 {
		oT = models.OpenTickets{OpenTickets: openTickets, Total: len(openTickets)}
	}
	if len(pendingTickets) > 0 {

		pT = models.PendingTickets{PendingTickets: pendingTickets, Total: len(pendingTickets)}
	}
	if len(resolvedTickets) > 0 {
		rT = models.ResolvedTickets{ResolvedTickets: resolvedTickets, Total: len(resolvedTickets)}
	}
	if len(closedTickets) > 0 {
		cT = models.ClosedTickets{ClosedTickets: closedTickets, Total: len(closedTickets)}
	}

	allTickets = models.AllTickets{
		OpenTickets:     oT,
		PendingTickets:  pT,
		ResolvedTickets: rT,
		ClosedTickets:   cT,
	}

	/*status := ticket.Status.(float64)
	ticket.Status = models.Status(status).String()

	priority := ticket.Priority.(float64)
	ticket.Priority = models.Priority(priority).String()

	source := ticket.Source.(float64)
	ticket.Source = models.Source(source).String()
	*/

	respondWithJSON(w, http.StatusOK, allTickets)
	return
}
