document.addEventListener('DOMContentLoaded', () => {
    // Elementos del DOM
    const textInput = document.getElementById('textInput');
    const analyzeBtn = document.getElementById('analyzeBtn');
    const loading = document.getElementById('loading');
    const error = document.getElementById('error');
    const result = document.getElementById('result');
    const sentimentElement = document.getElementById('sentiment');
    const scoreElement = document.getElementById('score');

    // Configuración
    const API_URL = 'http://localhost:8080/analyze';
    const SENTIMENT_COLORS = {
        'POSITIVO': '#00ff9d',
        'NEGATIVO': '#ff0055',
        'NEUTRAL': '#ffaa00'
    };

    // Funciones auxiliares
    const showLoading = (show) => {
        loading.classList.toggle('visible', show);
        analyzeBtn.disabled = show;
    };

    const showError = (message) => {
        error.textContent = message;
        error.classList.add('visible');
        result.classList.remove('visible');
    };

    const showResult = (sentiment, score) => {
        error.classList.remove('visible');
        result.classList.add('visible');
        
        sentimentElement.textContent = sentiment;
        sentimentElement.style.color = SENTIMENT_COLORS[sentiment] || '#ffffff';
        scoreElement.textContent = `Intensidad: ${(score * 100).toFixed(1)}%`;
    };

    // Función principal de análisis
    const analyzeSentiment = async (text) => {
        try {
            const response = await fetch(API_URL, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ text })
            });

            if (!response.ok) {
                throw new Error('Error en la respuesta del servidor');
            }

            const data = await response.json();
            
            if (data.error) {
                throw new Error(data.error);
            }

            return data;
        } catch (error) {
            throw new Error(error.message || 'Error al conectar con el servidor');
        }
    };

    // Event listeners
    analyzeBtn.addEventListener('click', async () => {
        const text = textInput.value.trim();
        
        if (!text) {
            showError('Por favor ingresa algún texto para analizar');
            return;
        }

        showLoading(true);
        error.classList.remove('visible');

        try {
            const result = await analyzeSentiment(text);
            showResult(result.sentiment, result.score);
        } catch (err) {
            showError(err.message);
        } finally {
            showLoading(false);
        }
    });
});