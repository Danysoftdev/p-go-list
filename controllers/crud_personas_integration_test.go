//go:build integration
// +build integration

package controllers_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/danysoftdev/p-go-list/config"
	"github.com/danysoftdev/p-go-list/controllers"
	"github.com/danysoftdev/p-go-list/models"
	"github.com/danysoftdev/p-go-list/repositories"
	"github.com/danysoftdev/p-go-list/services"
	"go.mongodb.org/mongo-driver/bson"
)

func TestEndpointsControllerIntegration(t *testing.T) {
	ctx := context.Background()

	// 1. Iniciar contenedor MongoDB
	req := testcontainers.ContainerRequest{
		Image:        "mongo:6.0",
		ExposedPorts: []string{"27017/tcp"},
		WaitingFor:   wait.ForListeningPort("27017/tcp").WithStartupTimeout(20 * time.Second),
	}
	mongoC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	assert.NoError(t, err)
	defer mongoC.Terminate(ctx)

	// 2. Configurar entorno
	endpoint, err := mongoC.Endpoint(ctx, "")
	assert.NoError(t, err)

	os.Setenv("MONGO_URI", "mongodb://"+endpoint)
	os.Setenv("MONGO_DB", "testdb")
	os.Setenv("COLLECTION_NAME", "personas_test")

	err = config.ConectarMongo()
	assert.NoError(t, err)
	defer config.CerrarMongo()

	repositories.SetCollection(config.Collection)
	services.SetPersonaRepository(repositories.RealPersonaRepository{})

	// 3. Limpiar colección
	_, err = config.Collection.DeleteMany(ctx, bson.M{})
	assert.NoError(t, err)

	// 4. Insertar persona
	persona := models.Persona{
		Documento: "999",
		Nombre:    "Lucía",
		Apellido:  "Pérez",
		Edad:      30,
		Correo:    "lucia@example.com",
		Telefono:  "3111234567",
		Direccion: "Calle 123",
	}
	_, err = config.Collection.InsertOne(ctx, persona)
	assert.NoError(t, err)

	// 5. Configurar router y endpoint
	router := mux.NewRouter()
	router.HandleFunc("/personas", controllers.ObtenerPersonas).Methods("GET")

	// 6. Ejecutar request
	reqObtener := httptest.NewRequest("GET", "/personas", nil)
	resObtener := httptest.NewRecorder()
	router.ServeHTTP(resObtener, reqObtener)

	assert.Equal(t, http.StatusOK, resObtener.Code)

	// 7. Leer y deserializar la respuesta
	content, _ := io.ReadAll(resObtener.Body)

	var personas []models.Persona
	err = json.Unmarshal(content, &personas)
	assert.NoError(t, err)
	assert.Len(t, personas, 1)
	assert.Equal(t, "Lucía", personas[0].Nombre)
	assert.Equal(t, "999", personas[0].Documento)
}
