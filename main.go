package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"time"

	"github.com/julienschmidt/httprouter"
)

const ADD_PREFIX = "Adding: "

var list = PageData{Messages: make([]Message, 0)}

type timerMiddleware struct {
	next http.Handler
}

func NewTimerMiddleware(next http.Handler) *timerMiddleware {
	return &timerMiddleware{next: next}
}

func (m *timerMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	st := time.Now()
	m.next.ServeHTTP(w, r)
	log.Printf("served %s [%s] in %s\n", r.URL, r.Method, time.Since(st))
}

type PageData struct {
	Messages []Message
}

type Message struct {
	Text string
}

func main() {
	start := time.Now()

	file, err := os.OpenFile("messages.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	tpl := template.Must(template.ParseFiles("index.html"))
	atpl := template.Must(template.ParseFiles("admin.html"))
	log.Println("template loaded")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		tpl.Execute(w, list)
	})
	router.POST("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		tpl.Execute(w, list)
	})
	router.POST("/save", addHandler)
	router.GET("/admin/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		atpl.Execute(w, nil)
	})
	router.GET("/admin/ops/", adminHandler)
	router.GET("/admin/hydrate", hydrateHandler)

	log.Println("routes loaded")
	m := NewTimerMiddleware(router)
	server := &http.Server{
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 16,
		Addr:           ":" + port,
		Handler:        m,
	}

	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Panicf("cannot listen: %s", err)
	}

	log.Printf("listening on port %s (started in %s)", port, time.Since(start))
	log.Fatal(server.Serve(l))
}

func addHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	m := r.FormValue("message")
	log.Printf("%s%s\n", ADD_PREFIX, m)
	list.Messages = append(list.Messages, Message{Text: m})
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func hydrateHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	f, err := os.Open("messages.log")
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "unable to access the file")
		return
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		l := s.Text()
		if strings.Contains(l, ADD_PREFIX) {
			m := strings.Split(l, ADD_PREFIX)
			list.Messages = append(list.Messages, Message{Text: m[1]})
		}
	}

	if err := s.Err(); err != nil {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "error during file access")
		return
	}

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func adminHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	qv := r.URL.Query()
	cmd := qv.Get("cmd")
	arg := strings.Split(qv.Get("args"), ",")
	ll := fmt.Sprintf("serving %s", cmd)
	for _, v := range arg {
		ll = fmt.Sprintf("%s %s", ll, v)
	}
	log.Println(ll)
	c := exec.Command(cmd, arg...)
	out, err := c.CombinedOutput()
	if err != nil {
		log.Printf("unable to serve page: %s\n", err.Error())
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "unable to run the command")
		return
	}
	output := string(out[:])
	log.Printf("result: %s\n", output)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, output)
}
