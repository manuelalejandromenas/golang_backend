package main

/*
	Se utilizaron dos librerias externas:
	gorilla/mux para el manejo de peticiones HTTP
	lib/pq como driver basico para la conexion con la base de datos en postgres

	De las librerias propias de go:
	"database/sql" para SQL, "encoding/json" para los JSON en el HTTP,
	"fmt" para impresion en flujos, "log" para control de errores, "net/http"
	para peticiones http, "strconv" para conversiones de string
*/

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strconv"
)

/*
	Interfaz CRUD, la idea es que cada clase de la REST API implemente esta
	interfaz, se crea pensando a futuro, ya que solo se maneja una clase (Recetas)
*/
type CRUD interface {
	existeEnBD() bool
	crearEnBD()
	consultarEnBD()
	actualizarEnBD()
	eliminarEnBD()
}

/*
	Datos de la conexion con la BD, no se usa encriptacion al tratarse de una prueba
*/
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "M10//29//1996"
	dbname   = "recetas_cocina"
)

/*	Funcion de nombre autoindicativo */
func con_comillas(palabra string) string {
	return fmt.Sprintf(`'%v'`, palabra)
}

/*  Funcion que crea una conexion con la base de datos, y la deja abierta, con
el objetivo de ser usada posteriormente por otras funciones / metodos */
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

/*  "Clase" basica de recetas, forma parte del DAO de la REST API */
type Receta struct {
	IdReceta     int
	Nombre       string
	Descripcion  string
	Ingredientes string
	Pasos        string
}

/*  Metodo para verificar existencia, de la clase receta */
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

/*  Metodo para crear en la base de datos, leyendo los atributos  en receta */
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

/*  Metodo para consultar en la base de datos, leyendo los atributos  en receta */
func (r *Receta) consultarEnBD() {
	var base_datos *sql.DB = crearConexionBD()
	var db sql.DB = *base_datos
	defer db.Close()

	err := db.QueryRow(`SELECT nombre, descripcion, ingredientes, pasos FROM receta WHERE id_receta=$1;`, r.IdReceta).Scan(&r.Nombre, &r.Descripcion, &r.Ingredientes, &r.Pasos)
	if err != nil {
		panic(err)
	}
}

/*  Metodo para actualizar en la base de datos, leyendo los atributos  en receta */
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

/*  Metodo para eliminar en la base de datos, leyendo los atributos  en receta */
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

/*  Funcion para listar todas las recetas en la base de datos, en otros lenguajes
seria un metodo de clase, al no encontrarse la forma adecuada de hacerlo en go,
se creo una funcion. */
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

/*  Pruebas del modelo DAO, solo funcionan si la base de datos esta recien creada
    y no tiene informacion alguna */
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

/*  RecetaJSON, es la misma clase receta sin su estructura, se crea con motivos
de facilitar implementacion */
type RecetaJSON struct {
	Nombre       string
	Descripcion  string
	Ingredientes string
	Pasos        string
}

/* Endpoint para listar recetas  (GET /recetas)
   Tiene dos posibles salidas:
   (StatusInternalServerError, listado_vacio) -> Error interno
   (StatusOk, listado de recetas) -> Consulta exitosa
*/
func ListarRecetasEndpoint(w http.ResponseWriter, r *http.Request) {
	var recetas []Receta = listarRecetas()

	json_bytes, err := json.Marshal(recetas)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "{}")
		panic(err)
	} else {
		listado_recetas := string(json_bytes[:])
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%v", listado_recetas)
	}
}

/* Endpoint para ver una receta  (GET /recetas/{id})
   Tiene cuatro posibles salidas:
   (StatusInternalServerError, listado_vacio) -> Error interno
   (StatusBadRequest, listado_vacio) -> Mala peticion
   (StatusNotFound, listado_vacio) -> No se encontro la receta con esa id
   (StatusOk, receta) -> JSON que representa la receta
*/
func VerRecetaEndpoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "{}")

	} else {
		receta := Receta{id, "", "", "", ""}
		existe := receta.existeEnBD()

		if existe {
			receta.consultarEnBD()
			json_bytes, err := json.Marshal(receta)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "{}")
				panic(err)
			} else {
				receta_json := string(json_bytes[:])
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, "%v", receta_json)
			}

		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "{}")
		}
	}
}

/* Endpoint para crear recetas (POST /recetas)
   Tiene dos posibles salidas:
   (StatusInternalServerError, listado_vacio) -> Error interno
   (StatusOk, receta) -> Impresion de la receta creada (no JSON)
*/
func CrearRecetaEndpoint(w http.ResponseWriter, r *http.Request) {
	var receta_json RecetaJSON
	err := json.NewDecoder(r.Body).Decode(&receta_json)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "{}")
		panic(err)
	} else {
		receta := Receta{-1, receta_json.Nombre, receta_json.Descripcion, receta_json.Ingredientes, receta_json.Pasos}
		receta.crearEnBD()
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%v", receta)
	}
}

/* Endpoint para modificar recetas (POST /recetas/{id})
   Tiene cuatro posibles salidas:
   (StatusInternalServerError, listado_vacio) -> Error interno
   (StatusBadRequest, listado_vacio) -> Mala peticion
   (StatusNotFound, listado_vacio) -> No se encontro la receta con esa id
   (StatusOk, receta) -> Impresion de la receta modificada (no JSON)
*/
func ModificarRecetaEndpoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "{}")
	} else {
		var receta_json RecetaJSON
		_ = json.NewDecoder(r.Body).Decode(&receta_json)

		verificar_existencia := Receta{id, "", "", "", ""}
		existe := verificar_existencia.existeEnBD()
		if existe {
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "{}")
				panic(err)
			} else {
				receta := Receta{id, receta_json.Nombre, receta_json.Descripcion, receta_json.Ingredientes, receta_json.Pasos}
				receta.actualizarEnBD()
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, "%v", receta)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "{}")
		}
	}
}

/* Endpoint para eliminar recetas (DELETE /recetas/{id})
   Tiene tres posibles salidas
   (StatusBadRequest, listado_vacio) -> Mala peticion
   (StatusNotFound, listado_vacio) -> No se encontro la receta con esa id
   (StatusOk, receta) -> Impresion de la receta eliminada (no JSON)
*/
func EliminarRecetaEndpoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "{}")
	} else {
		receta := Receta{id, "", "", "", ""}
		existe := receta.existeEnBD()
		if existe {
			receta.consultarEnBD()
			receta.eliminarEnBD()
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%v", receta)
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "{}")
		}
	}
}

/*
	Funcion principal, aqui se asocian los endpoints a sus respectivas rutas
	Muy importante uso de la libreria gorilla/mux
*/
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/recetas", ListarRecetasEndpoint).Methods("GET")
	router.HandleFunc("/recetas/{id}", VerRecetaEndpoint).Methods("GET")
	router.HandleFunc("/recetas", CrearRecetaEndpoint).Methods("POST")
	router.HandleFunc("/recetas/{id}", ModificarRecetaEndpoint).Methods("POST")
	router.HandleFunc("/recetas/{id}", EliminarRecetaEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":80", router))
}
