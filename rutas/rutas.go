package rutas

import (
	"net/http"
	"sistemas/connection"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type Empleado struct {
	Id     int
	Nombre string
	Correo string
}

var plantilla = template.Must(template.ParseGlob("plantillas/*"))

func Inicio(w http.ResponseWriter, r *http.Request) {

	conexionEstablecida := connection.Conexionbd()
	registros, err := conexionEstablecida.Query("SELECT * FROM empleados")

	if err != nil {
		panic(err.Error())
	}

	empleado := Empleado{}
	arregloEmpleado := []Empleado{}

	for registros.Next() {
		var id int
		var nombre, correo string
		err = registros.Scan(&id, &nombre, &correo)
		if err != nil {
			panic(err.Error())
		}
		empleado.Id = id
		empleado.Nombre = nombre
		empleado.Correo = correo

		arregloEmpleado = append(arregloEmpleado, empleado)

	}

	plantilla.ExecuteTemplate(w, "inicio", arregloEmpleado)

}

func Crear(w http.ResponseWriter, r *http.Request) {

	plantilla.ExecuteTemplate(w, "crear", nil)

}

func Insertar(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		nombre := r.FormValue("nombre")
		correo := r.FormValue("correo")

		conexionEstablecida := connection.Conexionbd()
		insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO empleados(nombre, correo) VALUES (?,?)")

		if err != nil {
			panic(err.Error())
		}
		insertarRegistros.Exec(nombre, correo)

		http.Redirect(w, r, "/", 301)

	}

}

func Editar(w http.ResponseWriter, r *http.Request) {

	idEmpleado := r.URL.Query().Get("id")
	conexionEstablecida := connection.Conexionbd()
	registros, err := conexionEstablecida.Query("SELECT * FROM empleados WHERE id=?", idEmpleado)

	if err != nil {
		panic(err.Error())
	}

	empleado := Empleado{}

	for registros.Next() {
		var id int
		var nombre, correo string
		err = registros.Scan(&id, &nombre, &correo)
		if err != nil {
			panic(err.Error())
		}
		empleado.Id = id
		empleado.Nombre = nombre
		empleado.Correo = correo
	}

	plantilla.ExecuteTemplate(w, "editar", empleado)
}

func Actualizar(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		id := r.FormValue("id")
		nombre := r.FormValue("nombre")
		correo := r.FormValue("correo")

		conexionEstablecida := connection.Conexionbd()
		insertarRegistros, err := conexionEstablecida.Prepare("UPDATE empleados SET nombre=?, correo=? WHERE id=?")

		if err != nil {
			panic(err.Error())
		}
		insertarRegistros.Exec(nombre, correo, id)

		http.Redirect(w, r, "/", 301)

	}

}

func Borrar(w http.ResponseWriter, r *http.Request) {

	idEmpleado := r.URL.Query().Get("id")
	conexionEstablecida := connection.Conexionbd()
	borrarRegistros, err := conexionEstablecida.Prepare("DELETE FROM empleados WHERE id=?")

	if err != nil {
		panic(err.Error())
	}
	borrarRegistros.Exec(idEmpleado)

	http.Redirect(w, r, "/", 301)
}
