package graficos

import (
	"encoding/base64"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

// StartServer inicia el servidor web.
func StartServer(port, imageDir string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		imagenes, _ := obtenerImagenCodificada(imageDir)

		renderizarPagina(w, imagenes)
	})

	http.ListenAndServe(":"+port, mux)
}

// RenderizarPagina renderiza la página HTML.
func renderizarPagina(w http.ResponseWriter, imagen Imagen) {
	tmpl, err := template.New("index").Parse(indexHTML)
	if err != nil {
		http.Error(w, "Error con el HTML", http.StatusInternalServerError)
		return
	}

	data := struct {
		Imagen Imagen
	}{
		Imagen: imagen,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error con el HTML", http.StatusInternalServerError)
	}
}

type Imagen struct {
	Nombre string
	Base64 string
}

func obtenerImagenCodificada(rutaImagen string) (Imagen, error) {
	var imagen Imagen

	contenido, err := ioutil.ReadFile(rutaImagen)
	if err != nil {
		return Imagen{}, err
	}

	nombre := filepath.Base(rutaImagen)
	base64str := base64.StdEncoding.EncodeToString(contenido)

	imagen = Imagen{
		Nombre: nombre,
		Base64: base64str,
	}

	return imagen, nil
}
