# INDICE
1. CATEGORIAS
  - [GET CATEGORIES](#get-categories)
  - [GET CATEGORY BY ID](#get-category-by-id)
  - [CREATE CATEGORY](#create-category)
  - [UPDATE CATEGORY](#update-category)
  - [DELETE CATEGORY](#delete-category)
2. MARCA
  - [GET MARCAS](#get-marcas)
  - [GET MARCA BY ID](#get-marca-by-id)
  - [CREATE MARCA](#create-marca)
  - [UPDATE MARCA](#update-marca)
  - [DELETE MARCA](#delete-marca)

# GET CATEGORIES

**ROUTE : ** `/category`

**METHOD : ** `GET`

**DESCRIPTION : ** `Devuelve todas las categorias en la web`

**RESPONSE : **
```json
{
  "data":[
    {
      "id_category": "int",
      "name": "string",
      "description": "string"
    }
  ]
}
```
**RESPONSE ERROR : **
```json
{
  "data": {
  	"error": "string",
    "message": "string",
    "status": "int"
  }
}
```
# GET CATEGORY BY ID

**ROUTE : ** `/category/{id}`

**METHOD : ** `GET`

**DESCRIPTION : ** `Devuelve una categoria por su id`

**RESPONSE : **
```json
{
  "data":{
    "id_category": "int",
    "name": "string",
    "description": "string"
  }
}
```
**RESPONSE ERROR : **
```json
{
  "data": {
  	"error": "string",
    "message": "string",
    "status": "int"
  }
}
```
# CREATE CATEGORY

**ROUTE : ** `/category`

**METHOD : ** `POST`

**DESCRIPTION : ** `Crea una categoria`

**BODY : **
```json
{
  "data": {
    "name": "string", *requerido
    "description": "string" *requerido
  }
}
```
**RESPONSE : **
```json
{
  "data":{
    "id_category": "int",
    "name": "string",
    "description": "string"
  }
}
```
**RESPONSE ERROR : **
```json
{
  "data": {
  	"error": "string",
    "message": "string",
    "status": "int"
  }
}
```
# UPDATE CATEGORY

**ROUTE : ** `/category/{id}`

**METHOD : ** `PUT`

**DESCRIPTION : ** `Actualiza una categoria por su id`

**BODY : **
```json
{
  "data": {
    "name": "string", *opcional
    "description": "string" *opcional
  }
}
```
**RESPONSE : **
```json
{
  "data":{
    "id_category": "int",
    "name": "string",
    "description": "string"
  }
}
```
**RESPONSE ERROR : **
```json
{
  "data": {
  	"error": "string",
    "message": "string",
    "status": "int"
  }
}
```
# DELETE CATEGORY

**ROUTE : ** `/category/{id}`

**METHOD : ** `DELETE`

**DESCRIPTION : ** `Elimina una categoria`

**RESPONSE : **
```json
{
  *sin contenido
}
```
**RESPONSE ERROR : **
```json
{
  "data": {
  	"error": "string",
    "message": "string",
    "status": "int"
  }
}
```
# GET MARCAS

**ROUTE : ** `/marca`

**METHOD : ** `GET`

**DESCRIPTION : ** `Devuelve todas las marcas en la web`

**RESPONSE : **
```json
{
  "data":[
    {
      "id_marca": "int",
      "name": "string",
      "description": "string",
      "image": "string"
    }
  ]
}
```
**RESPONSE ERROR : **
```json
{
  "data": {
  	"error": "string",
    "message": "string",
    "status": "int"
  }
}
```
# GET MARCA BY ID

**ROUTE : ** `/marca/{id}`

**METHOD : ** `GET`

**DESCRIPTION : ** `Devuelve una marca por su id`

**RESPONSE : **
```json
{
  "data":{
    "id_marca": "int",
    "name": "string",
    "description": "string",
    "image": "string"
  }
}
```
**RESPONSE ERROR : **
```json
{
  "data": {
  	"error": "string",
    "message": "string",
    "status": "int"
  }
}
```
# CREATE MARCA

**ROUTE : ** `/marca`

**METHOD : ** `POST`

**DESCRIPTION : ** `Crea una marca`

**BODY : **
```json
{
  "data": {
    "name": "string", *requerido
    "description": "string", *requerido
    "image": "string" *requerido
  }
}
```
**RESPONSE : **
```json
{
  "data":{
    "id_marca": "int",
    "name": "string",
    "description": "string",
    "image": "string"
  }
}
```
**RESPONSE ERROR : **
```json
{
  "data": {
  	"error": "string",
    "message": "string",
    "status": "int"
  }
}
```
# UPDATE MARCA

**ROUTE : ** `/marca/{id}`

**METHOD : ** `PUT`

**DESCRIPTION : ** `Actualiza una marca por su id`

**BODY : **
```json
{
  "data": {
    "name": "string", *opcional
    "description": "string", *opcional
    "image": "string" *opcional
  }
}
```
**RESPONSE : **
```json
{
  "data":{
    "id_marca": "int",
    "name": "string",
    "description": "string",
    "image": "string"
  }
}
```
**RESPONSE ERROR : **
```json
{
  "data": {
  	"error": "string",
    "message": "string",
    "status": "int"
  }
}
```
# DELETE MARCA

**ROUTE : ** `/marca/{id}`

**METHOD : ** `DELETE`

**DESCRIPTION : ** `Elimina una marca`

**RESPONSE : **
```json
{
  *sin contenido
}
```
**RESPONSE ERROR : **
```json
{
  "data": {
  	"error": "string",
    "message": "string",
    "status": "int"
  }
}
```
