//go:build integration

package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"testing"
)

// TestBackoffWithVegeta simulates traffic using Vegeta and observes the backoff behavior.
func TestBackoffWithVegeta(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Internal Server Error")
	}))
	defer server.Close()

	targets := fmt.Sprintf("GET %s\n", server.URL)
	err := os.WriteFile("vegeta-targets.txt", []byte(targets), 0644)
	if err != nil {
		t.Fatalf("Failed to write Vegeta targets file: %v", err)
	}

	resultsFile := "vegeta-results.bin"
	cmd := exec.Command("vegeta", "attack", "-duration=10s", "-rate=10", "-targets=vegeta-targets.txt", "-output", resultsFile)
	err = cmd.Run()
	if err != nil {
		t.Fatalf("Vegeta attack failed: %v", err)
	}

	cmd = exec.Command("vegeta", "report", resultsFile)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Vegeta report failed: %v", err)
	}

	t.Logf("Vegeta Report:\n%s", string(output))
}
