package app

import (
	"github.com/kosha/freshdesk-connector/pkg/httpclient"
	"net/http"
)

// getAgents godoc
// @Summary Get freshdesk agents
// @Description List all agents
// @Tags agents
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Agent
// @Failure      400  {object} string "bad request"
// @Failure      403  {object}  string "permission denied"
// @Failure      404  {object}  string "not found"
// @Failure      500  {object}  string "internal server error"
// @Router /api/v1/agents [get]
func (a *App) getAgents(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	agents := httpclient.GetAgents(a.Cfg.GetFreshDeskURL(), a.Cfg.GetApiKey())

	respondWithJSON(w, http.StatusOK, agents)
}
