package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    // Define the port to listen on
    port := "8085"

    // Open the CSV file
    file, err := os.Open("file.csv")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Parse the CSV file
    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        log.Fatal(err)
    }

    // Create an HTTP handler to serve the CSV data as an HTML table
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "<html><body><table>")
        for _, row := range records {
            fmt.Fprintln(w, "<tr>")
            for _, col := range row {
                fmt.Fprintf(w, "<td>%s</td>", col)
            }
            fmt.Fprintln(w, "</tr>")
        }
        fmt.Fprintln(w, "</table></body></html>")
    })

    // Start the server
    log.Printf("Listening on :%s...", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatal(err)
    }
}
