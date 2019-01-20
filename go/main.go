package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"
)

// =====================================================================

func get_factors(n int) []int {
    factors := []int{}
    for p := 2; n > 1 ; p++ {
        // If n % p == 0, p must be prime. We have already checked
        // against everything smaller.
        for n % p == 0 {
            n = n/p
            factors = append(factors, p)
        }
    }
    return factors
}

// =====================================================================

func factors(w http.ResponseWriter, r *http.Request) {
    // Report hits in the terminal.
    fmt.Printf("Hit URL: %s%s\n", r.Host, r.URL.Path)
    // Grab the given number. Third-party tools do this better.
    n, err := strconv.Atoi( r.URL.Path[9:] )
    if err == nil && n > 0 {
        out, _ := json.MarshalIndent(get_factors(n), "", "    ")
        fmt.Fprintf(w, string(out) + "\n")
    } else {
        w.WriteHeader(400)
        fmt.Fprintln(w, "{\"error\": \"Bad input\"}")
    }
}

// =====================================================================

func catchall(w http.ResponseWriter, r *http.Request) {
    // Report hits in the terminal.
    fmt.Printf("Hit URL: %s%s\n", r.Host, r.URL.Path)
    w.WriteHeader(404)
    fmt.Fprintln(w, "{\"error\": \"You're lost\"}")
}

// =====================================================================

func main() {
    handler := http.NewServeMux()
    handler.HandleFunc("/factors/", factors)
    handler.HandleFunc("/", catchall)
    err := http.ListenAndServe(":8081", handler)
    log.Fatal(err)
}
