package myhttp
import (
	"fmt"
	"log"
	"net/http"
)

func response(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func Myhttp() {
	http.HandleFunc("/", response)
	err := http.ListenAndServe("9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}