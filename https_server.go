package main 
import (
	"fmt"
	"net/http"
)

// handler es la función que maneja las peticiones GET en el endpoint "/hello"
func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "¡Hola, mundo!")
	} else {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func main() {
	// Registra el handler para el endpoint "/hello"
	http.HandleFunc("/hello", handler)

	// Inicia el servidor en el puerto 8080
	fmt.Println("Servidor escuchando en http://localhost:8080/hello")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
