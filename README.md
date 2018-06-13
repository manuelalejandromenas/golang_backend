# golang_backend
REST API para recetas de cocina

# Servidor de implementacion 
Se implemento en un servidor de Amazon, su direccion ip publica es:
ec2-18-188-191-107.us-east-2.compute.amazonaws.com

# Sobre el repositorio
En recetas_cocina esta toda la implementacion de la API basica

En tutorial estan los archivos correspondientes a un tutorial medianamente extensivo de go,
de este link:
https://tour.golang.org

En para_compilar esta el archivo .go principal con la implementacion de la API basica,
se creo esta carpeta para facilitar la compilacion en el server

# Modelo SQL
Se utilizo Postgres como el servidor de base de datos, no se uso un ORM para la implementacion del DAO,
pero se hizo uso del driver basico de postgres y se simplifico el codigo correspondiente a la consultar
en base de datos

En recetas_cocina/datos.sql y recetas_cocina/recetas_cocina.sql se encuentran algunos ejemplos de datos y su modelo respectivo.

Los otros archivos .sql corresponden a un modelo inicial no implementado mas extenso


# Endpoints

## Endpoint para listar recetas  (GET /recetas)
Tiene dos posibles salidas:
(StatusInternalServerError, listado_vacio) -> Error interno
(StatusOk, listado de recetas) -> Consulta exitosa

## Endpoint para ver una receta  (GET /recetas/{id})
Tiene cuatro posibles salidas:
(StatusInternalServerError, listado_vacio) -> Error interno
(StatusBadRequest, listado_vacio) -> Mala peticion
(StatusNotFound, listado_vacio) -> No se encontro la receta con esa id
(StatusOk, receta) -> JSON que representa la receta

## Endpoint para crear una receta (POST /recetas)
Tiene dos posibles salidas:
(StatusInternalServerError, listado_vacio) -> Error interno
(StatusOk, receta) -> Impresion de la receta creada (no JSON)

## Endpoint para modificar recetas (POST /recetas/{id})
Tiene cuatro posibles salidas:
(StatusInternalServerError, listado_vacio) -> Error interno
(StatusBadRequest, listado_vacio) -> Mala peticion
(StatusNotFound, listado_vacio) -> No se encontro la receta con esa id
(StatusOk, receta) -> Impresion de la receta modificada (no JSON)

## Endpoint para eliminar recetas (DELETE /recetas/{id})
Tiene tres posibles salidas
(StatusBadRequest, listado_vacio) -> Mala peticion
(StatusNotFound, listado_vacio) -> No se encontro la receta con esa id
(StatusOk, receta) -> Impresion de la receta eliminada (no JSON)



   
