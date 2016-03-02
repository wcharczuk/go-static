package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/blendlabs/go-util"
)

var portFlag string

func port() string {
	envPort := os.Getenv("PORT")
	if len(envPort) != 0 {
		return envPort
	}
	return "8080"
}

var pathFlag string

func path() string {
	if len(pathFlag) != 0 {
		return pathFlag
	}

	return pwd()
}

func pwd() string {
	pwd, pwdErr := os.Getwd()
	if pwdErr != nil {
		err := fmt.Errorf(pwdErr.Error())
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(1)
	}
	return pwd
}

func parseFlags() {
	flag.StringVar(&portFlag, "port", "8080", "Port to listen on.")
	flag.StringVar(&pathFlag, "path", "", "Path to serve.")
	flag.Parse()
}

func main() {
	parseFlags()

	servePort := port()
	servePath := path()

	fmt.Printf("Static File Server Serving: `%s` Listening on: %s\n", servePath, servePort)

	logger := log.New(os.Stdout, "", 0)

	fs := FileServer(logger, http.Dir(servePath))
	log.Fatal(http.ListenAndServe(":"+servePort, fs))
}

// FileServer returns a new (logged) FileHandler.
func FileServer(log *log.Logger, root http.Dir) http.Handler {
	return &FileHandler{Root: root, Handler: http.FileServer(root), Log: log}
}

// FileHandler is a logged version of http.fileHandler
type FileHandler struct {
	Root    http.Dir
	Handler http.Handler
	Log     *log.Logger
}

func (f *FileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now().UTC()
	f.Handler.ServeHTTP(w, r)
	end := time.Now().UTC()

	ip := util.GetIP(r)
	delta := end.Sub(start)
	contentLength := w.Header().Get("Content-Length")

	if len(contentLength) == 0 {
		contentLength = "cached"
	} else {
		contentLength = fmt.Sprintf("%s bytes", contentLength)
	}

	f.Log.Printf("%s - %s - %s - %s - %v - %s", start.Format(time.RFC3339), ip, r.Method, r.URL.Path, delta, contentLength)
}
