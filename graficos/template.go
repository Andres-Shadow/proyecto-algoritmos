package graficos

const indexHTML = `
<!DOCTYPE html>
<html>
<head>
    <title>Servidor Web en Go</title>
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.5.2/css/bootstrap.min.css">
    <style>

		.container {
			position: absolute;
			top: 50%;
			left: 50%;
			transform: translate(-50%, -50%);
		}
        .fondo-cielo {
            background-color: #87CEEB; 
        }
        .card {
			background-color: #000; 
			color: #fff;
			margin: 10px auto; /* Centra el recuadro horizontalmente */
			display: inline-block; /* Ajusta el ancho al contenido y permite centrarlo */
		}
        .card-img-top {
            margin: -10px -10px 0 -10px; 
        }
        .card-title {
            color: #fff; 
            margin: 0; 
            padding: 10px; 
        }
        .recuadro-inferior {
            background-color: #000; 
            color: #fff; 
            text-align: center; 
            padding: 10px; 
            position: bottom; 
            bottom: 0; 
        }
    </style>
</head>
<body>
    <div class="container">

        <div class="fondo-cielo text-white p-3">
            <h1 >Proyecto Final Analisis de Algoritmos</h1>
        </div>

        <br/>
        <div class="row">
            <div class="col-md-4 mb-4">
                <div class="card">
                    <img class="card-img-top" src="data:image/png;base64,{{ .Imagen.Base64 }}" alt="{{ .Imagen.Nombre }}">
                </div>
            </div>
        </div>

        <!-- Recuadro inferior centrado y sin conectar con las paredes -->
        <div class="recuadro-inferior p3">
            <p>Analisis de algoritmos</p>
            <p>Andrés Mauricio Dussán Bastidas</p>
            <p>Luis Fernando Arenas Ramirez</p>
        </div>
        
    </div>
</body>
</html>

`
