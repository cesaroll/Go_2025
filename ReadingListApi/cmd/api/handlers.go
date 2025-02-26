package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}

	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}

func (app *application) getCreateBooksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "get all books")
	}
	if r.Method == http.MethodPost {
		fmt.Fprintln(w, "create a new book")
	}
}

func (app *application) getUpdateDeleteBooksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			fmt.Fprintln(w, "get a book")
		case http.MethodPut:
			fmt.Fprintln(w, "update a book")
		case http.MethodDelete:
			fmt.Fprintln(w, "delete a book")
	}
}
