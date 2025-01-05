package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	cohere "github.com/cohere-ai/cohere-go/v2"
	cohereclient "github.com/cohere-ai/cohere-go/v2/client"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

type SentimentRequest struct {
	Text string `json:"text"`
}

type SentimentResponse struct {
	Sentiment string  `json:"sentiment"`
	Score     float64 `json:"score"`
	Error     string  `json:"error,omitempty"`
}

func analyzeSentiment(client *cohereclient.Client, text string) (*SentimentResponse, error) {
	response, err := client.Chat(
		context.Background(),
		&cohere.ChatRequest{
			Message: "Analiza el sentimiento y responde solamente con una palabra: EMOCIONADO, NEGATIVO o NEUTRAL del siguiente texto: " + text,
		},
	)

	if err != nil {
		return nil, err
	}

	// Aquí podrías implementar una lógica más sofisticada para procesar la respuesta
	return &SentimentResponse{
		Sentiment: response.Text,
		Score:     0.5, // Valor por defecto
	}, nil
}

func main() {
	// Configuración inicial
	// Cargar variables de entorno desde .env
	if err := godotenv.Load(); err != nil {
		log.Println("No se pudo cargar el archivo .env")
	}

	// Obtener token de variable de entorno
	cohereToken := os.Getenv("COHERE_API_KEY")
	if cohereToken == "" {
		log.Fatal("COHERE_API_KEY no está configurada")
	}

	// Inicializar cliente de Cohere
	client := cohereclient.NewClient(cohereclient.WithToken(cohereToken))

	// Configurar CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // En producción, especifica los dominios permitidos
		AllowedMethods: []string{"POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	})

	// Handler para el análisis de sentimientos
	http.HandleFunc("/analyze", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}

		var req SentimentRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response, err := analyzeSentiment(client, req.Text)
		if err != nil {
			json.NewEncoder(w).Encode(SentimentResponse{
				Error: "Error al analizar sentimiento: " + err.Error(),
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	// Iniciar servidor
	handler := corsHandler.Handler(http.DefaultServeMux)
	port := ":8080"
	log.Printf("Servidor iniciado en %s", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
