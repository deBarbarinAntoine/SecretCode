package SecretCode

import "net/http"

//	routes
//
// initialises all the routes.
func routes() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/index/treatment", treatmentHandler)
}
