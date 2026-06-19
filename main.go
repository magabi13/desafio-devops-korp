package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Resposta struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}

var requestsCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "korp_http_requests_total",
		Help: "Quantidade total de requisicoes recebidas no endpoint /projeto-korp",
	},
	[]string{"path"},
)

func init() {
	prometheus.MustRegister(requestsCounter)
}

func projetoKorpHandler(w http.ResponseWriter, r *http.Request) {
	requestsCounter.WithLabelValues("/projeto-korp").Inc()

	w.Header().Set("Content-Type", "application/json")
	horarioUTC := time.Now().UTC().Format(time.RFC3339)

	resposta := Resposta{
		Nome:    "Projeto Korp",
		Horario: horarioUTC,
	}

	json.NewEncoder(w).Encode(resposta)
}

func main() {
	http.HandleFunc("/projeto-korp", projetoKorpHandler)
	http.Handle("/metrics", promhttp.Handler())

	log.Println("Servidor rodando na porta 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
