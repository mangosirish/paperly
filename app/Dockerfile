FROM golang:1.23-alpine

# Define el directorio de trabajo dentro del contenedor
WORKDIR /app

# Instala Air, una herramienta para recarga en caliente durante el desarrollo
RUN go install github.com/air-verse/air@latest

# Copia los archivos de dependencias y descarga los módulos
COPY go.mod go.sum ./
RUN go mod download

# Copia el resto del código de la aplicación
COPY . .

# Establece el PATH para incluir los binarios de Go (incluyendo el de Air)
ENV PATH="/go/bin:${PATH}"

CMD [ "air" ]