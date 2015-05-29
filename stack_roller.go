package main

import (
  "net/http"
  "os"
  "github.com/gorilla/mux"
  "fmt"
  "time"
  "math/rand"
)



func main() {
  r := mux.NewRouter()
  r.HandleFunc("/", home_handler)
  r.HandleFunc("/{p}", home_handler)

  http.Handle("/", r)

  err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)

  if err != nil {
    panic(err)
  }
}


func home_handler(res http.ResponseWriter, req *http.Request) {
  rand.Seed(time.Now().UnixNano())
  
  databases := [...]string{"mysql", "mongodb", "redis", "memcached", "sqlite", "dynamodb", "text file"}
  languages := [...]string{"ruby", "go", "php", "node.js", "python", "java"}

  database := databases[rand.Intn(len(databases))]
  language := languages[rand.Intn(len(languages))]

  fmt.Fprintf(res,
`<html>
  <head>
    <title>Stack Roller</title>
  </head>
  <body style="text-align:center;">
    <h1>Stack Roller</h1>
    <h2>Having trouble designing your stack? Why not random?</h2>
    <div>
      <p style="font-size: 16px;">You should use <strong>%s</strong>, and store it in <strong>%s</strong>.</p>
    </div>
  </body>
</html>`, language, database)
}