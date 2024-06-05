package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/jdkato/prose/v2"
)

type Request struct {
	Text      string `json:"text"`
	MaxTokens int    `json:"max_tokens"`
}

type Response struct {
	Chunks       []string `json:"chunks"`
	ResponseTime float64  `json:"response_time"`
}

func SplitTextIntoChunks(text string, maxTokens int) ([]string, error) {
	doc, err := prose.NewDocument(text)
	if err != nil {
		return nil, err
	}

	var chunks []string
	var currentChunk []string
	currentLength := 0

	for _, sent := range doc.Sentences() {
		sentenceTokens := strings.Fields(sent.Text)
		sentenceLength := len(sentenceTokens)

		if currentLength+sentenceLength > maxTokens {
			chunks = append(chunks, strings.Join(currentChunk, " "))
			currentChunk = sentenceTokens
			currentLength = sentenceLength
		} else {
			currentChunk = append(currentChunk, sentenceTokens...)
			currentLength += sentenceLength
		}
	}

	if len(currentChunk) > 0 {
		chunks = append(chunks, strings.Join(currentChunk, " "))
	}

	return chunks, nil
}

func chunkHandler(w http.ResponseWriter, r *http.Request) {
	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if req.MaxTokens <= 0 {
		http.Error(w, "MaxTokens must be greater than zero", http.StatusBadRequest)
		return
	}

	start := time.Now()
	chunks, err := SplitTextIntoChunks(req.Text, req.MaxTokens)
	if err != nil {
		http.Error(w, "Error processing text", http.StatusInternalServerError)
		return
	}

	resp := Response{
		Chunks:       chunks,
		ResponseTime: float64(time.Since(start)) / float64(time.Second),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s %v", r.Method, r.RequestURI, r.RemoteAddr, time.Since(start))
	})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/chunk", chunkHandler).Methods("POST")

	// Add logging middleware
	r.Use(loggingMiddleware)

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	addr := fmt.Sprintf(":%s", port)
	fmt.Printf("Server running on port %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
