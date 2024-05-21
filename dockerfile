# Usar a imagem base do Go
FROM golang:1.16

# Definir o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copiar o arquivo go.mod e go.sum
COPY go.mod go.sum ./

# Baixar as dependências
RUN go mod download

# Copiar o código-fonte da aplicação
COPY . .

# Compilar a aplicação
RUN go build -o loadtester

# Definir o comando de entrada
ENTRYPOINT ["./loadtester"]
