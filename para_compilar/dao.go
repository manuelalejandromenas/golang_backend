package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
)

type CRUD interface {
	existeEnBD() bool
	crearEnBD()
	consultarEnBD()
	actualizarEnBD()
	eliminarEnBD()
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "M10//29//1996"
	dbname   = "recetas_cocina"
)

func con_comillas(palabra string) string {
	return fmt.Sprintf(`'%v'`, palabra)
}

func crearConexionBD() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	var base_datos *sql.DB = db
	return base_datos
}

type Receta struct {
	IdReceta     int
	Nombre       string
	Descripcion  string
	Ingredientes string
	Pasos        string
}

func (r *Receta) existeEnBD() bool {
	var base_datos *sql.DB = crearConexionBD()
	var db sql.DB = *base_datos
	defer db.Close()

	var nombre string
	err := db.QueryRow(`SELECT nombre FROM receta WHERE id_receta=$1`, r.IdReceta).Scan(&nombre)

	existe := true
	if err == sql.ErrNoRows {
		existe = false
	}

	return existe
}

func (r *Receta) crearEnBD() {
	var base_datos *sql.DB = crearConexionBD()
	var db sql.DB = *base_datos
	defer db.Close()

	sqlStatement := `INSERT INTO receta(nombre, descripcion, ingredientes, pasos) VALUES($1, $2, $3, $4) RETURNING id_receta;`

	id := 0
	err := db.QueryRow(sqlStatement, con_comillas(r.Nombre), con_comillas(r.Descripcion), con_comillas(r.Ingredientes), con_comillas(r.Pasos)).Scan(&id)
	r.IdReceta = id
	if err != nil {
		panic(err)
	}
}

func (r *Receta) consultarEnBD() {
	var base_datos *sql.DB = crearConexionBD()
	var db sql.DB = *base_datos
	defer db.Close()

	err := db.QueryRow(`SELECT nombre, descripcion, ingredientes, pasos FROM receta WHERE id_receta=$1;`, r.IdReceta).Scan(&r.Nombre, &r.Descripcion, &r.Ingredientes, &r.Pasos)
	if err != nil {
		panic(err)
	}
}

func (r *Receta) actualizarEnBD() {
	var base_datos *sql.DB = crearConexionBD()
	var db sql.DB = *base_datos
	defer db.Close()

	sqlStatement := `UPDATE receta SET nombre = $2, descripcion = $3, ingredientes = $4, pasos = $5 WHERE id_receta = $1;`
	_, err := db.Exec(sqlStatement, r.IdReceta, con_comillas(r.Nombre), con_comillas(r.Descripcion), con_comillas(r.Ingredientes), con_comillas(r.Pasos))
	if err != nil {
		panic(err)
	}
}

func (r *Receta) eliminarEnBD() {
	var base_datos *sql.DB = crearConexionBD()
	var db sql.DB = *base_datos
	defer db.Close()

	sqlStatement := `DELETE FROM receta WHERE id_receta = $1;`
	_, err := db.Exec(sqlStatement, r.IdReceta)
	if err != nil {
		panic(err)
	}
}

func listarRecetas() []Receta {
	var base_datos *sql.DB = crearConexionBD()
	var db sql.DB = *base_datos
	defer db.Close()

	rows, err := db.Query("SELECT * FROM receta;")
	if err != nil {
		panic(err)
	}

	var recetas []Receta
	defer rows.Close()

	for rows.Next() {
		var (
			id_receta    int
			nombre       string
			descripcion  string
			ingredientes string
			pasos        string
		)
		if err := rows.Scan(&id_receta, &nombre, &descripcion, &ingredientes, &pasos); err != nil {
			panic(err)
		}
		var receta Receta
		receta.IdReceta = id_receta
		receta.Nombre = nombre
		receta.Descripcion = descripcion
		receta.Ingredientes = ingredientes
		receta.Pasos = pasos

		recetas = append(recetas, receta)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}

	return recetas

}

