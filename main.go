package main

import (
	"fmt"
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
		fmt.Errorf(pwdErr)
		os.Exit(1)
	}
	panic(http.ListenAndServe(":"+port(), http.FileServer(http.Dir(pwd))))
}
