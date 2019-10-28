package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"Tools"
	"github.com/gorilla/mux"
	"theAuto"
)

const url_Albar string = "https://rent.albar.co.il/en/rent-in-israel/"
const url_Hertz string = "https://www.hertz.co.il/en/HomePage/.aspx"
const url_local string = "http://127.0.0.1:80"
const url_theAuto string = "http://www.theauto.info"
const url_SvetlanaSher string = "http://www.svetlanasher.com"

func SvetlanaSherIndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello www.svetlanasher.com!") // send data to client side
}

func VendorGetPrices() {

}

func VendorReserveCar() {

}

func CustomerReservationRequest() {}

func CustomerGetPersonalInfo() {}

func StartClient() {
	fmt.Println("Starting Client...")
	//TODO Open client to external web server
	var err error

	/*
			exec.Command("/usr/bin/google-chrome-stable", url_local).Start()
			if err != nil {
				log.Fatal(err)
			}

		exec.Command("/usr/bin/google-chrome-stable", url_local).Start()
		if err != nil {
			log.Fatal(err)
		}
	*/
	resp, err := http.Get(url_local)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		os.Exit(1)
	}

	resp1, err := http.Get(url_theAuto)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		os.Exit(1)
	}

	/*resp, err = http.PostForm("url_Albar",
	url.Values{"key": {"Value"}, "id": {"123"}})
	*/
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)

	defer resp1.Body.Close()
	io.Copy(os.Stdout, resp1.Body)

	/*
	 	exec.Command("/usr/bin/google-chrome-stable", url_local).Start()
	 	if err != nil {
	 		log.Fatal(err)
	 	}
	 	resp, err = http.Get(url_local)
	 	if err != nil {
	 		os.Stderr.WriteString(err.Error())
	 		os.Stderr.WriteString("\n")
	 		os.Exit(1)
	 	}
	 	defer resp.Body.Close()
	 	io.Copy(os.Stdout, resp.Body)
	 }

	*/
}

func main() {
	fmt.Printf(" Proccess %d\n", Tools.CheckMyExtIP())
	r := mux.NewRouter()
	tA := r.Host("www.theauto.info").Subrouter()
	svetlanaSher := r.Host("www.svetlanasher.com").Subrouter()

	tA.HandleFunc("/", theAuto.TheAutoIndexHandler)
	tA.HandleFunc("/cars/", theAuto.TheAutoIndexHandlerCars)

	svetlanaSher.HandleFunc("/", SvetlanaSherIndexHandler)

	// Bind to a port and pass our router in
	go log.Fatal(http.ListenAndServe(":80", r))
	// StartClient()
}
