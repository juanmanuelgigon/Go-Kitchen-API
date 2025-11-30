FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# Compila el binario en /app/dist
RUN go build -o /app/dist .

# Usa la ruta absoluta para ejecutar el binario
CMD ["/app/dist"]