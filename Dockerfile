# Usa una imagen base de Go
FROM golang:1.23.1-alpine

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos go.mod y go.sum y descarga las dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copia el resto de los archivos de la aplicación
COPY . .

# Compila la aplicación
RUN go build -o main ./cmd/main.go

# Expone el puerto en el que la aplicación se ejecutará
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]