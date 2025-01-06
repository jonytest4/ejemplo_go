# Ejemplo de Uso de Golang como API

Este proyecto es un ejemplo práctico del uso de **Golang** como API en combinación con tecnologías front-end y Python. La aplicación incluye un backend desarrollado en Go y un frontend manejado mediante Flask (Python), HTML, CSS y JavaScript. 

## ¿Qué es Go (Golang)?

**Golang**, desarrollado por Google, es un lenguaje de programación diseñado para ser simple, eficiente y escalable. Su sintaxis clara y concisa, junto con su compilación rápida y capacidad para manejar concurrencia de manera eficiente, lo convierten en una excelente opción para proyectos de alto rendimiento.

### Características principales:
1. **Simplicidad y legibilidad:**
   - Golang prioriza una sintaxis clara y sencilla, facilitando el aprendizaje y la implementación rápida de proyectos.

2. **Compilación rápida:**
   - Genera binarios portátiles y eficientes en tiempo récord, mejorando tanto el desarrollo como el despliegue.

3. **Gestión automática de memoria:**
   - Incluye un **garbage collector**, un mecanismo que gestiona automáticamente la memoria del programa. Esto reduce errores comunes como fugas de memoria al liberar recursos no utilizados.

4. **Concurrencia eficiente:**
   - Utiliza **gorutinas**, que son hilos ligeros de ejecución. Estas permiten manejar miles de tareas concurrentes simultáneamente con un uso mínimo de recursos, optimizando aplicaciones que necesitan alto rendimiento.

5. **Compatibilidad multiplataforma:**
   - Soporta múltiples sistemas operativos como Windows, macOS y Linux.

6. **Ecosistema robusto y comunidad activa:**
   - Ofrece una amplia gama de herramientas, bibliotecas y soporte comunitario.

7. **Paradigma de programación:**
   - Es un lenguaje imperativo y concurrente que aprovecha al máximo los procesadores multinúcleo. Aunque incluye elementos de POO, no admite jerarquías ni herencia.

## Descripción del Proyecto

El proyecto demuestra cómo utilizar Golang como núcleo para una API REST mientras se integra con un frontend sencillo basado en Flask y tecnologías web estándar. Incluye funcionalidades para gestionar solicitudes y respuestas con soporte para herramientas modernas como AWS SDK y otras bibliotecas populares.

## Tecnologías Utilizadas

### Backend (Golang)
- **Librerías:**
  - [aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) `v1.32.7`
  - [smithy-go](https://github.com/aws/smithy-go) `v1.22.1`
  - [cohere-go](https://github.com/cohere-ai/cohere-go) `v2.12.3`
  - [google/uuid](https://github.com/google/uuid) `v1.4.0`
  - [godotenv](https://github.com/joho/godotenv) `v1.5.1`
  - [rs/cors](https://github.com/rs/cors) `v1.11.1`

### Frontend
- **Lenguajes:** HTML, CSS, JavaScript
- **Framework:** Flask (Python)

## Requisitos Previos

### Herramientas necesarias
1. **Golang** (versión `1.23.4`): Descárgala desde [aquí](https://go.dev/dl/).
2. **Python** (versión `3.13.1`): Descárgalo desde [aquí](https://www.python.org/downloads/).
3. **Git**
4. Instalación de los módulos necesarios para Go y Python.

## Instalación y Ejecución

### Paso 1: Clonar el repositorio
```bash
git clone https://github.com/jonytest4/ejemplo_go.git
cd ejemplo_go
```

### Paso 2: Configurar el backend (Golang)
1. Navega a la carpeta `backend`:
   ```bash
   cd backend
   ```
2. Instala las dependencias necesarias con el siguiente comando:
   ```bash
   go get nombre_dependencia
   ```
3. Ejecuta el archivo principal:
   ```bash
   go run main.go
   ```

### Paso 3: Configurar el frontend (Flask)
1. Navega a la carpeta del frontend:
   ```bash
   cd ../frontend
   ```
2. Instala las dependencias necesarias de Python (si las hay).
   ```bash
   pip install -r requirements.txt
   ```
3. Ejecuta la aplicación Flask:
   ```bash
   python app.py
   ```

### Paso 4: Acceso a la aplicación
- El backend estará disponible en el puerto configurado (por defecto `localhost:8080`).
- El frontend estará disponible en el puerto configurado por Flask (por defecto `localhost:5000`).

## Contribuciones

Si deseas contribuir a este proyecto, sigue estos pasos:
1. Haz un fork del repositorio.
2. Crea una nueva rama para tu funcionalidad: `git checkout -b feature/nueva-funcionalidad`.
3. Realiza tus cambios y haz un commit: `git commit -m "Añadida nueva funcionalidad"`.
4. Envía un pull request.

## Licencia

Este proyecto está bajo la Licencia MIT. Consulta el archivo `LICENSE` para más información.

---

¡Gracias por explorar este proyecto! Si tienes preguntas o sugerencias, no dudes en abrir un issue.
