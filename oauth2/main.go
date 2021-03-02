package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func getEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error to loading .evn file")
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/oauth/github", startGithubOauth)
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
	getEnv()
	var githubOauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("CLIENTID"),
		ClientSecret: os.Getenv("CLIENTSECRET"),
		Endpoint:     github.Endpoint,
	}

	redirectURL := githubOauthConfig.AuthCodeURL("0000")
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}