func pruebasDAO() {
	var receta_a_crear Receta = Receta{1, "Nombre 1", "Descripcion 1", "Ingredientes 1", "Pasos 1"}
	fmt.Printf("¿Existe en BD (1)?: %v", receta_a_crear.existeEnBD())
	fmt.Printf("CREAR RECETA 1")

	receta_a_crear.IdReceta = 3
	receta_a_crear.crearEnBD()

	fmt.Printf("¿Existe en BD (1)?: %v", receta_a_crear.existeEnBD())

	fmt.Printf("CONSULTAR RECETA 1")
	var receta_a_consultar Receta = Receta{1, "", "", "", ""}
	receta_a_consultar.consultarEnBD()

	fmt.Printf("Receta 1: %v", receta_a_consultar)

	fmt.Printf("ACTUALIZAR RECETA 1")
	var receta_a_modificar Receta = Receta{1, "2", "2", "2", "2"}
	receta_a_modificar.actualizarEnBD()

	var receta_a_consultar_2 Receta = Receta{1, "", "", "", ""}
	receta_a_consultar_2.consultarEnBD()
	fmt.Printf("Receta 1: %v", receta_a_consultar_2)

	fmt.Printf("ELIMINAR RECETA 1")
	var receta_a_eliminar Receta = Receta{1, "", "", "", ""}

	fmt.Printf("¿Existe en BD (1)?: %v", receta_a_eliminar.existeEnBD())
	receta_a_eliminar.eliminarEnBD()
	fmt.Printf("¿Existe en BD (1) despues de eliminacion?: %v", receta_a_eliminar.existeEnBD())

}

type RecetaJSON struct {
	nombre string
	descripcion string
	ingredientes string
	pasos string
}

func crearReceta(json string) {
	var json_bytes []byte = []byte(json)
	var receta_json RecetaJSON
	err := json.Unmarshal(json_bytes, &receta_json)
	if ( err != nil ) {
		panic(err)
	}
	receta := Receta{-1,receta_json.nombre,receta_json.descripcion,receta_json.ingredientes,receta_json.pasos}
	
	receta.crearEnBD()
}

func consultarReceta(id_receta int) string {
	var json string
	receta := Receta{id_receta,"","","",""}
	
	bool existe = receta.existeEnBD()
	if existe {
		receta.consultarEnBD()
		json_bytes, err := json.Marshall(receta)
		if ( err != nil ) {
			panic(err)
		}
		json = string(json_bytes[:])
	} else {
		json = "NO SE ENCONTRO RECETA."
	}
}

func listarRecetasEnJson() []string {
	var recetas []Receta = listarRecetas()
	var recetas_en_json []string
	for i, receta := range recetas {
		json_bytes, err := json.Marshall(receta)
		if ( err != nil ) {
			panic(err)
		}
		json = string(json_bytes[:])
		append(recetas_en_json, json)
	}
	
	return recetas_en_json
}

func modificarReceta(id_receta int, json string) bool {
	
	receta := Receta{id_receta,"","","",""}
	bool existe = receta.existeEnBD()
	if existe {
		var json_bytes []byte = []byte(json)
		var receta_json RecetaJSON
		err := json.Unmarshal(json_bytes, &receta_json)
		if ( err != nil ) {
			panic(err)
		}
		receta := Receta{id_receta,receta_json.nombre,receta_json.descripcion,receta_json.ingredientes,receta_json.pasos}
		receta.actualizarEnBD()
	}
	return existe
	
}

func eliminarReceta(id_receta int) bool {
	var json string
	receta := Receta{id_receta,"","","",""}
	
	bool existe = receta.existeEnBD()
	if existe {
		receta.eliminarEnBD()
	} 
	return existe
}

func pruebasJSON() {
	fmt.Printf("\nLISTADO\n")
	var recetas []string = listarRecetasEnJson()
	for i, v := range recetas {
		fmt.Printf("%v\n", v)
	}
	fmt.Printf("\nCREAR\n")
	crearReceta(`{"nombre":"PRUEBA 1","descripcion":"PRUEBA 1","ingredientes":"PRUEBA 1", "pasos":"PRUEBA 1"}`)
	fmt.Printf("\nCONSULTAR\n")
	fmt.Printf(consultarReceta(2))
	fmt.Printf("\nMODIFICAR\n")
	se_pudo_modificar := modificarReceta(3, `{"nombre":"PRUEBA 1","descripcion":"PRUEBA 1","ingredientes":"PRUEBA 1", "pasos":"PRUEBA 1"}`)
	
	fmt.Printf("\nLISTADO\n")
	var recetas2 []string = listarRecetasEnJson()
	for i, v := range recetas2 {
		fmt.Printf("%v\n", v)
	}
	
	
	
}
func main() {
	//pruebasDAO()
	pruebasJSON()
}
