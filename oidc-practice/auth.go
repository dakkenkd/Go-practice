package main

import (
	"log"
	"net/http"

	oidc "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

func main() {
	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {

		provider, err := oidc.NewProvider(r.Context(), issuer)
		if err != nil {
			log.Fatal(err)

		}
		config := oauth2.Config{
			ClientID:     "Bk6HNFLoneeoOEcKClEr510oG3Gp7Sn9",
			ClientSecret: "W0E-8-uNZX8qB4toIy8mnhY1NP8xchc46d1XB0aNXLvyhtrXQi72w04CfXMHTvLk",
			Endpoint:     provider.Endpoint(),
			RedirectURL:  "http://localhost:3000/callback",
			Scopes:       []string{oidc.ScopeOpenID},
		}

		state := "fjaposijfaoe"

		authURL := config.AuthCodeURL(state)
		http.Redirect(w, r, authURL, http.StatusFound)
	})

	// サーバ起動
	http.ListenAndServe(":3000", nil)
}
