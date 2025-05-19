package services

import (
	"github.com/danysoftdev/p-go-list/models"
	"github.com/danysoftdev/p-go-list/repositories"
)

var Repo repositories.PersonaRepository

func SetPersonaRepository(r repositories.PersonaRepository) {
	Repo = r
}

func ListarPersonas() ([]models.Persona, error) {
	return Repo.ObtenerPersonas()
}
