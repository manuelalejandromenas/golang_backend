package dao

import "fmt"

type GuardarEnBD interface {
	guardarEnBD()
}

type LeerDesdeBD interface {
	leerDesdeBD() []interface{}
}

type ImprimirAJSON() interface {
	imprimirAJSON() string
}

type LeerJSON() interface {
	leerJSON()  
}

