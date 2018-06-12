
CREATE TABLE plato (
	id_plato SERIAL,
	nombre	varchar(30) NOT NULL,
	descripcion varchar(200),
	PRIMARY KEY (id_plato)
);


CREATE TABLE receta (
	serial  SERIAL,
	id_plato int,
	FOREIGN KEY (id_plato) REFERENCES plato(id_plato),
	PRIMARY KEY (serial, id_plato)
);

CREATE TABLE ingrediente (
	id_ingrediente SERIAL,
	nombre varchar(30) NOT NULL,
	descripcion varchar(100),
	PRIMARY KEY (id_ingrediente)
);

CREATE TABLE unidad_medida (
	id_unidad_medida SERIAL,
	abreviacion varchar(20) NOT NULL,
	nombre varchar(30) NOT NULL,
	descripcion varchar(100),
	PRIMARY KEY (id_unidad_medida)
);

CREATE TABLE ingredientes_en_receta (
	serial_receta int, 
	id_plato_receta int,
	id_ingrediente int,
	id_unidad_medida int,
	cantidad int NOT NULL,
	FOREIGN KEY (serial_receta, id_plato_receta) REFERENCES receta (serial, id_plato),
	FOREIGN KEY (id_ingrediente) REFERENCES ingrediente (id_ingrediente),
	FOREIGN KEY (id_unidad_medida) REFERENCES unidad_medida (id_unidad_medida),
	PRIMARY KEY (serial_receta, id_plato_receta, id_ingrediente)
);

CREATE TABLE pasos_receta (
	serial_receta int,
	id_plato_receta int,
	serial SERIAL,
	descripcion varchar(200) NOT NULL,
	FOREIGN KEY (serial_receta, id_plato_receta) REFERENCES receta (serial, id_plato),
	PRIMARY KEY (serial_receta, id_plato_receta, serial)
);
	
