package app

import (
	"github.com/kosha/freshdesk-connector/pkg/httpclient"
	"net/http"
)

// getAccounts godoc
// @Summary Get account detail
// @Description List account metadata
// @Tags account
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Account
// @Router /api/v1/account [get]
func (a *App) getAccounts(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	accounts := httpclient.GetAccounts(a.Cfg.GetFreshDeskURL(), a.Cfg.GetApiKey())

	respondWithJSON(w, http.StatusOK, accounts)
}
