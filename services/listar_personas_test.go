package services_test

import (
	"errors"
	"testing"

	"github.com/danysoftdev/p-go-list/models"
    "github.com/danysoftdev/p-go-list/services"
	"github.com/danysoftdev/p-go-list/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestListarPersonas_Success(t *testing.T) {
	mockRepo := new(mocks.MockPersonaRepo)
    services.Repo = mockRepo

	personasMock := []models.Persona{
		{Documento: "123", Nombre: "Ana", Apellido: "Díaz"},
		{Documento: "456", Nombre: "Luis", Apellido: "Pérez"},
	}

	mockRepo.On("ObtenerPersonas").Return(personasMock, nil)
	services.SetPersonaRepository(mockRepo) // 👈 Aquí se inyecta el mock

	personas, err := services.ListarPersonas()

	assert.NoError(t, err)
	assert.Equal(t, 2, len(personas))
	assert.Equal(t, "Ana", personas[0].Nombre)
	mockRepo.AssertExpectations(t)
}

func TestListarPersonas_Error(t *testing.T) {
	mockRepo := new(mocks.MockPersonaRepo)
    services.Repo = mockRepo

	mockRepo.On("ObtenerPersonas").Return([]models.Persona(nil), errors.New("fallo al obtener"))
	services.SetPersonaRepository(mockRepo)

	personas, err := services.ListarPersonas()

	assert.Error(t, err)
	assert.Nil(t, personas)
	mockRepo.AssertExpectations(t)
}
