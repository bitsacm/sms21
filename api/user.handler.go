package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dush-t/sms21/db/models"
	"github.com/dush-t/sms21/util"
)

func GoogleAuthScreen(data models.Models) http.Handler {
	redirect_uri := "http://localhost:8080/auth/google/redirect"
	// client_secret := "uj-pqT0g_HNkFkJjy_3_skza"
	client_id := "214753833156-tufte7ehqjvesud51t5ta3uo59bo6ol4.apps.googleusercontent.com"
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		newUrl := "https://accounts.google.com/o/oauth2/v2/auth/oauthchooseaccount?access_type=offline&scope=https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.email%20https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.profile&response_type=code&prompt=consent&client_id=" + client_id + "&redirect_uri=" + redirect_uri + "&flowName=GeneralOAuthFlow&state=%2F3%2Fhome"
		http.Redirect(w, r, newUrl, http.StatusSeeOther)
	})
}

func GoogleTokenExchange(data models.Models) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		access_token := util.GetTokenFromCode(r.URL.Query()["code"][0])

		fmt.Println("token is:", access_token)

		u := models.GetUserData(access_token)

		us, err := data.Users.GetUserByUsername(u.Username)

		if err != nil {
			err = data.Users.Add(u)
		}

		fmt.Println("User is: %+v", us)

		tokenString, err := us.GenerateJWT()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Error generating token for user:", err)
			return
		}

		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "jwt", Value: tokenString, Expires: expiration}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusOK)
	})
}
