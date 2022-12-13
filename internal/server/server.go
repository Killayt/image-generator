package server

import (
	"net/http"

	"github.com/Killayt/image-generator/configs"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func rend(w http.ResponseWriter, msg string) {
	_, err := w.Write([]byte(msg))
	if err != nil {
		log.Println(err)
	}
}

func imgHandler(w http.ResponseWriter, r *http.Request) {
	rend(w, "img\n")
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	rend(w, "PONG\n")
}

func robotsHandler(w http.ResponseWriter, r *http.Request) {
	rend(w, "robots\n")
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	rend(w, "favicon\n")
}

func Run(conf configs.ConfI) {
	r := mux.NewRouter()

	r.HandleFunc("/", imgHandler)
	r.HandleFunc("/ping", pingHandler)
	r.HandleFunc("/robots", robotsHandler)
	r.HandleFunc("/favicon", faviconHandler)
	log.Println("Server is starting ...")
	if err := http.ListenAndServe(":"+conf.GetPort(), r); err != nil {
		log.Fatal(err)
	}
}
