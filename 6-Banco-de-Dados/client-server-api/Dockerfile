# Utiliza uma imagem base do Golang
FROM golang

# Cria um diretório de trabalho
WORKDIR /app

# Copia os arquivos Go do servidor
COPY server/*.go ./

# Copia os arquivos Go do cliente
COPY client/*.go ./client/

# Copia os arquivos Go do pacote fetch
# COPY fetch/*.go ./fetch/

# Copia os arquivos de dependência do Go
COPY go.mod go.sum ./

# Cria o diretório para o banco de dados
RUN mkdir /db

# Cria o arquivo de banco de dados vazio dentro do diretório criado
RUN touch /db/cotacoes.db

# Compila o binário do servidor
RUN go build -o server .

# Garante que o binário tenha permissões de execução
RUN chmod +x server

# Define o comando a ser executado quando o container iniciar
CMD ["./server"]