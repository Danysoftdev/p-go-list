package repositories

import "github.com/danysoftdev/p-go-list/models"

type PersonaRepository interface {
	ObtenerPersonas() ([]models.Persona, error)
}
