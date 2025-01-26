package main

import (
	"errors"
	"testing"
	"time"
)

func TestIsPrime(t *testing.T) {
	tests := []struct {
		input  int
		output bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{13, true},
		{15, false},
	}

	for _, test := range tests {
		result := isPrime(test.input)
		if result != test.output {
			t.Errorf("isPrime(%d) = %v; expected %v", test.input, result, test.output)
		}
	}
}

func TestNextPrime(t *testing.T) {
	tests := []struct {
		input  int
		output int
	}{
		{1, 2},
		{2, 3},
		{3, 5},
		{5, 7},
		{13, 17},
	}

	for _, test := range tests {
		result := nextPrime(test.input)
		if result != test.output {
			t.Errorf("nextPrime(%d) = %d; expected %d", test.input, result, test.output)
		}
	}
}

func TestPrimeBackoff(t *testing.T) {
	backoff := NewPrimeBackoff(10 * time.Second)

	// Test the first few prime intervals
	expectedPrimes := []int{2, 3, 5, 7, 11}
	for _, expected := range expectedPrimes {
		interval, shouldContinue := backoff.NextBackoff()
		if !shouldContinue {
			t.Error("NextBackoff returned false before max elapsed time")
		}
		if interval != time.Duration(expected)*time.Second {
			t.Errorf("NextBackoff returned interval %v; expected %v seconds", interval, expected)
		}
	}
}

func TestRetry(t *testing.T) {
	attempts := 0
	operation := func() error {
		attempts++
		if attempts < 3 {
			return errors.New("temporary failure")
		}
		return nil
	}

	err := Retry(operation, 10*time.Second)
	if err != nil {
		t.Errorf("Retry failed: %v", err)
	}
	if attempts != 3 {
		t.Errorf("Expected 3 attempts; got %d", attempts)
	}
}
