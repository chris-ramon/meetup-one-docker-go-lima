package main

import (
	"log"
	"net/http"
	"text/template"
	"time"
)

const (
	Waiting = "Esperando"
	Done    = "Terminado"
)

var (
	tasks []Task
)

func init() {
	tasks = []Task{
		Task{Name: "Enviar mails", Status: Waiting, Time: 4},
		Task{Name: "Generar PDFS", Status: Waiting, Time: 7},
		Task{Name: "Buscar y Actualizar Github Perfiles", Status: Waiting, Time: 9},
	}
}

type Task struct {
	Name   string
	Status string
	Time   time.Duration
}

func (t *Task) Run() {
	//ejecutarse la tarea y cambiar de estado
	time.Sleep(t.Time * time.Second)
	t.Status = Done
	log.Printf("task done: %v", t)
}

func RunTasks() {
	for _, task := range tasks {
		go task.Run()
	}
}

func RunHandler(w http.ResponseWriter, r *http.Request) {
	go RunTasks()
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, tasks)
}

func main() {
	log.Printf("Server running on port: :8080")
	http.HandleFunc("/run", RunHandler)
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":8080", nil)
}
