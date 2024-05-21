### README

# LoadTester

LoadTester é uma aplicação em Go projetada para realizar testes de carga em serviços web. O usuário pode especificar a URL do serviço, o número total de requisições e a quantidade de chamadas simultâneas via linha de comando. A aplicação gera um relatório detalhado após a execução dos testes.

## Estrutura do Projeto

A estrutura do projeto é a seguinte:

```
stress-test-go/
│
├── main.go
├── Dockerfile
├── go.mod
└── go.sum
```

- **main.go**: Contém o código-fonte principal da aplicação.
- **Dockerfile**: Arquivo de configuração do Docker para criar a imagem da aplicação.
- **go.mod**: Gerencia as dependências do projeto Go.
- **go.sum**: Mantém um hash das dependências para garantir a integridade do projeto.

## Pré-requisitos

Antes de começar, certifique-se de ter o seguinte instalado:

- [Docker](https://www.docker.com/)
- [Go](https://golang.org/)

## Como Construir e Executar a Aplicação

### Passo 1: Construir a Imagem Docker

Para construir a imagem Docker da aplicação, execute o seguinte comando no diretório raiz do projeto:

```sh
docker build -t loadtester .
```

Este comando cria uma imagem Docker chamada `loadtester` a partir do Dockerfile.

### Passo 2: Executar o Contêiner Docker

Para executar o contêiner Docker com a aplicação LoadTester, use o seguinte comando:

```sh
docker run loadtester --url=http://google.com --requests=1000 --concurrency=10
```

Substitua `http://google.com` pela URL do serviço que você deseja testar, `1000` pelo número total de requisições que deseja fazer, e `10` pela quantidade de chamadas simultâneas.

## Parâmetros CLI

A aplicação aceita os seguintes parâmetros via linha de comando:

- `--url`: URL do serviço a ser testado.
- `--requests`: Número total de requisições.
- `--concurrency`: Número de chamadas simultâneas.

## Relatório de Saída

Após a execução dos testes, a aplicação gera um relatório com as seguintes informações:

- Tempo total gasto na execução.
- Quantidade total de requisições realizadas.
- Quantidade de requisições com status HTTP 200.
- Distribuição de outros códigos de status HTTP (como 404, 500, etc.).

## Exemplo de Uso

Aqui está um exemplo de como usar a aplicação:

```sh
docker run loadtester --url=http://example.com --requests=500 --concurrency=20
```

Este comando realiza 500 requisições para `http://example.com` com uma concorrência de 20 chamadas simultâneas.

## Contribuição

Contribuições são bem-vindas! Se você encontrar problemas ou tiver sugestões de melhorias, sinta-se à vontade para abrir uma issue ou enviar um pull request.

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).