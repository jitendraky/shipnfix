package shipnfixtask

import (
	"../../app"
	templates ".."
	"net/http"
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {

	session, err := app.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templates.RenderTemplate(w, "shipnfixtask", session.Values["profile"])
}
