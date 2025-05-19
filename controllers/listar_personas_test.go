package controllers_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danysoftdev/p-go-list/controllers"
	"github.com/danysoftdev/p-go-list/models"
	"github.com/danysoftdev/p-go-list/services"
	"github.com/danysoftdev/p-go-list/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestObtenerPersonasController_Success(t *testing.T) {
	mockRepo := new(mocks.MockPersonaRepo)
	services.SetPersonaRepository(mockRepo)

	mockData := []models.Persona{
		{Documento: "1", Nombre: "Ana"},
		{Documento: "2", Nombre: "Luis"},
	}

	mockRepo.On("ObtenerPersonas").Return(mockData, nil)

	req := httptest.NewRequest(http.MethodGet, "/personas", nil)
	rr := httptest.NewRecorder()

	controllers.ObtenerPersonas(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var respuesta []models.Persona
	err := json.NewDecoder(rr.Body).Decode(&respuesta)
	assert.NoError(t, err)
	assert.Len(t, respuesta, 2)
	assert.Equal(t, "Ana", respuesta[0].Nombre)

	mockRepo.AssertExpectations(t)
}

func TestObtenerPersonasController_Error(t *testing.T) {
	mockRepo := new(mocks.MockPersonaRepo)
	services.SetPersonaRepository(mockRepo)

	mockRepo.On("ObtenerPersonas").Return([]models.Persona{}, errors.New("fallo inesperado"))

	req := httptest.NewRequest(http.MethodGet, "/personas", nil)
	rr := httptest.NewRecorder()

	controllers.ObtenerPersonas(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Contains(t, rr.Body.String(), "Error al obtener personas")

	mockRepo.AssertExpectations(t)
}