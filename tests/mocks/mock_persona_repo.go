package mocks

import (
	"github.com/danysoftdev/p-go-list/models"
	"github.com/stretchr/testify/mock"
)

// MockPersonaRepo implementa la interfaz PersonaRepository para pruebas
type MockPersonaRepo struct {
	mock.Mock
}

func (m *MockPersonaRepo) ObtenerPersonas() ([]models.Persona, error) {
	args := m.Called()
	return args.Get(0).([]models.Persona), args.Error(1)
}
