package app

import (
	"github.com/kosha/freshdesk-connector/pkg/httpclient"
	"net/http"
)

// getGroups godoc
// @Summary Get groups detail
// @Description List all groups
// @Tags groups
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Group
// @Router /api/v1/groups [get]
func (a *App) getGroups(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	agents := httpclient.GetGroups(a.Cfg.GetFreshDeskURL(), a.Cfg.GetApiKey())

	respondWithJSON(w, http.StatusOK, agents)
}
