/* First thing to do will be to declare package */
package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

/* If we want to export this function its name needs to be capitalized */
func HandleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HELLO GORILLA\n"))
}

/* This will run when the package is imported */
func init() {
	r := mux.Newrouter()
	r.HandleFun("/", HandleHome)

	//Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
