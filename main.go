package main

import "fmt"
import "net/http"

func main() {
handlerIndex := func(w http.ResponseWriter, r *http.Request) {
fmt. Fprintln(w, "Hallo nama saya m hafidz Firmansyah")
}
http.HandleFunc("/", handlerIndex)

fmt.Println("server started at 43.173.13.18:9014")
http.ListenAndServe(":9014", nil)

}
