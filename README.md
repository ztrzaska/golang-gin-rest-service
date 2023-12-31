# Goolang gin REST API with go-swagger and mongo database

### Get started

This project demonstrates sample microservice, which use golang, gin REST API ang go-swagger.


### Swagger installation

```
git clone https://github.com/go-swagger/go-swagger
cd go-swagger
go install ./cmd/swagger
```

### Swagger doc API generation

```
swagger generate spec -o ./cmd/api/swaggerui/swagger.json --scan-models
```

### Mongo database

```
docker pull mongo
docker run -dp 27017:27017 mongo
```

### Reference Documentation

For further reference, please consider the following sections:

* [Golang documentation](https://go.dev/doc/)
* [Gin library documentation](https://gin-gonic.com/)
* [Go-swagger](https://github.com/go-swagger/go-swagger/)
