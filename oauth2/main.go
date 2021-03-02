package main

import (
	"fmt"
	"net/http"
)

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
		<title>HMAC Example</title>
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

}
