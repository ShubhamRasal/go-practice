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
	ch := make(chan [2]int) // UNBUFFERED — no second argument

	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("  UNBUFFERED CHANNEL DEMO")
	fmt.Printf("  %d jobs | calc (concurrent) + email (1 at a time)\n", jobs)
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println()

	start := time.Now()

	// All jobs run concurrently — calculations happen in parallel
	for i := 1; i <= jobs; i++ {
		go func(id int) {
			t := time.Now()
			result := runCalculation(id)

			fmt.Printf("  🔢 Job %-2d calc done in %v — now WAITING to send...\n",
				id, time.Since(t).Round(time.Millisecond))

			// Goroutine BLOCKS here until the email sender (main) reads from ch
			// Even though calculation is done, goroutine is stuck
			// because email sender is busy with a previous email (400ms)
			ch <- [2]int{id, result}

			fmt.Printf("  🔢 Job %-2d unblocked — goroutine now free\n", id)
		}(i)
	}

	// Email sender reads from channel ONE AT A TIME — not in a goroutine
	// Each read unblocks exactly ONE goroutine
	// But then immediately takes 400ms for the email
	// Next goroutine stays blocked during those 400ms
	for i := 0; i < jobs; i++ {
		msg := <-ch               // unblocks one goroutine
		sendEmail(msg[0], msg[1]) // 400ms — next goroutine stuck until this finishes
	}

	fmt.Println()
	fmt.Printf("  ⏱  Total time: %v\n", time.Since(start).Round(time.Millisecond))
	fmt.Println()
	fmt.Println("  WHY SLOW?")
	fmt.Println("  Goroutines finish calc but pile up at ch <- result")
	fmt.Println("  Email sender reads one, takes 400ms, then reads next")
	fmt.Println("  Goroutines are held hostage by the email rate limit")
}

// Expected output:
//
// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
//   UNBUFFERED CHANNEL DEMO
//   5 jobs | calc (concurrent) + email (1 at a time)
// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
//
//   🔢 Job 3  calc done in 210ms — now WAITING to send...
//   🔢 Job 1  calc done in 212ms — now WAITING to send...
//   🔢 Job 5  calc done in 213ms — now WAITING to send...
//   🔢 Job 2  calc done in 214ms — now WAITING to send...
//   🔢 Job 4  calc done in 215ms — now WAITING to send...
//   📧 Email sent  — job 3  | primes: 8713         ← 400ms later
//   🔢 Job 3  unblocked — goroutine now free
//   📧 Email sent  — job 1  | primes: 8713         ← another 400ms
//   🔢 Job 1  unblocked — goroutine now free
//   ...
//
//   ⏱  Total time: ~2.2s
//      (210ms calc + 5 × 400ms emails = 2.21s)
