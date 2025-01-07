document.addEventListener('DOMContentLoaded', () => {
    // Elementos del DOM
    // Referencia al campo de texto donde el usuario ingresa el texto a analizar
    const textInput = document.getElementById('textInput');
    // Botón que el usuario presiona para iniciar el análisis
    const analyzeBtn = document.getElementById('analyzeBtn');
    // Elemento que muestra un indicador de carga mientras se procesa la solicitud
    const loading = document.getElementById('loading');
    // Elemento donde se muestran errores al usuario
    const error = document.getElementById('error');
    // Contenedor donde se muestran los resultados del análisis
    const result = document.getElementById('result');
    // Elemento que muestra el sentimiento detectado (POSITIVO, NEGATIVO, NEUTRAL)
    const sentimentElement = document.getElementById('sentiment');
    // Elemento que muestra la intensidad del sentimiento en porcentaje
    const scoreElement = document.getElementById('score');

    // Configuración
    // URL de la API donde se enviarán las solicitudes para el análisis
    const API_URL = 'http://localhost:8080/analyze';
    // Colores asociados a cada tipo de sentimiento, para hacer la UI más interactiva
    const SENTIMENT_COLORS = {
        'EMOCIONADO.': '#3380ff', // Verde para sentimientos positivos
        'NEGATIVO.': '#ff5733', // Rojo para negativos
        'NEUTRAL': '#ffaa00' // Amarillo para neutral
    };

    // Funciones auxiliares
    // Muestra u oculta el indicador de carga y deshabilita el botón mientras se procesa
    const showLoading = (show) => { 
        loading.classList.toggle('visible', show); // Agrega o quita la clase "visible"
        analyzeBtn.disabled = show; // Deshabilita o habilita el botón
    };
    // Muestra un mensaje de error en la interfaz
    const showError = (message) => {
        error.textContent = message; // Escribe el mensaje en el contenedor de error
        error.classList.add('visible'); // Hace visible el mensaje de error
        result.classList.remove('visible'); // Oculta los resultados previos, si los hay
    };
    // Muestra los resultados del análisis en la interfaz
    const showResult = (sentiment, score) => {
        error.classList.remove('visible'); // Asegura que no se vea el error
        result.classList.add('visible'); // Muestra el contenedor de resultados
        
        sentimentElement.textContent = sentiment.toUpperCase(); // Escribe el sentimiento detectado
        sentimentElement.style.color = SENTIMENT_COLORS[sentiment] || '#ffffff'; // Cambia el color según el sentimiento
        scoreElement.textContent = `Intensidad: ${(score * 100).toFixed(1)}%`; // Muestra la intensidad como porcentaje
    };

    // Función principal de análisis
    // Envía una solicitud a la API para analizar el texto ingresado por el usuario
    const analyzeSentiment = async (text) => {
        try {
            // Realiza una solicitud POST a la API
            const response = await fetch(API_URL, {
                method: 'POST', // Método HTTP usado
                headers: {
                    'Content-Type': 'application/json', // Especifica que el cuerpo está en formato JSON
                },
                body: JSON.stringify({ text }) // Envía el texto como un objeto JSON
            });
            // Si la respuesta no es exitosa (status diferente de 200), lanza un error
            if (!response.ok) {
                throw new Error('Error en la respuesta del servidor');
            }
            // Convierte la respuesta a JSON
            const data = await response.json();
            // Si la API devuelve un error, lanza ese mensaje como excepción
            if (data.error) {
                throw new Error(data.error);
            }
            // Retorna los datos obtenidos de la API
            return data;
        } catch (error) {
            // Si ocurre cualquier error en la solicitud, lanza un mensaje genérico
            throw new Error(error.message || 'Error al conectar con el servidor');
        }
    };

    // Event listeners
    // Agrega un evento al botón para iniciar el análisis cuando el usuario hace clic
    analyzeBtn.addEventListener('click', async () => {
        // Obtiene el texto ingresado por el usuario y elimina espacios innecesarios
        const text = textInput.value.trim();
        // Si el usuario no ingresó texto, muestra un error y detiene la ejecución
        if (!text) {
            showError('Por favor ingresa algún texto para analizar');
            return;
        }
        // Muestra el indicador de carga y oculta errores previos
        showLoading(true);
        error.classList.remove('visible');

        try {
            // Llama a la función de análisis y espera los resultados
            const result = await analyzeSentiment(text);
            // Muestra los resultados en la interfaz
            showResult(result.sentiment, result.score);
        } catch (err) {
            // Si ocurre un error, muestra el mensaje al usuario
            showError(err.message);
        } finally {
            // Oculta el indicador de carga cuando todo termina
            showLoading(false);
        }
    });
});