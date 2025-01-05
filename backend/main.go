package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	// Importamos la biblioteca oficial de Cohere para trabajar con su API
	cohere "github.com/cohere-ai/cohere-go/v2"
	cohereclient "github.com/cohere-ai/cohere-go/v2/client"

	// Para cargar variables de entorno desde un archivo .env
	"github.com/joho/godotenv"
	// Biblioteca para manejar CORS (Cross-Origin Resource Sharing)
	"github.com/rs/cors"
)

// Definimos la estructura para la solicitud que recibirá el servidor
type SentimentRequest struct {
	Text string `json:"text"` // Campo "text" que contiene el texto a analizar
}

// Definimos la estructura para la respuesta que enviará el servidor
type SentimentResponse struct {
	Sentiment string  `json:"sentiment"`       // El sentimiento detectado: EMOCIONADO, NEGATIVO o NEUTRAL
	Score     float64 `json:"score"`           // Puntuación asociada al sentimiento
	Error     string  `json:"error,omitempty"` // Mensaje de error
}

// Función para analizar el sentimiento usando la API de Cohere
func analyzeSentiment(client *cohereclient.Client, text string) (*SentimentResponse, error) {
	// Realizamos una llamada a la API de Cohere con un mensaje personalizado
	response, err := client.Chat(
		context.Background(),
		&cohere.ChatRequest{
			Message: "Analiza el sentimiento y responde solamente con una palabra: EMOCIONADO, NEGATIVO o NEUTRAL del siguiente texto: " + text,
		},
	)

	if err != nil {
		// Si ocurre un error al comunicarnos con Cohere, lo retornamos
		return nil, err
	}

	// Aquí podrías implementar una lógica más sofisticada para procesar la respuesta
	// Creamos una respuesta simple con el texto recibido desde la API
	return &SentimentResponse{
		Sentiment: response.Text, // Texto retornado por la API
		Score:     0.5,           // Puntuación fija como placeholder
	}, nil
}

func main() {
	// Configuración inicial
	// Cargar variables de entorno desde .env
	if err := godotenv.Load(); err != nil {
		log.Println("No se pudo cargar el archivo .env")
	}

	// Obtenemos el token de la API de Cohere desde las variables de entorno
	cohereToken := os.Getenv("COHERE_API_KEY")
	if cohereToken == "" {
		// Si no encontramos el token, terminamos el programa con un error
		log.Fatal("COHERE_API_KEY no está configurada")
	}

	// Inicializamos el cliente de Cohere con el token
	client := cohereclient.NewClient(cohereclient.WithToken(cohereToken))

	// Configurar CORS
	// Configuramos las reglas de CORS para permitir peticiones desde otros orígenes
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},               // En producción, debemos cambiar "*" por dominios específicos
		AllowedMethods: []string{"POST", "OPTIONS"}, // Métodos permitidos
		AllowedHeaders: []string{"Content-Type"},    // Encabezados permitidos
	})

	// Handler para el análisis de sentimientos
	// Definimos el handler para la ruta "/analyze"
	http.HandleFunc("/analyze", func(w http.ResponseWriter, r *http.Request) {
		// Verificamos que el método de la solicitud sea POST
		if r.Method != http.MethodPost {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}

		// Intentamos decodificar el cuerpo de la solicitud en la estructura SentimentRequest
		var req SentimentRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Llamamos a la función para analizar el sentimiento
		response, err := analyzeSentiment(client, req.Text)
		if err != nil {
			// En caso de error, respondemos con un mensaje de error
			json.NewEncoder(w).Encode(SentimentResponse{
				Error: "Error al analizar sentimiento: " + err.Error(),
			})
			return
		}
		// Configuramos el tipo de contenido de la respuesta como JSON
		w.Header().Set("Content-Type", "application/json")
		// Enviamos la respuesta codificada en formato JSON
		json.NewEncoder(w).Encode(response)
	})

	// Inicio del servidor
	// Iniciamos el servidor en el puerto 8080 con soporte para CORS
	handler := corsHandler.Handler(http.DefaultServeMux)
	port := ":8080"
	log.Printf("Servidor iniciado en %s", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
