package main

import (
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("<html><body>Hello world</body></html>"))
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}
