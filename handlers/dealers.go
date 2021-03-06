package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	m "git.mailbox.com/mailbox/models"
	u "git.mailbox.com/mailbox/utils"
)

func dealersHandler(db m.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dealers, err := db.GetDealers()
		if err != nil {
			log.Printf("Error fetching dealers from DB: %s", err)

			errResponse := m.Error{
				Code:    u.I32Ptr(http.StatusInternalServerError),
				Message: u.SPtr("Internal server error"),
			}
			marshalledError, err := json.Marshal(errResponse)
			if err != nil {
				log.Printf("failed to marshal the error response: %s", err)
				http.Error(w, "something went wrong", http.StatusInternalServerError)
			}
			http.Error(w, string(marshalledError), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		marshalledRes, err := json.Marshal(dealers)
		if err != nil {
			log.Printf("failed to marshal the success response: %s", err)
			http.Error(w, "something went wrong", http.StatusInternalServerError)
		}

		fmt.Fprintf(w, string(marshalledRes))
	}
}
