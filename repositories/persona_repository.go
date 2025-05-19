package repositories

import (
	"context"
	"time"

	"github.com/danysoftdev/p-go-list/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

// Permite inyectar la colección desde fuera (ideal para pruebas)
func SetCollection(c *mongo.Collection) {
	collection = c
}

// ObtenerPersonas devuelve una lista de todas las personas
func ObtenerPersonas() ([]models.Persona, error) {
	var personas []models.Persona
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var persona models.Persona
		if err := cursor.Decode(&persona); err != nil {
			return nil, err
		}
		personas = append(personas, persona)
	}

	return personas, nil
}

type RealPersonaRepository struct{}

func (r RealPersonaRepository) ObtenerPersonas() ([]models.Persona, error) {
	return ObtenerPersonas()
}
