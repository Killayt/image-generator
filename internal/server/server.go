package server

import (
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/Killayt/image-generator/configs"
	"github.com/Killayt/image-generator/pkg/img"
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
	buffer, err := img.GenerateFavicon()
	if err != nil {
		log.Error(err)
	}
	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Error(err)
	}
}

func Run(conf configs.ConfI) {
	r := mux.NewRouter()

	r.HandleFunc("/", imgHandler)
	r.HandleFunc("/ping", pingHandler)
	r.HandleFunc("/robots", robotsHandler)
	r.HandleFunc("/favicon", faviconHandler)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("Server is starting ...")
		if err := http.ListenAndServe(":"+conf.GetPort(), r); err != nil {
			log.Fatal(err)
		}
	}()

	signalValue := <-sigs
	signal.Stop(sigs)
	log.Println("STOP SIGNAL", signalValue)
}
