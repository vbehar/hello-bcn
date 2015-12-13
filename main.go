package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        log.Println("Processing a request")
        io.WriteString(w, fmt.Sprintf(`
<html>
<head>
<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap.min.css" integrity="sha384-1q8mTJOASx8j1Au+a5WDVnPi2lkFfwwEAa8hDDdjZlpLegxhjVME1fgjWPGmkzs7" crossorigin="anonymous">
</head>
<body>
<div style="text-align: center; padding: 40px 15px;">
<h1>Hello Barcelona !</h1>
<h2>I'm alive !</h2>
</div>
</body>
</html>
`))
    })
    log.Println("Starting on port 8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("Failed to start", err)
    }
}
