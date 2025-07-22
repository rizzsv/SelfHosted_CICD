package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"selfhosted-ci/config"
	"selfhosted-ci/executor"
	"selfhosted-ci/routes"

	"github.com/gin-gonic/gin"
)

type GithubWebhook struct {
	Repository struct {
		CloneURL string `json:"clone_url"`
		Name	 string `json:"name"`
	} `json:"repository"`
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	var payload GithubWebhook
	body, _ := io.ReadAll(r.Body)
	if err := json.Unmarshal(body, &payload); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		log.Printf("Error parsing webhook payload: %v", err)
		return
	}
	go executor.RunJob(payload.Repository.CloneURL)
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(w, "Job started for %s", payload.Repository.Name)
}

func main() {
	http.HandleFunc("/webhook", webhookHandler)
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	r := gin.Default()
	config.SetupCORS(r)
	routes.SetupRoutes(r)
	
	r.Run(":8080")
}