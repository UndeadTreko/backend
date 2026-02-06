package main

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/jwt"
)

func main() {
	//sevice account config
	var googleServiceAccountConfig = &jwt.Config{
		Email:      "",
		privateKey: []byte("YOUR_SERVICE_ACCOUNT_PRIVATE_KEY"),
		Scopes: []string{"https://www.googleapis.com/auth/cloud-platform",
			"https://www.googleapis.com/auth/drive"},
	}
	client := googleServiceAccountConfig.Client(oauth2.NoContext)
	client.Get("https://www.googleapis.com/auth/drive")
}
