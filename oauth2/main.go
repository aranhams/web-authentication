package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type githubResponse struct {
	Data struct {
		Viewer struct {
			ID string `json:"id"`
		}
	}
}

var githubConnections map[string]string

var githubOauthConfig = &oauth2.Config{
	ClientID:     "",
	ClientSecret: "",
	Endpoint:     github.Endpoint,
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/oauth/github", startGithubOauth)
	http.HandleFunc("/oauth2/receive", completeGithubOauth)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	html := `<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">
		<title>OAUHT2 Example</title>
	</head>
	<body>
		<form action="/oauth/github" method="POST">
			<input type="submit" value="Login with GitHub"/>
		</form>
	</body>
	</html>`
	fmt.Fprint(w, html)
}

func startGithubOauth(w http.ResponseWriter, r *http.Request) {
	redirectURL := githubOauthConfig.AuthCodeURL("0000")
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}

func completeGithubOauth(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	state := r.FormValue("state")

	if state != "0000" {
		http.Error(w, "State is incorrect", http.StatusBadRequest)
		return
	}

	token, err := githubOauthConfig.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, "Coudn't login", http.StatusInternalServerError)
		return
	}

	ts := githubOauthConfig.TokenSource(r.Context(), token)
	client := oauth2.NewClient(r.Context(), ts)

	requestBody := strings.NewReader(`{"query": "query {viewer {id}}"}`)
	resp, err := client.Post("https://api.github.com/graphql", "application/json", requestBody)
	if err != nil {
		http.Error(w, "Coudn't get user", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var gr githubResponse
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		http.Error(w, "Githyb invalid response", http.StatusInternalServerError)
	}

	githubID := gr.Data.Viewer.ID
	userID, ok := githubConnections[githubID]
	if !ok {
		// New user - create account
	}

	log.Println(userID)
	// Login to account userID using JWT
}
