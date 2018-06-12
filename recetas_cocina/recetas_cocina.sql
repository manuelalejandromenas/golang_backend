CREATE TABLE receta (
	id_receta SERIAL,
	nombre	varchar(30) NOT NULL,
	descripcion varchar(200),
	ingredientes varchar(300),
	pasos varchar(300),
	PRIMARY KEY (id_receta)
);