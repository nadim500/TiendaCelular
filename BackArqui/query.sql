# CREANDO LA BASE DE DATOS
CREATE DATABASE test

# NOS CONECTAMOS \c test
# CREAMOS LAS TABLAS CORRESPONDIENTES

# Tabla categoria("Alta","Media")
CREATE TABLE category(
    id_category serial PRIMARY KEY,
    name VARCHAR NOT NULL,
    description VARCHAR NOT NULL
)

# Tabla marca de celular("Motorola","Samsung","etc")
CREATE TABLE marca(
    id_marca serial PRIMARY KEY,
    name VARCHAR NOT NULL,
    description VARCHAR NOT NULL,
    image VARCHAR NOT NULL
)

# Tabla producto(El celular en si)
CREATE TABLE product(
    id_product serial PRIMARY KEY,
    name VARCHAR NOT NULL,
    first_description VARCHAR NOT NULL,
    second_description VARCHAR NOT NULL,
    feature VARCHAR NOT NULL,
    datecreated DATE NOT NULL DEFAULT CURRENT_DATE,
    images VARCHAR[],
    marca_id integer NOT NULL,
    category_id integer NOT NULL,
    FOREIGN KEY (marca_id) REFERENCES marca (id_marca),
    FOREIGN KEY (category_id) REFERENCES category (id_category)
)

# Tabla cliente(Cliente para la compra de celular)
CREATE TABLE client(
    id_client serial PRIMARY KEY,
    name VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    phone VARCHAR NOT NULL,
    direction VARCHAR NOT NULL,
    document VARCHAR NOT NULL,
    email_verified BOOLEAN,
    password VARCHAR,
    hash_password VARCHAR NOT NULL
)

# Tabla comentario(Comentario para el celular, el comentario es unico para el cliente en un celular)
CREATE TABLE annotation(
    client_id INTEGER NOT NULL,
    product_id INTEGER NOT NULL,
    datecreated DATE NOT NULL DEFAULT CURRENT_DATE,
    annotation VARCHAR NOT NULL,
    PRIMARY KEY (client_id,product_id),
    FOREIGN KEY (client_id) REFERENCES client (id_client),
    FOREIGN KEY (product_id) REFERENCES product (id_product)
)

#  Tabla bill(Tabla para la factura)
CREATE TABLE bill(
    id_bill serial PRIMARY KEY,
    datecreated DATE NOT NULL DEFAULT CURRENT_DATE,
    mount FLOAT,
    client_id INTEGER NOT NULL,
    FOREIGN KEY (client_id) REFERENCES client (id_client)
)

# Tabla product_has_bill(Tabla para unir bill y product)
CREATE TABLE product_has_bill(
    product_id INTEGER NOT NULL,
    bill_id INTEGER NOT NULL,
    mount FLOAT,
    quantity INTEGER NOT NULL,
    PRIMARY KEY (product_id,bill_id),
    FOREIGN KEY (product_id) REFERENCES product (id_product),
    FOREIGN KEY (bill_id) REFERENCES bill (id_bill)
)

# Tabla connectivity(Tabla para conectividad)
CREATE TABLE connectivity(
    id_connectivity serial PRIMARY KEY,
    name VARCHAR NOT NULL,
    description VARCHAR NOT NULL,
    image VARCHAR NOT NULL
)

# Tabla product_has_connectivity
CREATE TABLE product_has_connectivity(
    product_id INTEGER NOT NULL,
    connectivity_id INTEGER NOT NULL,
    PRIMARY KEY (product_id,connectivity_id),
    FOREIGN KEY (product_id) REFERENCES product (id_product),
    FOREIGN KEY (connectivity_id) REFERENCES connectivity (id_connectivity)
)