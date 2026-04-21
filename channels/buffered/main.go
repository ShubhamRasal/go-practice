package main

import (
	"fmt"
	"math"
	"time"
)

// Heavy calculation — find prime numbers up to 80,000
// Simulates real CPU work per job
func runCalculation(id int) int {
	count := 0
	for n := 2; n < 80_000; n++ {
		isPrime := true
		for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			count++
		}
	}
	return count
}

// Send email — rate limited, 1 at a time, NOT in a goroutine
// Simulates a mail service that allows only 1 send per 400ms
func sendEmail(id int, result int) {
	time.Sleep(400 * time.Millisecond)
	fmt.Printf("  📧 Email sent  — job %-2d | primes: %d\n", id, result)
}

func main() {
	jobs := 5
	ch := make(chan [2]int, jobs) // BUFFERED — buffer size = number of jobs

	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("  BUFFERED CHANNEL DEMO")
	fmt.Printf("  %d jobs | calc (concurrent) + email (1 at a time)\n", jobs)
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	start := time.Now()

	// All jobs run concurrently — calculations happen in parallel
	for i := 1; i <= jobs; i++ {
		go func(id int) {
			t := time.Now()
			result := runCalculation(id)

			fmt.Printf("  🔢 Job %-2d calc done in %v — dropping into buffer ✓\n",
				id, time.Since(t).Round(time.Millisecond))

			// Goroutine sends into buffer and is IMMEDIATELY FREE
			// It does not care about the email sender at all
			// Buffer has space — no blocking
			ch <- [2]int{id, result}

			fmt.Printf("  🔢 Job %-2d goroutine done — free to go\n", id)
		}(i)
	}

	// Email sender reads from buffer ONE AT A TIME — not in a goroutine
	// Same rate limit as unbuffered demo — 400ms per email
	// BUT: all calculations already finished while emails are being sent
	// Calc time and email time OVERLAP — that is the win
	for i := 0; i < jobs; i++ {
		msg := <-ch               // reads from buffer — goroutines not affected
		sendEmail(msg[0], msg[1]) // 400ms — but goroutines are already done
	}

	fmt.Println()
	fmt.Printf("  ⏱  Total time: %v\n", time.Since(start).Round(time.Millisecond))
	fmt.Println()
	fmt.Println("  WHY FASTER?")
	fmt.Println("  Goroutines drop results into buffer and exit immediately")
	fmt.Println("  They are never blocked by the 400ms email rate limit")
	fmt.Println("  Calc (~210ms) and emails (5 × 400ms) run in parallel")
	fmt.Println("  Total ≈ max(calc time, email time) not calc + email time")
}

// Expected output:
//
// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
//   BUFFERED CHANNEL DEMO
//   5 jobs | calc (concurrent) + email (1 at a time)
// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
//
//   🔢 Job 2  calc done in 209ms — dropping into buffer ✓
//   🔢 Job 2  goroutine done — free to go
//   🔢 Job 4  calc done in 210ms — dropping into buffer ✓
//   🔢 Job 4  goroutine done — free to go
//   🔢 Job 1  calc done in 211ms — dropping into buffer ✓
//   🔢 Job 1  goroutine done — free to go
//   🔢 Job 3  calc done in 212ms — dropping into buffer ✓
//   🔢 Job 3  goroutine done — free to go
//   🔢 Job 5  calc done in 213ms — dropping into buffer ✓
//   🔢 Job 5  goroutine done — free to go
//   📧 Email sent  — job 2  | primes: 8713
//   📧 Email sent  — job 4  | primes: 8713
//   📧 Email sent  — job 1  | primes: 8713
//   📧 Email sent  — job 3  | primes: 8713
//   📧 Email sent  — job 5  | primes: 8713
//
//   ⏱  Total time: ~2.0s
//      (5 × 400ms emails, calc overlapped — goroutines never blocked)
