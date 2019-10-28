package Tools

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func CheckMyExtIP() (status int) {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		os.Exit(1)
	}
	defer resp.Body.Close()
	fmt.Printf("Payload is %d", &(resp.Body))
	fmt.Printf("\nGetting external IP,  Status =  %s, ext. IP is ", resp.Status)
	io.Copy(os.Stdout, resp.Body)

	//TODO convert string from "http://myexternalip.com/raw" to IP using regular expression and return it. Copy resp.Body to buffer
	return resp.StatusCode
}
