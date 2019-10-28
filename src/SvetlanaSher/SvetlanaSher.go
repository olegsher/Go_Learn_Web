package SvetlanaSher

import (
	"net/http"
)

func SvetlanaSherIndexHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello www.svetlanasher.com!!!!") // send data to client side
	http.Handle("/", http.FileServer(http.Dir("./static")))
}
