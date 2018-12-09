package shipnfixmytask

import (
	"../../app"
	templates ".."
	"net/http"
)

func MytaskHandler(w http.ResponseWriter, r *http.Request) {

	session, err := app.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templates.RenderTemplate(w, "shipnfixmytask", session.Values["profile"])
}
