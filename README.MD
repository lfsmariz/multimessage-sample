# Instruções para Execução do Projeto

Este projeto utiliza Docker Compose para orquestrar os serviços necessários e Go para a implementação da lógica de consumo de mensagens.

## Pré-requisitos

- Docker e Docker Compose instalados em sua máquina.
- Go instalado (versão 1.23 ou superior).

## Executando o Docker Compose

1. Navegue até o diretório onde está localizado o arquivo `docker-compose.yml`.

   ```bash
   cd multimessage-sample
   ```

2. Execute o Docker Compose para iniciar os serviços.

   ```bash
   docker compose up -d
   ```

   O parâmetro `-d` executa os serviços em segundo plano.

3. Verifique se os serviços estão em execução.

   ```bash
   docker compose ps
   ```

## Executando o arquivo main.go

1. Navegue até o diretório onde está localizado o arquivo `main.go`.

   ```bash
   cd multimessage-sample
   ```

2. Execute o arquivo Go.

   ```bash
   go run main.go
   ```

## Executando o arquivo main.go da pasta sample-api-consumer

1. Navegue até o diretório `sample-api-consumer`.

   ```bash
   cd multimessage-sample/sample-api-consumer
   ```

2. Execute o arquivo Go.

   ```bash
   go run main.go
   ```

## Considerações Finais

Certifique-se de que todos os serviços estejam funcionando corretamente e que as dependências do Go estejam instaladas. Para instalar as dependências, você pode usar:

```bash
go mod tidy
```


Se você encontrar algum problema, verifique os logs dos serviços com:

```bash
docker compose logs -f
```
