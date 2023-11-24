package main

import (
	"backend/api/routes"
	"context"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	firebase "firebase.google.com/go"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
)

func main() {
	// Inicializa la configuración de Firebase

	// Buscar el archivo de credenciales en el directorio actual
	matchingPattern := "./gha-creds-*.json"
	matches, err := filepath.Glob(matchingPattern)
	if err != nil {
		fmt.Printf("Error al buscar el archivo de credenciales: %v\n", err)

		return
	}

	if len(matches) == 0 {
		fmt.Println("No se encontraron archivos de credenciales.")

		return
	}

	// Utilizar el primer archivo coincidente (puedes ajustar esto según tus necesidades)
	pathToCredentials := matches[0]

	ctx := context.Background()
	opt := option.WithCredentialsFile(pathToCredentials)
	config := &firebase.Config{ProjectID: "test-5eebf"}
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatalln(err)
	}

	// Inicializa el cliente de Firestore
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	// Cerrar el cliente de Firestore cuando ya no se necesite
	defer client.Close()

	// Crear un enrutador utilizando gorilla/mux
	r := mux.NewRouter()
	routes.ConfigureRoutes(r, app)
	log.Println("ListenOn")
	log.Fatal(http.ListenAndServe(":3000", r))
}
