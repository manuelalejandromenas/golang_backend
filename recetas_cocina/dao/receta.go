package dao

import "fmt"

type Receta struct {
	Serial int
	IdPlato int
}

type IngredientesEnReceta struct {
	SerialReceta int
	IdPlatoReceta int
	IdIngrediente int
	IdUnidadMedida int
	Cantidad int
}

type PasosEnReceta struct {
	SerialReceta int
	IdPlatoReceta int
	Serial int
	Descripcion string
}
