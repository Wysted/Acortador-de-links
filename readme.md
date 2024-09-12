# Acortador de links

Hecho con Go - GinGonic - PostgreSQL

## Iniciar proyecto

```sh
go run .
```

## Ejemplo de peticiones

#### Subir enlace

POST http://localhost:8081/

```json
{
    "ID": "1",
    "Name": "WystedGithub",
    "URL": "https://github.com/Wysted"
}
```

PATCH http://localhost:8081/:id

### Actualizar enlace

```json
{
    "url": "http://www.google.com/"
}
```

o

```json
{
    "name": "Github"
}
```

### Eliminar enlace

DELETE http://localhost:8081/:id

### Usar api

GET http://localhost:8081/:name
