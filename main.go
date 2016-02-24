package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func port() string {
	envPort := os.Getenv("PORT")
	if len(envPort) != 0 {
		return envPort
	}
	return "8080"
}

func main() {
	pwd, pwdErr := os.Getwd()
	if pwdErr != nil {
		fmt.Errorf(pwdErr.Error())
		os.Exit(1)
	}
	fmt.Printf("Static File Server Listening on: %s\n", port())
	fs := http.FileServer(http.Dir(pwd))
	log.Fatal(http.ListenAndServe(":"+port(), fs))
}
