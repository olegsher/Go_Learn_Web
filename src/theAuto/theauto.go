package theAuto

import (
	"fmt"
	"net/http"
)

func TheAutoIndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello www.theauto.info!") // send data to client side
}

func TheAutoIndexHandlerCars(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Cars!") // send data to client side
}
