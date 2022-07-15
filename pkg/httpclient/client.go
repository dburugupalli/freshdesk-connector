package httpclient

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/kosha/freshdesk-connector/pkg/models"
	"io/ioutil"
	"net/http"
	neturl "net/url"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func makeHttpReq(apiKey string, req *http.Request) []byte {
	req.Header.Add("Authorization", "Basic "+basicAuth(apiKey, "X"))
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	return bodyBytes
}

func GetAllTickets(url string, apiKey, pageNum string) []*models.Ticket {
	req, err := http.NewRequest("GET", url+"/api/v2/tickets?page="+pageNum, nil)
	if err != nil {
		return nil
	}
	var tickets []*models.Ticket

	res := makeHttpReq(apiKey, req)

	// Convert response body to target struct
	err = json.Unmarshal(res, &tickets)
	if err != nil {
		return nil
	}
	return tickets
}

func GetAgents(url string, apiKey string) []*models.Agent {
	req, err := http.NewRequest("GET", url+"/api/v2/agents", nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var agents []*models.Agent

	res := makeHttpReq(apiKey, req)

	// Convert response body to target struct
	err = json.Unmarshal(res, &agents)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return agents
}

func GetAccounts(url string, apiKey string) *models.Account {
	req, err := http.NewRequest("GET", url+"/api/v2/account", nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var accounts *models.Account

	res := makeHttpReq(apiKey, req)

	// Convert response body to target struct
	err = json.Unmarshal(res, &accounts)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return accounts
}

func GetGroups(url string, apiKey string) []*models.Group {
	req, err := http.NewRequest("GET", url+"/api/v2/groups", nil)
	if err != nil {
		return nil
	}
	var groups []*models.Group

	res := makeHttpReq(apiKey, req)

	// Convert response body to target struct
	err = json.Unmarshal(res, &groups)
	if err != nil {
		return nil
	}
	return groups
}

func SearchTickets(url string, apiKey, query string, page string) *models.SearchResults {

	if page == "" {
		page = "1"
	}
	baseUrl := url + "/api/v2/search/tickets?page=" + page + "&query=" + neturl.QueryEscape(query)

	/*if agentId != "" && status == 0 {
		baseUrl = baseUrl + "\"agent_id:" + agentId + "\""
	}
	if status != 0 && agentId == "" {
		baseUrl = baseUrl + "\"status:" + strconv.Itoa(status) + "\""
	}
	if status != 0 && agentId != "" {
		baseUrl = baseUrl + "\"agent_id:" + agentId + "%20OR%20status:" + strconv.Itoa(status) + "\""
	}
		if query != "" {
			baseUrl += query + "\""
		}
	*/
	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		return nil
	}
	var searchResult *models.SearchResults

	res := makeHttpReq(apiKey, req)

	// Convert response body to target struct
	err = json.Unmarshal(res, &searchResult)
	if err != nil {
		return nil
	}
	return searchResult
}

func GetTicketsPerAgent(url string, apiKey, agentId string, page string) *models.SearchResults {

	if page == "" {
		page = "1"
	}
	baseUrl := url + "/api/v2/search/tickets?page=" + page + "&query="

	if agentId != "" {
		baseUrl = baseUrl + "\"agent_id:" + agentId + "\""
	}
	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		return nil
	}
	var searchResult *models.SearchResults

	res := makeHttpReq(apiKey, req)

	// Convert response body to target struct
	err = json.Unmarshal(res, &searchResult)
	if err != nil {
		return nil
	}
	return searchResult
}

func GetTicketsCreatedAt(url string, apiKey, createdAt string, page string) *models.SearchResults {

	if page == "" {
		page = "1"
	}
	baseUrl := url + "/api/v2/search/tickets?page=" + page + "&query="

	if createdAt != "" {
		baseUrl = baseUrl + "\"created_at:>'" + createdAt + "'\""
	}
	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		return nil
	}
	var searchResult *models.SearchResults

	res := makeHttpReq(apiKey, req)

	// Convert response body to target struct
	err = json.Unmarshal(res, &searchResult)
	if err != nil {
		return nil
	}
	return searchResult
}

func GetTicketsPerGroup(url string, apiKey, groupId string, page string) *models.SearchResults {

	if page == "" {
		page = "1"
	}
	baseUrl := url + "/api/v2/search/tickets?page=" + page + "&query="

	if groupId != "" {
		baseUrl = baseUrl + "\"group_id:" + groupId + "\""
	}
	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		return nil
	}
	var searchResult *models.SearchResults

	res := makeHttpReq(apiKey, req)

	// Convert response body to target struct
	err = json.Unmarshal(res, &searchResult)
	if err != nil {
		return nil
	}
	return searchResult
}

func GetSingleTicket(url, id string, apiKey string) *models.Ticket {
	req, err := http.NewRequest("GET", url+"/api/v2/tickets/"+id+"+?include=requester,company,stats", nil)
	if err != nil {
		return nil
	}
	var ticket *models.Ticket

	res := makeHttpReq(apiKey, req)
	// Convert response body to target struct
	err = json.Unmarshal(res, &ticket)
	if err != nil {
		return nil
	}
	return ticket
}

func CreateTicket(url, apiKey string, ticket *models.Ticket) (string, error) {

	jsonReq, err := json.Marshal(ticket)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", url+"/api/v2/tickets", bytes.NewBuffer(jsonReq))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	return string(makeHttpReq(apiKey, req)), nil
}

func RemoveTicket(url, id, apiKey string) (string, error) {
	req, err := http.NewRequest(http.MethodDelete, url+"/api/v2/tickets/"+id, nil)
	if err != nil {
		return "", err
	}
	bodyString := makeHttpReq(apiKey, req)
	return string(bodyString), nil
}
