package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

// isPrime checks if a number is prime.
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// nextPrime returns the next prime number after the given number.
func nextPrime(n int) int {
	n++
	for !isPrime(n) {
		n++
	}
	return n
}

// PrimeBackoff implements a backoff algorithm based on prime numbers.
type PrimeBackoff struct {
	currentPrime int
	startTime    time.Time
	maxElapsed   time.Duration
}

// NewPrimeBackoff creates a new PrimeBackoff instance.
func NewPrimeBackoff(maxElapsed time.Duration) *PrimeBackoff {
	return &PrimeBackoff{
		currentPrime: 1, // Start before the first prime (2)
		startTime:    time.Now(),
		maxElapsed:   maxElapsed,
	}
}

// NextBackoff returns the next backoff interval based on prime numbers.
// It returns false if the maximum elapsed time has been reached.
func (pb *PrimeBackoff) NextBackoff() (time.Duration, bool) {
	if time.Since(pb.startTime) >= pb.maxElapsed {
		return 0, false
	}

	pb.currentPrime = nextPrime(pb.currentPrime)
	backoffInterval := time.Duration(pb.currentPrime) * time.Second

	return backoffInterval, true
}

// Retry executes the provided function with prime-based backoff.
// It stops retrying if the function returns nil or the maximum elapsed time is reached.
func Retry(operation func() error, maxElapsed time.Duration) error {
	backoff := NewPrimeBackoff(maxElapsed)

	for {
		err := operation()
		if err == nil {
			return nil
		}

		nextInterval, shouldContinue := backoff.NextBackoff()
		if !shouldContinue {
			return fmt.Errorf("max retry time reached: %w", err)
		}

		fmt.Printf("Operation failed. Retrying in %v seconds...\n", nextInterval.Seconds())
		time.Sleep(nextInterval)
	}
}

// HTTP handler to simulate failures
func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintln(w, "Internal Server Error")
}

func main() {
	http.HandleFunc("/", handler)
	go http.ListenAndServe(":8080", nil)

	operation := func() error {
		resp, err := http.Get("http://localhost:8080")
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return errors.New("server returned non-200 status")
		}
		return nil
	}

	err := Retry(operation, 30*time.Second)
	if err != nil {
		fmt.Println("Operation failed:", err)
	} else {
		fmt.Println("Operation succeeded!")
	}
}
