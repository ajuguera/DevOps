# Etapa de construcción
FROM golang:alpine AS build-env

WORKDIR /app

# Copiamos los archivos y compilamos
COPY . .
RUN go build -o main .

# Etapa final
FROM alpine

WORKDIR /app

# Copiamos el binario desde la etapa de construcción
COPY --from=build-env /app/main .

# Exponer el puerto 8081
EXPOSE 8081

# Ejecutar la aplicación
CMD ["./main"]
