package main

import (
	"net/http"

	"github.com/crnvl96/golang-webserver/helpers"
)

func main() {
	http.HandleFunc("/cotacao", getDailyQuotation)
	http.ListenAndServe(":8080", nil)
}

func getDailyQuotation(w http.ResponseWriter, r *http.Request) {
	fetchCtx, fetchCancel := helpers.GenerateContextWithTimeout(200)
	defer fetchCancel()

	data, err := helpers.FetchAPI(fetchCtx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	db, err := helpers.BootstrapDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	defer db.Close()

	err = helpers.CreateDBTables(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	dbCtx, dbCancel := helpers.GenerateContextWithTimeout(10)
	defer dbCancel()

	err = helpers.InsertIntoDB(dbCtx, db, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write([]byte(data))
}
