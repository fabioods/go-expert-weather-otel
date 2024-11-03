# Weather Lab with otel

Laboratório do go expert onde temos duas apps, a primeira é uma aplicação de cep e na sequência consome uma API de previsão do tempo e retorna a previsão do tempo para uma cidade.
Na segunda teremos uma app que consumira essa primeira app e retornará a previsão do tempo para uma cidade.
E teremos o otel configurado para monitorar as duas apps.

## Rodando a aplicação

Para rodar a aplicação, basta executar o seguinte comando na raiz do projeto:

```bash
docker-compose up -d --build
```

## Testando a aplicação

Para testar a aplicação, basta acessar o seguinte URL no seu navegador:

```bash
http://localhost:8081/weather/cep/84430000
```

Onde o 84430000 é o CEP da cidade que você deseja consultar a previsão do tempo.

## Vendo os traces no Jaeger

Para ver os traces no Jaeger, basta acessar a seguinte URL no seu navegador:

```bash
http://localhost:16686/search
```
