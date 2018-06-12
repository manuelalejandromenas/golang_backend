package recetas_cocina

import (
	"database/sql"
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

	sqlStatement := `INSERT INTO receta(nombre, descripcion, ingredientes, pasos) VALUES $1, $2, $3, $4);`

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

	err := db.QueryRow(`SELECT nombre, descripcion, ingredientes, pasos FROM receta WHERE id_receta=$1`, r.IdReceta).Scan(&r.Nombre, &r.Descripcion, &r.Ingredientes, &r.Pasos)
	if err != nil {
		panic(err)
	}
}

func (r *Receta) actualizarEnBD() {
	var base_datos *sql.DB = crearConexionBD()
	var db sql.DB = *base_datos
	defer db.Close()

	sqlStatement := `UPDATE receta SET nombre = $2, descripcion = $3, ingredientes = $4, pasos = $5 WHERE id_receta = $1`
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

func listarRecetas() []receta {
	var base_datos *sql.DB = crearConexionBD()
	var db sql.DB = *base_datos
	defer db.Close()

	rows, err := db.Query("SELECT name FROM users;", age)
	if err != nil {
		panic(err)
	}

	var recetas []receta
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
			log.Fatal(err)
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
		log.Fatal(err)
	}

	return recetas

}

func pruebasDAO() {
	var receta_a_crear Receta = Receta{1, "Nombre 1", "Descripcion 1", "Ingredientes 1", "Pasos 1"}
	fmt.Printf("¿Existe en BD (1)?: %v", receta1.existeEnBD())
	fmt.Printf("CREAR RECETA 1")

	receta_a_crear.IdReceta = 3
	receta_a_crear.crearEnBD()

	fmt.Printf("¿Existe en BD (1)?: %v", receta1.existeEnBD())

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

func main() {
	pruebasDAO()

}
