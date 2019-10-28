package SvetlanaSher

import (
	"fmt"
	"net/http"
)

func SvetlanaSherIndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello www.svetlanasher.com!!!!") // send data to client side
}
