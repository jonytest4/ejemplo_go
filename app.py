from flask import Flask, send_from_directory
import os
from flask_cors import CORS

app = Flask(__name__, static_folder='static', template_folder='templates')

# Habilitar CORS para permitir solicitudes de tu front-end en el puerto 3000
CORS(app, origins="http://localhost:3000", supports_credentials=True)

# Ruta para servir archivos estáticos
@app.route('/')
def index():
    return send_from_directory(os.path.join(app.root_path, 'templates'), 'index.html')

# Ruta para servir archivos estáticos (JS, CSS, imágenes)
@app.route('/static/<path:path>')
def send_static(path):
    return send_from_directory(os.path.join(app.root_path, 'static'), path)

# Ejecutar el servidor Flask en el puerto 3000
if __name__ == "__main__":
    app.run(host='0.0.0.0', port=3000)
