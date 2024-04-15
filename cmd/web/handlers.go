package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/morhaham/snippetbox/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
	}

	data := &templateData{Snippets: snippets}
	app.render(w, r, "home.page.tmpl", data)
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrorNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	data := &templateData{
		Snippet: snippet,
	}
	app.render(w, r, "show.page.tmpl", data)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	expires := r.PostForm.Get("expires")

	validationErrors := make(map[string]string)
	if strings.TrimSpace(title) == "" {
		validationErrors["title"] = "Title is required"
	} else if len(title) > 100 {
		validationErrors["title"] = "Title is too long"
	}

	if strings.TrimSpace(content) == "" {
		validationErrors["content"] = "Content is required"
	}

	if strings.TrimSpace(expires) == "" {
		validationErrors["expires"] = "Expiry time is required"
	} else if expires != "365" && expires != "7" && expires != "1" {
		validationErrors["expires"] = "Invalid expiry time"
	}

	if len(validationErrors) > 0 {
		app.render(w, r, "create.page.tmpl", &templateData{
			FormErrors: validationErrors,
			FormData:   r.PostForm,
		})
		return
	}

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}

func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", nil)
}
